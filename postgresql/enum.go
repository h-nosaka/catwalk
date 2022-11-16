package postgresql

import (
	"bytes"
	"fmt"
	"strings"

	"github.com/gertd/go-pluralize"
	"github.com/iancoleman/strcase"
)

type IEnum struct {
	Column   string
	IsString bool
	Values   map[string]interface{}
}

func NewEnum(column string, isString *bool, values map[string]interface{}) IEnum {
	is := false
	if isString != nil {
		is = *isString
	}
	return IEnum{
		Column:   column,
		IsString: is,
		Values:   values,
	}
}

func (p *IEnum) GoCreate(table *ITable) string {
	buf := bytes.NewBuffer([]byte{})
	name := p.GoStructName(table)
	dataType := "string"
	if !p.IsString {
		dataType = "int"
	}
	buf.WriteString(fmt.Sprintf("type %s %s\n", name, dataType))
	buf.WriteString("const (\n")
	if p.IsString {
		for key, value := range p.Values {
			buf.WriteString(fmt.Sprintf("\t%s%s = %s(\"%s\")\n", name, key, name, value))
		}
	} else {
		for key, value := range p.Values {
			buf.WriteString(fmt.Sprintf("\t%s%s = %s(%d)\n", name, key, name, value))
		}
	}
	buf.WriteString(")\n\n")
	if p.IsString {
		buf.WriteString(fmt.Sprintf("func (p %s) String() string {\n", name))
		buf.WriteString("\treturn string(p)\n")
		buf.WriteString("}\n\n")
	} else {
		buf.WriteString(fmt.Sprintf("func (p %s) Integer() int {\n", name))
		buf.WriteString("\treturn int(p)\n")
		buf.WriteString("}\n\n")
	}
	return buf.String()
}

func (p *IEnum) GoStructName(table *ITable) string {
	con := pluralize.NewClient()
	return fmt.Sprintf("%s%s", strcase.ToCamel(con.Singular(table.Name)), strcase.ToCamel(p.Column))
}

func (p *IEnum) DartCreate(table *ITable) string {
	buf := bytes.NewBuffer([]byte{})
	name := p.GoStructName(table)
	// enum定義
	buf.WriteString(fmt.Sprintf("enum %s {\n", name))
	for key := range p.Values {
		buf.WriteString(fmt.Sprintf("\t%s,\n", strings.ToLower(key)))
	}
	buf.WriteString("}\n\n")
	// enumの変換定義
	buf.WriteString(fmt.Sprintf("extension %sExt on %s {\n", name, name))
	buf.WriteString("\tstatic final enumValues = {\n")
	if p.IsString {
		for key, value := range p.Values {
			buf.WriteString(fmt.Sprintf("\t\t%s.%s: '%s',\n", name, strings.ToLower(key), value))
		}
	} else {
		for key, value := range p.Values {
			buf.WriteString(fmt.Sprintf("\t\t%s.%s: %d,\n", name, strings.ToLower(key), value))
		}
	}
	buf.WriteString("\t};\n")
	if p.IsString {
		buf.WriteString("\tString get value => enumValues[this] ?? '';\n")
	} else {
		buf.WriteString("\tint get value => enumValues[this] ?? 0;\n")
	}
	buf.WriteString("}\n\n")
	return buf.String()
}
