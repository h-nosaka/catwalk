package catwalk

import (
	"bytes"
	"fmt"
	"strings"

	"github.com/h-nosaka/catwalk/base"
)

func (p *Table) SchemaTable() string {
	switch GetSchemaMode() {
	case SchemaModeMySQL:
		return fmt.Sprintf("`%s`.`%s`", p.Schema, p.Name)
	case SchemaModePostgres:
		return fmt.Sprintf(`"%s"."%s"`, p.Schema, p.Name)
	default:
		panic("未定義のデータベース")
	}
}

func (p *Table) sqlTableDrop() string {
	return GetParts("TableDrop", p.SchemaTable())
}

func (p *Table) sqlTableRename() string {
	return GetParts("TableRename", p.SchemaTable(), p.Rename)
}

func (p *Table) sqlTableCreateOpen() string {
	return GetParts("TableCreateOpen", p.SchemaTable())
}

func (p *Table) sqlTableCreateClose() string {
	if p.Comment == nil {
		return GetParts("TableCreateClose")
	}
	values := []interface{}{}
	switch GetSchemaMode() {
	case SchemaModeMySQL:
		values = append(values, *p.Comment)
	case SchemaModePostgres:
		values = append(values, p.SchemaTable(), *p.Comment)
	default:
		panic("未定義のデータベース")
	}
	return GetParts("TableCreateCloseWithComment", values...)
}

func (p *Table) sqlColumnSetComment() string {
	rs := []string{}
	mode := GetSchemaMode()
	for _, col := range p.Columns {
		if col.Comment != nil {
			values := []interface{}{}
			if mode == SchemaModePostgres {
				values = append(values, p.SchemaTable(), col.Name, *col.Comment)
				rs = append(rs, GetParts("ColumnSetComment", values...))
			}
		}
	}
	return strings.Join(rs, "")
}

func (p *Table) Create() string {
	buf := bytes.NewBuffer([]byte{})
	buf.WriteString(p.sqlTableCreateOpen())
	cols := []string{}
	for _, item := range p.Columns {
		cols = append(cols, item.Append())
	}
	buf.WriteString(fmt.Sprintf("\t%s\n", strings.Join(cols, ",\n\t")))
	// for _, index := range p.Indexes {
	// 	buf.WriteString(index.Append())
	// }
	buf.WriteString(p.sqlTableCreateClose())
	buf.WriteString(p.sqlColumnSetComment())
	indexes := []string{}
	for _, item := range p.Indexes {
		indexes = append(indexes, item.Create(p))
	}
	buf.WriteString(fmt.Sprintf("%s\n", strings.Join(indexes, "")))
	relations := []string{}
	for _, item := range p.Relations {
		relations = append(relations, item.Create(p))
	}
	buf.WriteString(fmt.Sprintf("%s\n", strings.Join(relations, "")))
	if p.Partitions != nil {
		buf.WriteString(p.Partitions.Create(p))
	}
	return buf.String()
}

func (p *Table) Drop() string {
	return p.sqlTableDrop()
}

func (p *Table) RenameTable() string {
	return p.sqlTableRename()
}

func (p *Table) Diff(src *[]Table) string {
	buf := bytes.NewBuffer([]byte{})
	var dest Table
	for _, item := range *src {
		if item.Name == p.Name {
			dest = item
		}
	}
	if dest.Name == "" {
		buf.WriteString(p.Create())
	} else if p.Create() != dest.Create() {
		if base.GetEnvBool("DEBUG", false) {
			fmt.Println("new:", p.Create())
			fmt.Println("old:", dest.Create())
		}
		// カラムの差分
		for _, item := range p.Columns {
			buf.WriteString(item.Diff(p, &dest))
		}
		for _, item := range dest.Columns {
			buf.WriteString(item.IsDrop(p))
		}
		// インデックスの差分
		for _, item := range p.Indexes {
			buf.WriteString(item.Diff(p, &dest))
		}
		for _, item := range dest.Indexes {
			buf.WriteString(item.IsDrop(p))
		}
		// 制約の差分
		for _, item := range p.Relations {
			if item.IsForeignKey {
				buf.WriteString(item.Diff(p, &dest))
			}
		}
		for _, item := range dest.Relations {
			if item.IsForeignKey {
				buf.WriteString(item.IsDrop(p))
			}
		}
		// パーティション
		if p.Partitions != nil {
			buf.WriteString(p.Partitions.Diff(p, dest.Partitions))
		}
		// カラムのリネーム
		for _, col := range p.Columns {
			if col.Rename != nil && col.Name != *col.Rename {
				buf.WriteString(col.RenameColumn(p))
			}
		}
	}
	// テーブルのリネーム
	if p.Rename != nil && p.Name != *p.Rename {
		buf.WriteString(p.RenameTable())
	}
	return buf.String()
}
