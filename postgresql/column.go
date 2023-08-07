package postgresql

import (
	"bytes"
	"fmt"

	"github.com/h-nosaka/catwalk/base"
	"golang.org/x/exp/slices"
)

type IColumn struct {
	Name     string  `gorm:"column:column_name"`
	DataType string  `gorm:"column:udt_name"`
	Defaults *string `gorm:"column:column_default" yaml:"default,omitempty"`
	Null     bool    `gorm:"->:false" yaml:"nullable,omitempty"`
	IsNull   string  `gorm:"column:is_nullable" yaml:"-"`
	CharaMax int     `gorm:"column:character_maximum_length" yaml:"chara_max,omitempty"`
	NumMax   int     `gorm:"column:numeric_precision" yaml:"num_max,omitempty"`
	Comment  *string `gorm:"column:comment" yaml:"comment,omitempty"`
	Using    *string `gorm:"->:false" yaml:"using,omitempty"`
	Rename   *string `gorm:"->:false" yaml:"rename,omitempty"`
}

func NewColumn(name string, dataType string, charaMax int, numMax int, defaults *string, nullable *bool, comment *string, using *string, rename *string) IColumn {
	null := true
	if nullable != nil {
		null = *nullable
	}
	return IColumn{
		Name:     name,
		DataType: dataType,
		CharaMax: charaMax,
		NumMax:   numMax,
		Defaults: defaults,
		Null:     null,
		Comment:  comment,
		Using:    using,
		Rename:   rename,
	}
}

func DefaultColumn(seq string, cols ...IColumn) []IColumn {
	rs := []IColumn{
		NewColumn("id", "int4", 0, 32, base.String(fmt.Sprintf("nextval('%s'::regclass)", seq)), base.Bool(false), base.String("primary key"), nil, nil),
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
		NewColumn("created_at", "timestamp", 0, 0, base.String("(now())::timestamp(0) without time zone"), nil, base.String("作成日"), nil, nil),
		NewColumn("updated_at", "timestamp", 0, 0, base.String("(now())::timestamp(0) without time zone"), nil, base.String("更新日"), nil, nil),
	}
}

func (p *IColumn) GetColumnType() string {
	switch p.DataType {
	case "int2":
		return "int2"
	case "int4":
		return "int4"
	case "int8":
		return "int8"
	case "varchar":
		if p.CharaMax > 0 {
			return fmt.Sprintf("character varying(%d)", p.CharaMax)
		}
		return "character varying"
	case "bpchar":
		if p.CharaMax > 0 {
			return fmt.Sprintf("character(%d)", p.CharaMax)
		}
		return "character"
	case "timestamp":
		return "timestamp without time zone"
	case "bool":
		return "boolean"
	case "_varchar":
		return "varchar[]"
	case "_int4":
		return "int4[]"
	case "_macaddr":
		return "macaddr[]"
	case "json", "text", "cidr", "inet", "macaddr", "date":
		return p.DataType
	default:
		panic(fmt.Sprintf("unknown udtname: %s", p.DataType))
	}
}

func (p *IColumn) GetDefault() string {
	if p.Defaults != nil && *p.Defaults != "" {
		return fmt.Sprintf(" DEFAULT %s", *p.Defaults)
	}
	return ""
}

func (p *IColumn) GetNullable() string {
	if !p.Null {
		return " NOT NULL"
	}
	return ""
}

func (p *IColumn) Append() string {
	return fmt.Sprintf(
		"\t%s %s%s%s",
		p.Name,
		p.GetColumnType(),
		p.GetDefault(),
		p.GetNullable(),
	)
}

func (p *IColumn) Create(t *ITable) string {
	return fmt.Sprintf(
		"ALTER TABLE %s%s ADD COLUMN %s;\n\n",
		t.SchemaName(),
		t.Name,
		p.Append(),
	)
}

func (p *IColumn) Drop(t *ITable) string {
	return fmt.Sprintf(
		"ALTER TABLE %s%s DROP COLUMN %s;\n\n",
		t.SchemaName(),
		t.Name,
		p.Name,
	)
}

