package catwalk

import (
	"bytes"
	"fmt"
	"os"
	"os/exec"
	"strings"

	"github.com/h-nosaka/catwalk/base"
	"github.com/h-nosaka/catwalk/mysql"
	"github.com/h-nosaka/catwalk/postgres"
	"github.com/iancoleman/strcase"
	"gopkg.in/yaml.v2"
)

func IsMySQL() bool {
	return base.GetEnv("RDB_TYPE", "") == "mysql"
}

func IsPostgres() bool {
	return base.GetEnv("RDB_TYPE", "") == "postgres"
}

func GetSchemaMode() SchemaMode {
	mode := SchemaModeES
	if IsMySQL() {
		mode = SchemaModeMySQL
	}
	if IsPostgres() {
		mode = SchemaModePostgres
	}
	return mode
}

func NewSchemaFromDB() *Schema {
	base.Init()
	rs := Schema{
		Name:   base.DBName,
		Tables: []Table{},
		Mode:   GetSchemaMode(),
	}
	// table取得
	data := rs.sqlTables()
	for _, record := range data {
		// fmt.Println(base.ToPrettyJson(record))
		name := record["TABLE_NAME"].(string)
		jcase := JsonCaseCamel
		if strcase.ToCamel(name) == name {
			jcase = JsonCasePascal
		}
		if strcase.ToSnake(name) == name {
			jcase = JsonCaseSnake
		}
		comment, ok := record["TABLE_COMMENT"].(string)
		if !ok {
			comment = ""
		}
		table := Table{
			Schema:    record["TABLE_SCHEMA"].(string),
			Name:      name,
			JsonCase:  jcase,
			Comment:   base.String(comment),
			Columns:   []Column{},
			Indexes:   []Index{},
			Relations: []Relation{},
			Enums:     []Enum{},
		}
		table.GetColumn()
		table.GetIndexes()
		table.GetForeignkeys()
		table.GetPartitions()
		rs.Tables = append(rs.Tables, table)
	}
	return &rs
}

func (p *Schema) sqlTables() []map[string]interface{} {
	switch p.Mode {
	case SchemaModeMySQL:
		return GetData("Tables", p.Name)
	case SchemaModePostgres:
		return GetData("Tables")
	}
	return []map[string]interface{}{}
}

func GetData(key string, values ...interface{}) []map[string]interface{} {
	mode := GetSchemaMode()
	var sql string
	var ok bool
	var rs []map[string]interface{}
	switch mode {
	case SchemaModeMySQL:
		sql, ok = mysql.SQLS[key]
	case SchemaModePostgres:
		sql, ok = postgres.SQLS[key]
	case SchemaModeES:
		panic("ES非対応")
	default:
		panic("未定義のデータベース")
	}
	if !ok {
		panic(fmt.Errorf("未定義のSQL: %s", key))
	}
	if err := base.DB.Raw(sql, values...).Scan(&rs).Error; err != nil {
		panic(err)
	}
	return rs
}

func GetParts(key string, values ...interface{}) string {
	var sql string
	var ok bool
	switch GetSchemaMode() {
	case SchemaModeMySQL:
		sql, ok = mysql.PARTS[key]
	case SchemaModePostgres:
		sql, ok = postgres.PARTS[key]
	case SchemaModeES:
		panic("ES非対応")
	default:
		panic("未定義のデータベース")
	}
	if !ok {
		panic(fmt.Errorf("未定義のSQLパーツ: %s", key))
	}
	return fmt.Sprintf(sql, values...)
}

func (p *Schema) CreateSql(filename string) {
	fp, err := os.Create(filename)
	if err != nil {
		panic(err)
	}
	defer fp.Close()
	buf := bytes.NewBuffer([]byte{})
	buf.WriteString(GetParts("SchemaCreate", p.Name))
	for _, item := range p.Tables {
		buf.WriteString(item.Create())
	}
	if _, err := fp.Write(buf.Bytes()); err != nil {
		fmt.Printf("WriteString Error: %s\n", err.Error())
	}
}

