package mysql

import (
	"bytes"
	"fmt"
	"strings"

	"github.com/h-nosaka/catwalk/base"
	"golang.org/x/exp/slices"
)

type IColumn struct {
	Name     string  `gorm:"column:COLUMN_NAME"`
	DataType string  `gorm:"column:COLUMN_TYPE"`
	Extra    *string `gorm:"column:EXTRA" yaml:"extra,omitempty"`
	Defaults *string `gorm:"column:COLUMN_DEFAULT" yaml:"default,omitempty"`
	Null     bool    `gorm:"->:false" yaml:"nullable,omitempty"`
	IsNull   string  `gorm:"column:IS_NULLABLE" yaml:"-"`
	Comment  *string `gorm:"column:COLUMN_COMMENT" yaml:"comment,omitempty"`
	Rename   *string `gorm:"->:false" yaml:"rename,omitempty"`
}

func NewColumn(name string, dataType string, extra *string, defaults *string, nullable *bool, comment *string, rename *string) IColumn {
	null := true
	if nullable != nil {
		null = *nullable
	}
	return IColumn{
		Name:     name,
		DataType: dataType,
		Extra:    extra,
		Defaults: defaults,
		Null:     null,
		Comment:  comment,
		Rename:   rename,
	}
}

func DefaultColumn(cols ...IColumn) []IColumn {
	rs := []IColumn{
		NewColumn("id", "bigint(20) unsigned", base.String("auto_increment"), nil, base.Bool(false), base.String("primary key"), nil),
	}
	rs = append(rs, cols...)
	rs = append(
		rs,
		TimestampColumn()...,
	)
	return rs
}

func TimestampColumn() []IColumn {
	return []IColumn{
		NewColumn("created_at", "timestamp", nil, base.String("CURRENT_TIMESTAMP"), nil, base.String("作成日"), nil),
		NewColumn("updated_at", "timestamp", nil, base.String("CURRENT_TIMESTAMP"), nil, base.String("更新日"), nil),
	}
}

func (p *IColumn) GetDefault() string {
	if p.Defaults != nil && *p.Defaults != "" {
		return fmt.Sprintf(" DEFAULT %s", *p.Defaults)
	}
	return ""
}

func (p *IColumn) GetExtra() string {
	if p.Extra != nil && *p.Extra != "" {
		return fmt.Sprintf(" %s", *p.Extra)
	}
	return ""
}

func (p *IColumn) GetNullable() string {
	if !p.Null {
		return " NOT NULL"
	}
	return " NULL"
}

func (p *IColumn) GetComment() string {
	if p.Comment != nil && *p.Comment != "" {
		return fmt.Sprintf(" COMMENT '%s'", *p.Comment)
	}
	return ""
}

func (p *IColumn) Append() string {
	return fmt.Sprintf(
		"\t%s %s",
		p.Name,
		p._append(),
	)
}

func (p *IColumn) _append() string {
	return fmt.Sprintf(
		"%s%s%s%s%s",
		p.DataType,
		p.GetDefault(),
		p.GetExtra(),
		p.GetNullable(),
		p.GetComment(),
	)
}

func (p *IColumn) Create(t *ITable) string {
	return fmt.Sprintf(
		"ALTER TABLE %s ADD %s %s;\n\n",
		t.Name,
		p.Name,
		p._append(),
	)
}

func (p *IColumn) Drop(t *ITable) string {
	return fmt.Sprintf(
		"ALTER TABLE %s DROP COLUMN `%s`;\n\n",
		t.Name,
		p.Name,
	)
}

func (p *IColumn) SetModify(t *ITable) string {
	return fmt.Sprintf(
		"ALTER TABLE %s MODIFY COLUMN `%s` %s;\n\n",
		t.Name,
		p.Name,
		p._append(),
	)
}

func (p *IColumn) RenameColumn(t *ITable) string {
	if p.Rename != nil {
		return fmt.Sprintf(
			"ALTER TABLE %s CHANGE `%s` %s %s;\n\n",
			t.Name,
			p.Name,
			*p.Rename,
			p._append(),
		)
	}
	return ""
}

func (p *IColumn) IsDrop(table *ITable) string {
	buf := bytes.NewBuffer([]byte{})
	dest := IColumn{}
	for _, item := range table.Columns {
		if item.Name == p.Name {
			dest = item
		}
	}
	if dest.Name == "" {
		buf.WriteString(p.Drop(table))
	}
	return buf.String()
}

func (p IColumn) Diff(table *ITable, diff *ITable) string {
	buf := bytes.NewBuffer([]byte{})
	dest := IColumn{}
	for _, item := range diff.Columns {
		if item.Name == p.Name {
			dest = item
		}
	}
	if dest.Name == "" {
		buf.WriteString(p.Create(table))
	} else if p.Create(table) != dest.Create(diff) {
		buf.WriteString(p.SetModify(table))
	}
	return buf.String()
}

func (p *IColumn) GetGoType() string {
	ctypes := strings.Split(p.DataType, "(")
	value := ""
	switch ctypes[0] {
	case "tinyint", "smallint", "mediumint", "int":
		if strings.Contains(p.DataType, "unsigned") {
			value = "uint"
		} else {
			value = "int"
		}
	case "bigint":
		if strings.Contains(p.DataType, "unsigned") {
			value = "uint64"
		} else {
			value = "int64"
		}
	case "float":
		value = "float32"
	case "double":
		value = "float64"
	case "char", "varchar", "text", "tinytext", "midiumtext", "longtext":
		value = "string"
	case "timestamp", "date", "datetime", "time":
		value = "time.Time"
	case "bool":
		value = "bool"
	case "blob", "tinyblob", "midiumblob", "longblob", "binary", "varbinary":
		value = "[]byte"
	default:
		panic(fmt.Sprintf("unknown datatype: %s", p.DataType))
	}
	if !p.Null {
		return value
	}
	return fmt.Sprintf("*%s", value)
}

func (p *IColumn) GetGoTag(table *ITable) string {
	ok := false
	for _, index := range table.Indexes {
		if index.Name == "PRIMARY KEY" {
			if !slices.Contains(index.Columns, p.Name) {
				ok = true
			}
		}
	}
	if ok {
		return ` gorm:"primarykey"`
	}
	switch p.DataType {
	case "_int4":
		return ` gorm:"type:integer[]"`
	case "_varchar", "_macaddr":
		return ` gorm:"type:text[]"`
	}
	return ""
}

func (p *IColumn) GetGoComment() string {
	if p.Comment != nil && len(*p.Comment) > 0 {
		return fmt.Sprintf(" // %s", *p.Comment)
	}
	return ""
}

func (p *IColumn) GetDartType() string {
	ctypes := strings.Split(p.DataType, "(")
	value := ""
	switch ctypes[0] {
	case "tinyint", "smallint", "mediumint", "int":
		value = "int"
	case "bigint":
		value = "int"
	case "float":
		value = "double"
	case "double":
		value = "double"
	case "char", "varchar", "text", "tinytext", "midiumtext", "longtext":
		value = "String"
	case "timestamp":
		value = "int"
	case "date", "datetime", "time":
		value = "String"
	case "bool":
		value = "bool"
	case "blob", "tinyblob", "midiumblob", "longblob", "binary", "varbinary":
		value = "List<int>"
	default:
		panic(fmt.Sprintf("unknown datatype: %s", p.DataType))
	}
	if !p.Null {
		return value
	}
	return fmt.Sprintf("%s?", value)
}
