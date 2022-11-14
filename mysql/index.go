package mysql

import (
	"bytes"
	"fmt"
	"strings"

	"gorm.io/gorm"
)

type IIndex struct {
	Name      string   `gorm:"column:INDEX_NAME"`
	Unique    bool     `gorm:"->:false"`
	NonUnique int      `gorm:"column:NON_UNIQUE" yaml:"-"`
	Columns   []string `gorm:"->:false"`
}

func NewIndex(name string, unique *bool, columns ...string) IIndex {
	uk := true
	if unique != nil {
		uk = *unique
	}
	return IIndex{
		Name:    name,
		Unique:  uk,
		Columns: columns,
	}
}

func DefaultIndex(indexes ...IIndex) []IIndex {
	rs := []IIndex{
		NewIndex("PRIMARY", nil, "id"),
	}
	rs = append(rs, indexes...)
	return rs
}

func (p *IIndex) GetColumn(db *gorm.DB, table *ITable) {
	db.Raw(fmt.Sprintf(GetIndexColumn, table.Schema, table.Name, p.Name)).Scan(&p.Columns)
}

func (p *IIndex) Append() string {
	if p.Name == "PRIMARY" {
		return fmt.Sprintf(
			"\tPRIMARY KEY (`%s`)",
			strings.Join(p.Columns, "`,`"),
		)
	}
	key := "KEY"
	if p.Unique {
		key = "UNIQUE KEY"
	}
	return fmt.Sprintf(
		"\t%s `%s` (`%s`) USING BTREE",
		key,
		p.Name,
		strings.Join(p.Columns, "`,`"),
	)
}

func (p *IIndex) Create(table *ITable) string {
	if p.Name == "PRIMARY" {
		return ""
		// return fmt.Sprintf(
		// 	"ALTER TABLE %s ADD CONSTRAINT %s PRIMARY KEY (%s);\n\n",
		// 	table.Name,
		// 	p.Name,
		// 	strings.Join(p.Columns, ","),
		// )
	}
	if p.Unique {
		return fmt.Sprintf(
			"CREATE UNIQUE INDEX %s USING BTREE ON %s (`%s`);\n\n",
			p.Name,
			table.Name,
			strings.Join(p.Columns, ","),
		)
	}
	return fmt.Sprintf(
		"CREATE INDEX %s USING BTREE ON %s (%s);\n\n",
		p.Name,
		table.Name,
		strings.Join(p.Columns, ","),
	)
}

func (p *IIndex) Drop(table *ITable) string {
	if p.Name == "PRIMARY" {
		return fmt.Sprintf(
			"ALTER TABLE %s DROP PRIMARY KEY %s;\n",
			table.Name,
			p.Name,
		)
	}
	return fmt.Sprintf(
		"ALTER TABLE %s DROP INDEX %s;\n",
		table.Name,
		p.Name,
	)
}

func (p IIndex) Diff(table *ITable, diff *ITable) string {
	buf := bytes.NewBuffer([]byte{})
	dest := IIndex{}
	for _, item := range diff.Indexes {
		if item.Name == p.Name {
			dest = item
		}
	}
	if dest.Name == "" {
		buf.WriteString(p.Create(table))
	} else if p.Create(table) != dest.Create(diff) {
		buf.WriteString(dest.Drop(diff))
		buf.WriteString(p.Create(table))
	}
	return buf.String()
}
