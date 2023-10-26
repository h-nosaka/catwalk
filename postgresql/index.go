package postgresql

import (
	"bytes"
	"fmt"
	"strings"

	"github.com/h-nosaka/catwalk/base"
)

type IIndex struct {
	Name           string   `gorm:"column:indexname"`
	ConstraintType *string  `yaml:"constraint_type,omitempty"`
	Columns        []string `gorm:"->:false"`
}

func NewIndex(name string, constraint *string, columns ...string) IIndex {
	return IIndex{
		Name:           name,
		ConstraintType: constraint,
		Columns:        columns,
	}
}

func DefaultIndex(name string, indexes ...IIndex) []IIndex {
	rs := []IIndex{
		NewIndex(name, base.String("PRIMARY KEY"), "id"),
	}
	rs = append(rs, indexes...)
	return rs
}

func (p *IIndex) GetColumn() {
	base.DB.Raw(GetIndexColumn, p.Name).Scan(&p.Columns)
}

func (p *IIndex) Create(t *ITable) string {
	if p.ConstraintType != nil {
		switch *p.ConstraintType {
		case "PRIMARY KEY":
			return fmt.Sprintf(
				"ALTER TABLE %s\"%s\" ADD CONSTRAINT %s PRIMARY KEY (%s);\n\n",
				t.SchemaName(),
				t.Name,
				p.Name,
				strings.Join(p.Columns, ","),
			)
		case "UNIQUE":
			return fmt.Sprintf(
				"ALTER TABLE %s\"%s\" ADD CONSTRAINT %s UNIQUE (%s);\n\n",
				t.SchemaName(),
				t.Name,
				p.Name,
				strings.Join(p.Columns, ","),
			)
		default:
			panic(fmt.Sprintf("unknown constraint_type: %s", *p.ConstraintType))
		}
	}
	return fmt.Sprintf(
		"CREATE INDEX %s ON %s\"%s\" (%s);\n\n",
		p.Name,
		t.SchemaName(),
		t.Name,
		strings.Join(p.Columns, ","),
	)
}

func (p *IIndex) Drop(t *ITable) string {
	if p.ConstraintType != nil {
		switch *p.ConstraintType {
		case "PRIMARY KEY":
			return fmt.Sprintf(
				"ALTER TABLE %s\"%s\" DROP CONSTRAINT %s;\n",
				t.SchemaName(),
				t.Name,
				p.Name,
			)
		case "UNIQUE":
			return fmt.Sprintf(
				"ALTER TABLE %s\"%s\" DROP CONSTRAINT %s;\n",
				t.SchemaName(),
				t.Name,
				p.Name,
			)
		default:
			panic(fmt.Sprintf("unknown constraint_type: %s", *p.ConstraintType))
		}
	}
	return fmt.Sprintf(
		"DROP INDEX IF EXISTS %s\"%s\" RESTRICT;\n",
		t.SchemaName(),
		p.Name,
	)
}

func (p *IIndex) Diff(t *ITable, src *[]IIndex) string {
	buf := bytes.NewBuffer([]byte{})
	dest := IIndex{}
	for _, item := range *src {
		if item.Name == p.Name {
			dest = item
		}
	}
	if dest.Name == "" {
		buf.WriteString(p.Create(t))
	} else if p.Create(t) != dest.Create(t) {
		buf.WriteString(dest.Drop(t))
		buf.WriteString(p.Create(t))
	}
	return buf.String()
}
