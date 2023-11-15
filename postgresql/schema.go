package postgresql

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"os/exec"
	"strings"

	"github.com/h-nosaka/catwalk/base"
	"gopkg.in/yaml.v2"
)

type ISchema struct {
	Name   string
	Tables []ITable
}

func NewSchema(yamlpath string) *ISchema {
	fp, err := os.Open(yamlpath)
	if err != nil {
		panic(err)
	}
	data, err := io.ReadAll(fp)
	if err != nil {
		panic(err)
	}
	rs := ISchema{}
	if err := yaml.Unmarshal(data, &rs); err != nil {
		panic(err)
	}
	return &rs
}

func NewSchemaFromDB() *ISchema {
	rs := ISchema{
		Name: base.DBName,
	}
	// table取得
	rs.Tables = []ITable{}
	tables := []ITable{}
	base.DB.Raw(GetTables).Scan(&tables)
	for _, table := range tables {
		table.GetColumn()
		table.GetIndexes()
		table.GetForeignkeys()
		rs.Tables = append(rs.Tables, table)
	}
	return &rs
}

func (p *ISchema) Dump() string {
	buf := bytes.NewBuffer([]byte{})

	for _, item := range p.Tables {
		buf.WriteString(item.Create())
	}

	return buf.String()
}

func (p *ISchema) Sql(filename string) {
	fp, err := os.Create(filename)
	if err != nil {
		panic(err)
	}
	defer fp.Close()
	if _, err := fp.WriteString(p.Dump()); err != nil {
		fmt.Printf("WriteString Error: %s\n", err.Error())
	}
}

func (p *ISchema) Yaml(filename string) {
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

func (p *ISchema) Diff(src *ISchema) string {
	buf := bytes.NewBuffer([]byte{})
	for _, item := range p.Tables {
		buf.WriteString(item.Diff(&src.Tables))
	}
	return buf.String()
}

func (p *ISchema) CreateDatabase(dbname ...string) {
	if dbname != nil && len(dbname[0]) > 0 {
		p.Name = dbname[0]
	}
	if err := base.DB.Exec(fmt.Sprintf(GetCreateDatabase, p.Name, base.GetEnv("RDB_USER", ""))).Error; err != nil {
		panic(err)
	}
	fmt.Printf("CreateDatabase comleted: %s\n", p.Name)
}

func (p *ISchema) Run() {
	base.Init()
	src := NewSchemaFromDB()
	diff := p.Diff(src)
	lines := strings.Split(diff, ";\n")
	count := 0
	after := []string{}
	if err := base.DB.Exec("CREATE EXTENSION IF NOT EXISTS \"uuid-ossp\";").Error; err != nil {
		fmt.Println(err)
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
				continue
				// panic(err)
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

func (p *ISchema) Model(output ...string) {
	base.Init()
	gopath := "./models/dump"
	if len(output) > 0 && len(output[0]) > 0 {
		gopath = output[0]
	}

	for _, table := range p.Tables {
		table.CreateGoModel(gopath)
	}
	if err := exec.Command("go", "fmt", fmt.Sprintf("%s/...", gopath)).Run(); err != nil {
		panic(err)
	}
	fmt.Printf("Model comleted: %d models\n", len(p.Tables))
}

func (p *ISchema) Fixture(output ...string) {
	base.Init()
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

func (p *ISchema) CreateSchema(output string) {
	for _, table := range p.Tables {
		filename := fmt.Sprintf("%s/%s.go", output, table.Name)
		if err := os.WriteFile(filename, table.CreateSchemaFile(), 0666); err != nil {
			panic(err)
		}
	}
	if err := exec.Command("go", "fmt", fmt.Sprintf("%s/...", output)).Run(); err != nil {
		panic(err)
	}
}

func (p *ISchema) ER(output string) {
	buf := bytes.NewBuffer([]byte{})
	buf.WriteString("```mermaid\nerDiagram\n\n")
	for _, item := range p.Tables {
		buf.WriteString(fmt.Sprintf("%s {\n", item.Name))
		for _, col := range item.Columns {
			comment := ""
			if col.Comment != nil {
				comment = *col.Comment
			}
			buf.WriteString(fmt.Sprintf("  %s %s \"%s\"\n", col.Name, col.DataType, comment))
		}
		buf.WriteString("}\n\n")
	}
	for _, item := range p.Tables {
		buf.WriteString(fmt.Sprintf("%%%% %s\n", item.Name))
		for _, rel := range item.Foreignkeys {
			if rel.HasOne {
				buf.WriteString(fmt.Sprintf("%s ||--|| %s : \"foreignkey\"\n", item.Name, rel.RefTable))
			} else {
				buf.WriteString(fmt.Sprintf("%s ||--o{ %s : \"foreignkey\"\n", item.Name, rel.RefTable))
			}
		}
		for _, rel := range item.Relations {
			if rel.HasOne {
				buf.WriteString(fmt.Sprintf("%s ||--|| %s : \"relation\"\n", item.Name, rel.RefTable))
			} else {
				buf.WriteString(fmt.Sprintf("%s ||--o{ %s : \"relation\"\n", item.Name, rel.RefTable))
			}
		}
		buf.WriteString("\n")
	}

	buf.WriteString("```\n")
	if err := os.WriteFile(output, buf.Bytes(), 0666); err != nil {
		panic(err)
	}
}
