package postgresql

import (
	"bytes"
	"fmt"

	"github.com/gertd/go-pluralize"
)

type PGForeignkey struct {
	TableSchema    string
	TableName      string
	ConstraintName string
	Columnname     string
	Reftable       string
	Refcolumn      string
}

func (p *PGForeignkey) Create() string {
	return fmt.Sprintf(
		"ALTER TABLE ONLY %s.%s ADD CONSTRAINT %s FOREIGN KEY (%s) REFERENCES %s.%s(%s);\n\n",
		p.TableSchema,
		p.TableName,
		p.ConstraintName,
		p.Columnname,
		p.TableSchema,
		p.Reftable,
		p.Refcolumn,
	)
}

func (p *PGForeignkey) Drop() string {
	return fmt.Sprintf(
		"ALTER TABLE %s.%s DROP CONSTRAINT %s;\n",
		p.TableSchema,
		p.TableName,
		p.ConstraintName,
	)
}

func (p PGForeignkey) Diff(src *[]PGForeignkey) string {
	buf := bytes.NewBuffer([]byte{})
	dest := PGForeignkey{}
	for _, item := range *src {
		if item.ConstraintName == p.ConstraintName {
			dest = item
		}
	}
	if dest.ConstraintName == "" {
		buf.WriteString(p.Create())
	} else if p.Create() != dest.Create() {
		buf.WriteString(dest.Drop())
		buf.WriteString(p.Create())
	}
	return buf.String()
}

func (p *PGForeignkey) GetRelation(reverse bool, model string, src *[]Mapping) string {
	for _, m := range *src {
		if reverse {
			if p.Reftable == m.Tablename && p.TableName == m.Reftable && p.Refcolumn == m.Columnname && p.Columnname == m.Refcolumn {
				if m.Association == "MANY" {
					con := pluralize.NewClient()
					return fmt.Sprintf("%s []%s", con.Plural(model), model)
				}
			}
		} else {
			if p.TableName == m.Tablename && p.Reftable == m.Reftable && p.Columnname == m.Columnname && p.Refcolumn == m.Refcolumn {
				if m.Association == "MANY" {
					con := pluralize.NewClient()
					return fmt.Sprintf("%s []%s", con.Plural(model), model)
				}
			}
		}
	}
	return fmt.Sprintf("%s *%s", model, model)
}
