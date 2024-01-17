package catwalk

import (
	"bytes"
	"fmt"
	"strings"
)

func (p *Column) Append() string {
	cols := []string{}
	mode := GetSchemaMode()
	switch mode {
	case SchemaModeMySQL:
		cols = append(cols, p.DataType.Mysql(p.Count))
	case SchemaModePostgres:
		cols = append(cols, p.DataType.Postgres(p.Count))
	}
	if p.Default != nil && *p.Default != "" && *p.Default != "NULL" {
		cols = append(cols, fmt.Sprintf("DEFAULT %s", *p.Default))
	}
	if p.Extra != nil && *p.Extra != "" {
		cols = append(cols, *p.Extra)
	}
	if !p.Null {
		cols = append(cols, "NOT NULL")
	} else {
		cols = append(cols, "NULL")
	}
	switch mode {
	case SchemaModeMySQL:
		if p.Comment != nil && *p.Comment != "" {
			cols = append(cols, fmt.Sprintf("COMMENT '%s'", *p.Comment))
		}
	}
	return GetParts("ColumnAppend", p.Name, strings.Join(cols, " "))
}

func (p *Column) Create(t *Table) string {
	return GetParts("ColumnCreate", t.SchemaTable(), p.Append())
}

func (p *Column) Diff(t *Table, diff *Table) string {
	buf := bytes.NewBuffer([]byte{})
	var dest Column
	for _, item := range diff.Columns {
		if item.Name == p.Name {
			dest = item
		}
	}
	if dest.Name == "" {
		buf.WriteString(p.Create(t))
	} else if p.Create(t) != dest.Create(diff) {
		// if base.GetEnvBool("DEBUG", false) {
		// 	fmt.Println(base.Diff(p.Create(t), dest.Create(diff)))
		// }
		buf.WriteString(p.SetModify(t, &dest))
	}
	return buf.String()
}

func (p *Column) valueSetOrDrop(ok bool) string {
	if ok {
		return "SET"
	}
	return "DROP"
}

func (p *Column) SetModify(t *Table, d *Column) string {
	switch GetSchemaMode() {
	case SchemaModeMySQL:
		return GetParts("ColumnSetModify", t.SchemaTable(), p.Append())
	case SchemaModePostgres:
		rs := []string{}
		if val := GetParts("ColumnSetModify", t.SchemaTable(), p.Name, "Type", p.DataType.Postgres()); val != GetParts("ColumnSetModify", d.Name, t.SchemaTable(), "Type", d.DataType.Postgres()) {
			rs = append(rs, val)
		}
		if val := GetParts("ColumnSetModify", t.SchemaTable(), p.Name, p.valueSetOrDrop(!p.Null), "NOT NULL"); val != GetParts("ColumnSetModify", d.Name, t.SchemaTable(), d.valueSetOrDrop(!d.Null), "NOT NULL") {
			rs = append(rs, val)
		}
		val := "DEFAULT"
		if p.Default != nil {
			val = fmt.Sprintf("DEFAULT %s", *p.Default)
		}
		dest := "DEFAULT"
		if d.Default != nil {
			val = fmt.Sprintf("DEFAULT %s", *d.Default)
		}
		if val := GetParts("ColumnSetModify", t.SchemaTable(), p.Name, p.valueSetOrDrop(p.Default != nil), val); val != GetParts("ColumnSetModify", p.Name, t.SchemaTable(), p.valueSetOrDrop(d.Default != nil), dest) {
			rs = append(rs, val)
		}
		return strings.Join(rs, "")
	}
	return ""
}

func (p *Column) RenameColumn(t *Table) string {
	if p.Rename != nil {
		switch GetSchemaMode() {
		case SchemaModeMySQL:
			return GetParts("ColumnRename", t.SchemaTable(), GetParts("ColumnAppend", strings.Join([]string{p.Name, *p.Rename}, "` `")))
		case SchemaModePostgres:
			return GetParts("ColumnRename", t.SchemaTable(), p.Name, *p.Rename)
		}
	}
	return ""
}

func (p *Column) Drop(t *Table) string {
	return GetParts("ColumnDrop", t.SchemaTable(), p.Name)
}

func (p *Column) IsDrop(t *Table) string {
	var dest *Column
	for _, item := range t.Columns {
		if item.Name == p.Name {
			dest = &item
		}
	}
	if dest == nil {
		return p.Drop(t)
	}
	return ""
}
