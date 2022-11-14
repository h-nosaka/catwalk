package postgresql

import (
	"bytes"
	"fmt"
	"strings"

	"gorm.io/gorm"
)

type PGIndex struct {
	Schemaname     string
	Tablename      string
	Indexname      string
	ConstraintType *string  `yaml:"constraint_type,omitempty"`
	Columns        []string `gorm:"->:false"`
}

func (p *PGIndex) GetColumn(db *gorm.DB) {
	db.Raw(`
		SELECT pg_attribute.attname
			FROM pg_catalog.pg_attribute
			LEFT JOIN pg_catalog.pg_class ON pg_attribute.attrelid = pg_class.oid
			WHERE pg_class.relname = ?
			ORDER BY pg_attribute.attnum;
	`, p.Indexname).Scan(&p.Columns)
}

func (p *PGIndex) Create() string {
	if p.ConstraintType != nil {
		switch *p.ConstraintType {
		case "PRIMARY KEY":
			return fmt.Sprintf(
				"ALTER TABLE %s.%s ADD CONSTRAINT %s PRIMARY KEY (%s);\n\n",
				p.Schemaname,
				p.Tablename,
				p.Indexname,
				strings.Join(p.Columns, ","),
			)
		case "UNIQUE":
			return fmt.Sprintf(
				"ALTER TABLE %s.%s ADD CONSTRAINT %s UNIQUE (%s);\n\n",
				p.Schemaname,
				p.Tablename,
				p.Indexname,
				strings.Join(p.Columns, ","),
			)
		default:
			panic(fmt.Sprintf("unknown constraint_type: %s", *p.ConstraintType))
		}
	}
	return fmt.Sprintf(
		"CREATE INDEX %s ON %s.%s (%s);\n\n",
		p.Indexname,
		p.Schemaname,
		p.Tablename,
		strings.Join(p.Columns, ","),
	)
}

func (p *PGIndex) Drop() string {
	if p.ConstraintType != nil {
		switch *p.ConstraintType {
		case "PRIMARY KEY":
			return fmt.Sprintf(
				"ALTER TABLE %s.%s DROP CONSTRAINT %s;\n",
				p.Schemaname,
				p.Tablename,
				p.Indexname,
			)
		case "UNIQUE":
			return fmt.Sprintf(
				"ALTER TABLE %s.%s DROP CONSTRAINT %s;\n",
				p.Schemaname,
				p.Tablename,
				p.Indexname,
			)
		default:
			panic(fmt.Sprintf("unknown constraint_type: %s", *p.ConstraintType))
		}
	}
	return fmt.Sprintf(
		"DROP INDEX IF EXISTS %s.%s RESTRICT;\n",
		p.Schemaname,
		p.Indexname,
	)
}

func (p PGIndex) Diff(src *[]PGIndex) string {
	buf := bytes.NewBuffer([]byte{})
	dest := PGIndex{}
	for _, item := range *src {
		if item.Indexname == p.Indexname {
			dest = item
		}
	}
	if dest.Indexname == "" {
		buf.WriteString(p.Create())
	} else if p.Create() != dest.Create() {
		buf.WriteString(dest.Drop())
		buf.WriteString(p.Create())
	}
	return buf.String()
}
