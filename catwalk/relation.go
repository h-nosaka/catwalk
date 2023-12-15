package catwalk

import (
	"bytes"
	"fmt"
)

func (p *Relation) SchemaRefTable(t *Table) string {
	switch GetSchemaMode() {
	case SchemaModeMySQL:
		return fmt.Sprintf("`%s`.`%s`", t.Schema, p.RefTable)
	case SchemaModePostgres:
		return fmt.Sprintf(`"%s"."%s"`, t.Schema, p.RefTable)
	default:
		panic("未定義のデータベース")
	}
}

func (p *Relation) Create(t *Table) string {
	if p.IsForeignKey {
		switch GetSchemaMode() {
		case SchemaModeMySQL:
			return GetParts("RelationCreate", t.SchemaTable(), p.Name, t.SchemaTable(), p.Column, p.SchemaRefTable(t), p.RefColumn)
		case SchemaModePostgres:
			return GetParts("RelationCreate", t.SchemaTable(), p.Name, p.Column, p.SchemaRefTable(t), p.RefColumn)
		}
	}
	return ""
}

func (p *Relation) Drop(t *Table) string {
	if p.IsForeignKey {
		return GetParts("RelationDrop", t.SchemaTable(), p.Name)
	}
	return ""
}

func (p *Relation) Diff(t *Table, diff *Table) string {
	buf := bytes.NewBuffer([]byte{})
	var dest Relation
	for _, item := range diff.Relations {
		if item.IsForeignKey {
			if item.Name == p.Name {
				dest = item
			}
		}
	}
	if dest.Name == "" {
		buf.WriteString(p.Create(t))
	} else if p.Create(t) != dest.Create(diff) {
		buf.WriteString(dest.Drop(diff))
		buf.WriteString(p.Create(t))
	}
	return buf.String()
}

func (p *Relation) IsDrop(t *Table) string {
	var dest *Relation
	for _, item := range t.Relations {
		if item.Name == p.Name {
			dest = &item
		}
	}
	if dest == nil {
		return p.Drop(t)
	}
	return ""
}
