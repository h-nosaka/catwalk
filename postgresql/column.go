package postgresql

import (
	"bytes"
	"fmt"

	"github.com/freebitdx/fbfiber/context"
)

type PGColumn struct {
	ColumnName             string
	UdtName                string
	ColumnDefault          string
	IsNullable             string
	CharacterMaximumLength int
	NumericPrecision       int
	Comment                *string `yaml:"comment,omitempty"`
	Using                  *string `gorm:"->:false" yaml:"using,omitempty"`
	Rename                 *string `gorm:"->:false" yaml:"rename,omitempty"`
}

func (p *PGColumn) GetColumnType() string {
	switch p.UdtName {
	case "int4":
		if p.NumericPrecision > 0 && p.NumericPrecision < 32 {
			return fmt.Sprintf("integer(%d)", p.NumericPrecision)
		}
		return "integer"
	case "varchar":
		if p.CharacterMaximumLength > 0 {
			return fmt.Sprintf("character varying(%d)", p.CharacterMaximumLength)
		}
		return "character varying"
	case "bpchar":
		if p.CharacterMaximumLength > 0 {
			return fmt.Sprintf("character(%d)", p.CharacterMaximumLength)
		}
		return "character"
	case "timestamp":
		return "timestamp without time zone"
	case "bool":
		return "boolean"
	case "int2":
		return "smallint"
	case "_varchar":
		return "varchar[]"
	case "_int4":
		return "int4[]"
	case "_macaddr":
		return "macaddr[]"
	case "json", "text", "cidr", "inet", "macaddr", "date":
		return p.UdtName
	default:
		panic(fmt.Sprintf("unknown udtname: %s", p.UdtName))
	}
}

func (p *PGColumn) GetDefault() string {
	if p.ColumnDefault != "" {
		return fmt.Sprintf(" DEFAULT %s", p.ColumnDefault)
	}
	return ""
}

func (p *PGColumn) GetNullable() string {
	if p.IsNullable == "NO" {
		return " NOT NULL"
	}
	return ""
}

func (p *PGColumn) Append() string {
	return fmt.Sprintf(
		"\t%s %s%s%s",
		p.ColumnName,
		p.GetColumnType(),
		p.GetDefault(),
		p.GetNullable(),
	)
}

func (p *PGColumn) Create(t *PGTable) string {
	return fmt.Sprintf(
		"ALTER TABLE %s.%s ADD COLUMN %s;\n\n",
		t.Schemaname,
		t.Tablename,
		p.Append(),
	)
}

func (p *PGColumn) Drop(t *PGTable) string {
	return fmt.Sprintf(
		"ALTER TABLE %s.%s DROP COLUMN %s;\n\n",
		t.Schemaname,
		t.Tablename,
		p.ColumnName,
	)
}

func (p *PGColumn) Type(t *PGTable) string {
	if p.Using != nil {
		return fmt.Sprintf(
			"ALTER TABLE %s.%s ALTER COLUMN %s TYPE %s USING %s;\n\n",
			t.Schemaname,
			t.Tablename,
			p.ColumnName,
			p.GetColumnType(),
			*p.Using,
		)
	}
	return fmt.Sprintf(
		"ALTER TABLE %s.%s ALTER COLUMN %s TYPE %s;\n\n",
		t.Schemaname,
		t.Tablename,
		p.ColumnName,
		p.GetColumnType(),
	)
}

func (p *PGColumn) SetNullable(t *PGTable) string {
	value := "DROP NOT NULL"
	if p.IsNullable == "NO" {
		value = "SET NOT NULL"
	}
	return fmt.Sprintf(
		"ALTER TABLE %s.%s ALTER COLUMN %s %s;\n\n",
		t.Schemaname,
		t.Tablename,
		p.ColumnName,
		value,
	)
}

func (p *PGColumn) SetDefault(t *PGTable) string {
	value := "DROP DEFAULT"
	if p.ColumnDefault != "" {
		value = fmt.Sprintf("SET %s", p.GetDefault())
	}
	return fmt.Sprintf(
		"ALTER TABLE %s.%s ALTER COLUMN %s %s;\n\n",
		t.Schemaname,
		t.Tablename,
		p.ColumnName,
		value,
	)
}

func (p *PGColumn) SetComment(t *PGTable) string {
	if p.Comment == nil {
		return ""
	}
	return fmt.Sprintf("COMMENT ON COLUMN %s.%s.%s IS '%s';\n\n", t.Schemaname, t.Tablename, p.ColumnName, *p.Comment)
}

func (p *PGColumn) DropComment(t *PGTable) string {
	return fmt.Sprintf("COMMENT ON COLUMN %s.%s.%s IS NULL;\n\n", t.Schemaname, t.Tablename, p.ColumnName)
}

func (p *PGColumn) RenameColumn(t *PGTable) string {
	if p.Rename != nil {
		return fmt.Sprintf(
			"ALTER TABLE %s.%s RENAME COLUMN %s TO %s;\n",
			t.Schemaname,
			t.Tablename,
			p.ColumnName,
			*p.Rename,
		)
	}
	return ""
}

func (p *PGColumn) IsDrop(src *[]PGColumn, table *PGTable) string {
	buf := bytes.NewBuffer([]byte{})
	dest := PGColumn{}
	for _, item := range *src {
		if item.ColumnName == p.ColumnName {
			dest = item
		}
	}
	if dest.ColumnName == "" {
		buf.WriteString(p.Drop(table))
	}
	return buf.String()
}

func (p PGColumn) Diff(src *[]PGColumn, table *PGTable) string {
	buf := bytes.NewBuffer([]byte{})
	dest := PGColumn{}
	for _, item := range *src {
		if item.ColumnName == p.ColumnName {
			dest = item
		}
	}
	if dest.ColumnName == "" {
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

func (p *PGColumn) GetGoType() string {
	value := ""
	switch p.UdtName {
	case "int4", "int2":
		value = "int"
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
		panic(fmt.Sprintf("unknown udtname: %s", p.UdtName))
	}
	if p.IsNullable == "NO" {
		return value
	}
	return fmt.Sprintf("*%s", value)
}

func (p *PGColumn) GetGoTag(table *PGTable, indexes *[]PGIndex) string {
	ok := false
	for _, index := range *indexes {
		if index.ConstraintType != nil && *index.ConstraintType == "PRIMARY KEY" && index.Schemaname == table.Schemaname && index.Tablename == table.Tablename {
			if !context.ArrayInclude(index.Columns, p.ColumnName) {
				ok = true
			}
		}
	}
	if ok {
		return ` gorm:"primarykey"`
	}
	switch p.UdtName {
	case "_int4":
		return ` gorm:"type:integer[]"`
	case "_varchar", "_macaddr":
		return ` gorm:"type:text[]"`
	}
	return ""
}

func (p *PGColumn) GetGoComment() string {
	if p.Comment != nil && len(*p.Comment) > 0 {
		return fmt.Sprintf(" // %s", *p.Comment)
	}
	return ""
}
