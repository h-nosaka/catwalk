package mysql

import (
	"bytes"
	"fmt"

	"github.com/freebitdx/fbfiber/context"
	"github.com/gertd/go-pluralize"
	"github.com/iancoleman/strcase"
)

type IForeignkey struct {
	Name      string `gorm:"column:CONSTRAINT_NAME"`
	Column    string `gorm:"column:COLUMN_NAME"`
	RefTable  string `gorm:"column:REFERENCED_TABLE_NAME"`
	RefColumn string `gorm:"column:REFERENCED_COLUMN_NAME"`
	HasOne    bool   `gorm:"->:false"`
	HasAny    bool   `gorm:"->:false"`
}

func NewFK(name string, column string, reftable string, refcolumn string, hasone bool, any bool) IForeignkey {
	return IForeignkey{
		Name:      name,
		Column:    column,
		RefTable:  reftable,
		RefColumn: refcolumn,
		HasOne:    hasone,
		HasAny:    any,
	}
}

func (p *IForeignkey) Create(table *ITable) string {
	// 外部キーはいったん設定しない方針にする
	return ""
	// if p.HasAny {
	// 	return ""
	// }
	// return fmt.Sprintf(
	// 	"ALTER TABLE %s ADD CONSTRAINT %s FOREIGN KEY %s(%s) REFERENCES %s(%s);\n\n",
	// 	table.Name,
	// 	p.Name,
	// 	p.RefTable,
	// 	p.RefColumn,
	// 	table.Name,
	// 	p.Column,
	// )
}

func (p *IForeignkey) Drop(table *ITable) string {
	return fmt.Sprintf(
		"ALTER TABLE %s DROP FOREIGN KEY %s;\n",
		table.Name,
		p.Name,
	)
}

func (p IForeignkey) Diff(table *ITable, diff *ITable) string {
	buf := bytes.NewBuffer([]byte{})
	dest := IForeignkey{}
	for _, item := range diff.Foreignkeys {
		if item.Name == p.Name {
			dest = item
		}
	}
	if dest.Name == "" {
		buf.WriteString(p.Create(table))
	} else if p.Create(table) != dest.Create(diff) {
		context.Ctx.Logger.Debug(p.Create(table))
		context.Ctx.Logger.Debug(dest.Create(diff))
		buf.WriteString(dest.Drop(diff))
		buf.WriteString(p.Create(table))
	}
	return buf.String()
}

func (p *IForeignkey) GetReference() string {
	if p.HasAny {
		return fmt.Sprintf("foreignKey:%s;references:%s", strcase.ToCamel(p.RefColumn), strcase.ToCamel(p.Column))
	}
	return fmt.Sprintf("foreignKey:%s;references:%s", strcase.ToCamel(p.Column), strcase.ToCamel(p.RefColumn))
}

func (p *IForeignkey) GetRelation() string {
	con := pluralize.NewClient()
	model := strcase.ToCamel(con.Singular(p.RefTable))
	if p.HasOne {
		return fmt.Sprintf("%s []%s", strcase.ToCamel(p.RefTable), model)
	}
	return fmt.Sprintf("%s *%s", model, model)
}

func (p *IForeignkey) GetDartRelation() string {
	con := pluralize.NewClient()
	model := strcase.ToCamel(con.Singular(p.RefTable))
	if p.HasOne {
		return fmt.Sprintf("late List<%s>? %s;", model, strcase.ToLowerCamel(p.RefTable))
	}
	return fmt.Sprintf("late %s? %s;", model, strcase.ToLowerCamel(con.Singular(p.RefTable)))
}
