package catwalk

import (
	"bytes"
	"fmt"
	"os"
	"strings"

	"github.com/h-nosaka/catwalk/base"
	"github.com/h-nosaka/catwalk/catwalk/mysql"
	"github.com/h-nosaka/catwalk/catwalk/postgres"
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
		table := Table{
			Schema:    record["TABLE_SCHEMA"].(string),
			Name:      name,
			JsonCase:  jcase,
			Comment:   base.String(record["TABLE_COMMENT"].(string)),
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
