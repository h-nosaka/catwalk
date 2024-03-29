package mysql

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
	Name   string // DBName
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
	base.DB.Raw(fmt.Sprintf(GetTables, base.DBName)).Scan(&tables)
	for _, table := range tables {
		table.GetColumn()
		table.GetIndexes()
		table.GetForeignkeys()
		table.GetPartitions()
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
		fmt.Printf("Write Error: %s\n", err.Error())
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
	if err := base.DB.Exec(fmt.Sprintf(GetCreateDatabase, p.Name)).Error; err != nil {
		panic(err)
	}
	fmt.Printf("CreateDatabase comleted: %s\n", p.Name)
}

func (p *ISchema) Run() {
	src := NewSchemaFromDB()
	diff := p.Diff(src)
	lines := strings.Split(diff, ";\n")
	count := 0
	for _, sql := range lines {
		if len(sql) > 10 {
			if err := base.DB.Exec(sql).Error; err != nil {
				panic(err)
			}
			count++
		}
	}
	fmt.Printf("Run comleted: %d sql\n", count)
}

func (p *ISchema) Model(output ...string) {
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
