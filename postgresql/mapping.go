package postgresql

import (
	"bytes"
	"fmt"
	"io"
	"os"

	"github.com/gertd/go-pluralize"
	"github.com/iancoleman/strcase"
	"gopkg.in/yaml.v2"
)

type Mappings struct {
	Mappings []Mapping
	Actions  []Action
	Enums    []Enum
}

type Mapping struct {
	Tablename   string
	Reftable    string
	Association string
	Columnname  string
	Refcolumn   string
}

type Action struct {
	Tablename string
	Imports   *[]string
	Action    string
}

type Enum struct {
	Tablename  string
	Columnname string
	Type       string
	Enums      map[string]interface{}
}

func (p Mappings) Load(filename string) *Mappings {
	fp, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	data, err := io.ReadAll(fp)
	if err != nil {
		panic(err)
	}
	if err := yaml.Unmarshal(data, &p); err != nil {
		panic(err)
	}
	return &p
}

func (p *Mapping) GetRelation(model string) string {
	if p.Association == "ONE" {
		return fmt.Sprintf("*%s", model)
	}
	return fmt.Sprintf("*[]%s", model)
}

func (p *Mapping) GetModel(model string) string {
	if p.Association == "ONE" {
		return model
	}
	con := pluralize.NewClient()
	return con.Plural(model)
}

func (p *Enum) GoCreate() string {
	buf := bytes.NewBuffer([]byte{})
	name := p.GoStructName()
	buf.WriteString(fmt.Sprintf("type %s %s\n", name, p.Type))
	buf.WriteString("const (\n")
	switch p.Type {
	case "string":
		for key, value := range p.Enums {
			buf.WriteString(fmt.Sprintf("\t%s%s = %s(\"%s\")\n", name, key, name, value))
		}
	case "int":
		for key, value := range p.Enums {
			buf.WriteString(fmt.Sprintf("\t%s%s = %s(%d)\n", name, key, name, value))
		}
	}
	buf.WriteString(")\n\n")
	switch p.Type {
	case "string":
		buf.WriteString(fmt.Sprintf("func (p %s) String() string {\n", name))
		buf.WriteString("\treturn string(p)\n")
		buf.WriteString("}\n\n")
	case "int":
		buf.WriteString(fmt.Sprintf("func (p %s) Integer() int {\n", name))
		buf.WriteString("\treturn int(p)\n")
		buf.WriteString("}\n\n")
	}
	return buf.String()
}

func (p *Enum) GoStructName() string {
	con := pluralize.NewClient()
	return fmt.Sprintf("%s%s", strcase.ToCamel(con.Singular(p.Tablename)), strcase.ToCamel(p.Columnname))
}