func (p *Schema) CreateYaml(filename string) {
	fp, err := os.Create(filename)
	if err != nil {
		panic(err)
	}
	defer fp.Close()
	data, err := yaml.Marshal(&p)
	if err != nil {
		panic(err)
	}
	if _, err := fp.Write(data); err != nil {
		fmt.Printf("WriteString Error: %s\n", err.Error())
	}
}

func (p *Schema) Diff(src *Schema) string {
	buf := bytes.NewBuffer([]byte{})
	for _, item := range p.Tables {
		buf.WriteString(item.Diff(&src.Tables))
	}
	return buf.String()
}

func (p *Schema) Run() {
	base.Init()
	src := NewSchemaFromDB()
	diff := p.Diff(src)
	lines := strings.Split(diff, ";\n")
	count := 0
	after := []string{}
	switch p.Mode {
	case SchemaModePostgres:
		if err := base.DB.Exec(`CREATE EXTENSION IF NOT EXISTS "uuid-ossp";`).Error; err != nil {
			panic(err)
		}
	}
	for _, sql := range lines {
		if len(sql) > 10 {
			// FOREIGN KEYは他テーブルとの兼ね合いがあるので、最後に実行する
			if strings.Contains(sql, "ALTER TABLE") && strings.Contains(sql, "FOREIGN KEY") {
				after = append(after, sql)
				continue
			}
			if err := base.DB.Exec(sql).Error; err != nil {
				fmt.Println(sql)
				// continue
				panic(err)
			}
			count++
		}
	}
	// FOREIGN KEYは最後にまとめて実行する
	for _, sql := range after {
		if err := base.DB.Exec(sql).Error; err != nil {
			fmt.Println(sql)
			panic(err)
		}
		count++
	}
	fmt.Printf("Run comleted: %d sql\n", count)
}

func (p *Schema) ForcePublic() *Schema {
	p.Name = "public"
	tables := []Table{}
	for _, t := range p.Tables {
		t.Schema = "public"
		tables = append(tables, t)
	}
	p.Tables = tables
	return p
}

func (p *Schema) CreateModel(output ...string) {
	gopath := "./app/models"
	if len(output) > 0 && len(output[0]) > 0 {
		gopath = output[0]
	}

	cnt := 0
	for _, table := range p.Tables {
		if table.CreateGoModel(gopath) {
			cnt++
		}
	}
	if err := exec.Command("goimports", "-w", fmt.Sprintf("%s/", gopath)).Run(); err != nil {
		panic(err)
	}
	if err := exec.Command("go", "fmt", fmt.Sprintf("%s/...", gopath)).Run(); err != nil {
		panic(err)
	}
	fmt.Printf("Model comleted: %d models\n", cnt)
}

func (p *Schema) CreateFixture(output ...string) {
	gopath := "./tests/fixtures"
	if len(output) > 0 && len(output[0]) > 0 {
		gopath = output[0]
	}

	cnt := 0
	for _, table := range p.Tables {
		if table.CreateGoFixture(gopath) {
			cnt++
		}
	}
	if err := exec.Command("goimports", "-w", fmt.Sprintf("%s/", gopath)).Run(); err != nil {
		panic(err)
	}
	if err := exec.Command("go", "fmt", fmt.Sprintf("%s/...", gopath)).Run(); err != nil {
		panic(err)
	}
	fmt.Printf("Fixture comleted: %d fixtures\n", cnt)
}

func (p *Schema) CreateSchema(output ...string) {
	gopath := "./platform/schema"
	if len(output) > 0 && len(output[0]) > 0 {
		gopath = output[0]
	}

	cnt := 0
	for _, table := range p.Tables {
		if table.CreateSchemaFile(gopath) {
			cnt++
		}
	}
	if err := exec.Command("goimports", "-w", fmt.Sprintf("%s/", gopath)).Run(); err != nil {
		panic(err)
	}
	if err := exec.Command("go", "fmt", fmt.Sprintf("%s/...", gopath)).Run(); err != nil {
		panic(err)
	}
	fmt.Printf("Schema comleted: %d schema\n", cnt)
}
