package postgresql

import (
	"bytes"
	"fmt"
	"regexp"

	"github.com/gertd/go-pluralize"
	"github.com/iancoleman/strcase"
)

type IForeignkey struct {
	Name      string `gorm:"column:constraint_name"`
	Column    string `gorm:"column:columnname"`
	RefTable  string `gorm:"column:reftable"`
	RefColumn string `gorm:"column:refcolumn"`
	HasOne    bool   `gorm:"->:false"`
	HasAny    bool   `gorm:"->:false"`
}

type IRelation struct {
	Column    string
	RefTable  string
	RefColumn string
	HasOne    bool
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

func (p *IForeignkey) Create(t *ITable) string {
	return fmt.Sprintf(
		"ALTER TABLE ONLY %s%s ADD CONSTRAINT %s FOREIGN KEY (%s) REFERENCES %s%s(%s);\n\n",
		t.SchemaName(),
		t.Name,
		p.Name,
		p.Column,
		t.SchemaName(),
		p.RefTable,
		p.RefColumn,
	)
}

func (p *IForeignkey) Drop(t *ITable) string {
	return fmt.Sprintf(
		"ALTER TABLE %s%s DROP CONSTRAINT %s;\n",
		t.SchemaName(),
		t.Name,
		p.Name,
	)
}

func (p IForeignkey) Diff(t *ITable, src *[]IForeignkey) string {
	buf := bytes.NewBuffer([]byte{})
	dest := IForeignkey{}
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

func (p *IForeignkey) GetReference() string {
	if p.HasAny {
		return fmt.Sprintf("foreignKey:%s;references:%s", strcase.ToCamel(p.RefColumn), strcase.ToCamel(p.Column))
	}
	return fmt.Sprintf("foreignKey:%s;references:%s", strcase.ToCamel(p.Column), strcase.ToCamel(p.RefColumn))
}

func (p *IForeignkey) GetRelation() string {
	con := pluralize.NewClient()
	model := strcase.ToCamel(con.Singular(p.RefTable))
	// key := strcase.ToCamel(strings.ReplaceAll(p.Column, "_id", ""))
	key := strcase.ToCamel(regexp.MustCompile("_id|Id$").ReplaceAllString(p.Column, ""))
	if key == "Id" || key == "" {
		key = model
	}
	if !p.HasOne {
		return fmt.Sprintf("%s []%s", con.Plural(key), model)
	}
	return fmt.Sprintf("%s *%s", key, model)
}

func (p *IRelation) GetReference(t *ITable) string {
	con := pluralize.NewClient()
	fmt.Println(strcase.ToCamel(p.Column), strcase.ToCamel(p.RefColumn), strcase.ToCamel(fmt.Sprintf("%s_id", con.Singular(t.Name))))
	if !p.HasOne && strcase.ToCamel(p.RefColumn) == strcase.ToCamel(fmt.Sprintf("%s_id", con.Singular(t.Name))) {
		if strcase.ToCamel(p.Column) == "Id" {
			return ""
		} else {
			return fmt.Sprintf("foreignKey:%s", strcase.ToCamel(p.Column))
		}
	}
	return fmt.Sprintf("foreignKey:%s;references:%s", strcase.ToCamel(p.Column), strcase.ToCamel(p.RefColumn))
}

func (p *IRelation) GetRelation() string {
	con := pluralize.NewClient()
	model := strcase.ToCamel(con.Singular(p.RefTable))
	// key := strcase.ToCamel(strings.ReplaceAll(p.Column, "_id", ""))
	key := strcase.ToCamel(regexp.MustCompile("_id|Id$").ReplaceAllString(p.Column, ""))
	if key == "Id" || key == "" {
		key = model
	}
	if !p.HasOne {
		return fmt.Sprintf("%s []%s", con.Plural(key), model)
	}
	return fmt.Sprintf("%s *%s", key, model)
}
