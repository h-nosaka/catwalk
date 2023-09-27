package mysql

import (
	"bytes"
	"encoding/json"
	"fmt"
	"sort"

	"github.com/gertd/go-pluralize"
	"github.com/iancoleman/strcase"
)

type EnumType uint

const (
	EnumTypeString   = EnumType(0)
	EnumTypeUint     = EnumType(1)
	EnumTypeBitfield = EnumType(2)
	EnumTypeUnkown   = EnumType(255)
)

func (p EnumType) String() string {
	switch p {
	case EnumTypeString:
		return "string"
	case EnumTypeUint:
		return "uint"
	case EnumTypeBitfield:
		return "uint64"
	}
	return "unknown"
}

func EnumTypes(key string) EnumType {
	switch key {
	case "string":
		return EnumTypeString
	case "uint":
		return EnumTypeUint
	case "uint64":
		return EnumTypeBitfield
	}
	return EnumTypeUnkown
}

func (p EnumType) MarshalJSON() ([]byte, error) {
	return json.Marshal(p.String())
}

func (p *EnumType) UnmarshalJSON(data []byte) error {
	var value string
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*p = EnumTypes(value)
	return nil
}

type IEnum struct {
	Column string
	Type   EnumType
	Values map[string]interface{}
}

type ISortValue struct {
	Key   string
	Value interface{}
}

func NewEnum(column string, types EnumType, values map[string]interface{}) IEnum {
	return IEnum{
		Column: column,
		Type:   types,
		Values: values,
	}
}

func (p *IEnum) GoCreate(table *ITable) string {
	buf := bytes.NewBuffer([]byte{})
	name := p.GoStructName(table)
	dataType := p.Type.String()
	buf.WriteString(fmt.Sprintf("type %s %s\n", name, dataType))

	values := []ISortValue{}
	for key, value := range p.Values {
		values = append(values, ISortValue{Key: key, Value: value})
	}
	switch p.Type {
	case EnumTypeString:
		sort.Slice(values, func(a, b int) bool { return values[a].Value.(string) < values[b].Value.(string) })
	case EnumTypeUint, EnumTypeBitfield:
		sort.Slice(values, func(a, b int) bool { return values[a].Value.(int) < values[b].Value.(int) })
	}

	buf.WriteString("const (\n")
	switch p.Type {
	case EnumTypeString:
		for _, value := range values {
			buf.WriteString(fmt.Sprintf("\t%s%s = %s(\"%s\")\n", name, value.Key, name, value.Value.(string)))
		}
	case EnumTypeUint:
		for _, value := range values {
			buf.WriteString(fmt.Sprintf("\t%s%s = %s(%d)\n", name, value.Key, name, value.Value.(int)))
		}
	case EnumTypeBitfield:
		ok := true
		for _, value := range values {
			if ok {
				buf.WriteString(fmt.Sprintf("\t%s%s %s = 1 << iota\n", name, value.Key, name))
				ok = false
			} else {
				buf.WriteString(fmt.Sprintf("\t%s%s\n", name, value.Key))
			}
		}
	}
	buf.WriteString(")\n\n")

	switch p.Type {
	case EnumTypeString:
		buf.WriteString(fmt.Sprintf("func (p %s) String() string {\n", name))
		buf.WriteString("\treturn string(p)\n")
		buf.WriteString("}\n\n")
	case EnumTypeUint:
		buf.WriteString(fmt.Sprintf("func (p %s) String() string {\n", name))
		buf.WriteString("\tswitch p {\n")
		for _, value := range values {
			buf.WriteString(fmt.Sprintf("\tcase %s%s:\n", name, value.Key))
			buf.WriteString(fmt.Sprintf("\t\treturn \"%s\"\n", value.Key))
		}
		buf.WriteString("\t}\n")
		buf.WriteString("\treturn \"\"\n")
		buf.WriteString("}\n\n")

		con := pluralize.NewClient()
		names := con.Plural(name)
		buf.WriteString(fmt.Sprintf("func %s(key string) %s {\n", names, name))
		buf.WriteString("\tswitch key {\n")
		for _, value := range values {
			buf.WriteString(fmt.Sprintf("\tcase \"%s\":\n", value.Key))
			buf.WriteString(fmt.Sprintf("\t\treturn %s%s\n", name, value.Key))
		}
		buf.WriteString("\t}\n")
		buf.WriteString("\treturn 0\n")
		buf.WriteString("}\n\n")

		buf.WriteString(fmt.Sprintf("func (p %s) MarshalJSON() ([]byte, error) {\n", name))
		buf.WriteString("\treturn json.Marshal(p.String())")
		buf.WriteString("}\n\n")

		buf.WriteString(fmt.Sprintf("func (p *%s) UnmarshalJSON(data []byte) error {\n", name))
		buf.WriteString("\tvar value string\n")
		buf.WriteString("\tif err := json.Unmarshal(data, &value); err != nil {\n")
		buf.WriteString("\t\treturn err")
		buf.WriteString("\t}\n")
		buf.WriteString(fmt.Sprintf("\t*p = %s(value)\n", names))
		buf.WriteString("\treturn nil\n")
		buf.WriteString("}\n\n")
	case EnumTypeBitfield:
		buf.WriteString(fmt.Sprintf("func (p %s) Check(flag %s) bool {\n", name, name))
		buf.WriteString("\treturn (p & flag) == flag\n")
		buf.WriteString("}\n\n")
	}
	return buf.String()
}

func (p *IEnum) GoStructName(table *ITable) string {
	con := pluralize.NewClient()
	return fmt.Sprintf("%s%s", strcase.ToCamel(con.Singular(table.Name)), strcase.ToCamel(p.Column))
}