func (p *IColumn) Type(t *ITable) string {
	if p.Using != nil {
		return fmt.Sprintf(
			"ALTER TABLE %s%s ALTER COLUMN %s TYPE %s USING %s;\n\n",
			t.SchemaName(),
			t.Name,
			p.Name,
			p.GetColumnType(),
			*p.Using,
		)
	}
	return fmt.Sprintf(
		"ALTER TABLE %s%s ALTER COLUMN %s TYPE %s;\n\n",
		t.SchemaName(),
		t.Name,
		p.Name,
		p.GetColumnType(),
	)
}

func (p *IColumn) SetNullable(t *ITable) string {
	value := "DROP NOT NULL"
	if !p.Null {
		value = "SET NOT NULL"
	}
	return fmt.Sprintf(
		"ALTER TABLE %s%s ALTER COLUMN %s %s;\n\n",
		t.SchemaName(),
		t.Name,
		p.Name,
		value,
	)
}

func (p *IColumn) SetDefault(t *ITable) string {
	value := "DROP DEFAULT"
	if p.Defaults != nil && *p.Defaults != "" {
		value = fmt.Sprintf("SET %s", p.GetDefault())
	}
	return fmt.Sprintf(
		"ALTER TABLE %s%s ALTER COLUMN %s %s;\n\n",
		t.SchemaName(),
		t.Name,
		p.Name,
		value,
	)
}

func (p *IColumn) SetComment(t *ITable) string {
	if p.Comment == nil {
		return ""
	}
	return fmt.Sprintf("COMMENT ON COLUMN %s%s.%s IS '%s';\n\n", t.SchemaName(), t.Name, p.Name, *p.Comment)
}

func (p *IColumn) DropComment(t *ITable) string {
	return fmt.Sprintf("COMMENT ON COLUMN %s%s.%s IS NULL;\n\n", t.SchemaName(), t.Name, p.Name)
}

func (p *IColumn) RenameColumn(t *ITable) string {
	if p.Rename != nil {
		return fmt.Sprintf(
			"ALTER TABLE %s%s RENAME COLUMN %s TO %s;\n",
			t.SchemaName(),
			t.Name,
			p.Name,
			*p.Rename,
		)
	}
	return ""
}

func (p *IColumn) IsDrop(src *[]IColumn, table *ITable) string {
	buf := bytes.NewBuffer([]byte{})
	dest := IColumn{}
	for _, item := range *src {
		if item.Name == p.Name {
			dest = item
		}
	}
	if dest.Name == "" {
		buf.WriteString(p.Drop(table))
	}
	return buf.String()
}

func (p IColumn) Diff(src *[]IColumn, table *ITable) string {
	buf := bytes.NewBuffer([]byte{})
	dest := IColumn{}
	for _, item := range *src {
		if item.Name == p.Name {
			dest = item
		}
	}
	if dest.Name == "" {
		buf.WriteString(p.Create(table))
	} else if p.Create(table) != dest.Create(table) {
		if p.SetNullable(table) != dest.SetNullable(table) {
			buf.WriteString(p.SetNullable(table))
		}
		if p.SetDefault(table) != dest.SetDefault(table) {
			buf.WriteString(p.SetDefault(table))
		}
		if p.Type(table) != dest.Type(table) {
			buf.WriteString(p.Type(table))
		}
	}
	if p.SetComment(table) != dest.SetComment(table) {
		if p.Comment != nil {
			buf.WriteString(p.SetComment(table))
		} else {
			buf.WriteString(p.DropComment(table))
		}
	}
	return buf.String()
}

func (p *IColumn) GetGoType() string {
	value := ""
	switch p.DataType {
	case "int4", "int2":
		value = "int"
	case "int8":
		value = "int64"
	case "varchar", "bpchar", "json", "text", "macaddr":
		value = "string"
	case "timestamp", "date":
		value = "time.Time"
	case "bool":
		value = "bool"
	case "_varchar", "_macaddr":
		value = "pq.StringArray"
	case "_int4":
		value = "pq.Int32Array"
	case "cidr", "inet":
		// value = "pgtype.Inet"
		value = "string" // TODO: Inet型だとSQLがうまくいかないケースがあるようなのでいったんString
	default:
		panic(fmt.Sprintf("unknown udtname: %s", p.DataType))
	}
	if !p.Null {
		return value
	}
	return fmt.Sprintf("*%s", value)
}

func (p *IColumn) GetGoTag(table *ITable) string {
	ok := false
	for _, index := range table.Indexes {
		if index.ConstraintType != nil && *index.ConstraintType == "PRIMARY KEY" {
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
