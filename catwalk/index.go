package catwalk

import (
	"bytes"
	"fmt"
	"strings"
)

func (p *Index) SchemaIndex(schema string) string {
	switch GetSchemaMode() {
	case SchemaModeMySQL:
		return fmt.Sprintf("`%s`.`%s`", schema, p.Name)
	case SchemaModePostgres:
		return fmt.Sprintf(`"%s"."%s"`, schema, p.Name)
	default:
		panic("未定義のデータベース")
	}
}

func (p *Index) Append() string {
	if p.Name != "PRIMARY" {
		return ""
	}
	return fmt.Sprintf(
		"\tPRIMARY KEY (`%s`)\n",
		strings.Join(p.Columns, "`,`"),
	)
}

func (p *Index) Create(t *Table) string {
	switch GetSchemaMode() {
	case SchemaModeMySQL:
		switch p.Type {
		case "PRIMARY KEY":
			return GetParts("IndexCreatePrimary", p.Name, t.SchemaTable(), strings.Join(p.Columns, "`,`"))
		case "UNIQUE":
			return GetParts("IndexCreateUnique", p.Name, t.SchemaTable(), strings.Join(p.Columns, "`,`"))
		default:
			return GetParts("IndexCreate", p.Name, t.SchemaTable(), strings.Join(p.Columns, "`,`"))
		}
	case SchemaModePostgres:
		switch p.Type {
		case "PRIMARY KEY":
			return GetParts("IndexCreatePrimary", t.SchemaTable(), p.Name, strings.Join(p.Columns, "\",\""))
		case "UNIQUE":
			return GetParts("IndexCreateUnique", t.SchemaTable(), p.Name, strings.Join(p.Columns, "\",\""))
		default:
			return GetParts("IndexCreate", p.Name, t.SchemaTable(), strings.Join(p.Columns, "\",\""))
		}
	}
	return ""
}

func (p *Index) Drop(t *Table) string {
	switch GetSchemaMode() {
	case SchemaModeMySQL:
		switch p.Type {
		case "PRIMARY KEY":
			return GetParts("IndexDropPrimary", t.Name, p.Name)
		case "UNIQUE":
			return GetParts("IndexDropUnique", t.Name, p.Name)
		default:
			return GetParts("IndexDrop", t.Name, p.Name)
		}
	case SchemaModePostgres:
		switch p.Type {
		case "PRIMARY KEY":
			return GetParts("IndexDropPrimary", t.SchemaTable(), p.Name)
		case "UNIQUE":
			return GetParts("IndexDropUnique", t.SchemaTable(), p.Name)
		default:
			return GetParts("IndexDrop", p.SchemaIndex(t.Schema))
		}
	}
	return ""
}

func (p *Index) Diff(t *Table, diff *Table) string {
	buf := bytes.NewBuffer([]byte{})
	var dest Index
	for _, item := range diff.Indexes {
		if item.Name == p.Name {
			dest = item
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

func (p *Index) IsDrop(t *Table) string {
	var dest *Index
	for _, item := range t.Indexes {
		if item.Name == p.Name {
			dest = &item
		}
	}
	if dest == nil {
		return p.Drop(t)
	}
	return ""
}
