package catwalk

import (
	"bytes"
	"fmt"
	"os"
	"strings"

	"github.com/gertd/go-pluralize"
	"github.com/iancoleman/strcase"
	"golang.org/x/exp/slices"
)

func (p *Table) CreateGoModel(path string) bool {
	buf := bytes.NewBuffer([]byte{})
	buf.WriteString("package models\n\n")
	// インポート定義
	imports := []string{`"time"`, `"gorm.io/gorm"`}
	for _, item := range p.Columns {
		switch item.DataType {
		case DataTypeJson:
			if !slices.Contains(imports, `"encoding/json"`) {
				imports = append(imports, `"encoding/json"`)
			}
		}
	}
	for _, item := range p.Enums {
		if item.Type == EnumTypeUint && !slices.Contains(imports, `"encoding/json"`) {
			imports = append(imports, `"encoding/json"`)
		}
	}

	buf.WriteString(fmt.Sprintf("import (\n\t%s\n)\n\n", strings.Join(imports, "\n\t")))
	// コメント処理
	if p.Comment != nil {
		buf.WriteString(fmt.Sprintf("// %s\n", *p.Comment))
	}
	con := pluralize.NewClient()
	table := con.Singular(p.Name)
	// Enumの定義
	for _, enum := range p.Enums {
		buf.WriteString(enum.GoCreate(p))
	}
	// 構造体定義
	buf.WriteString(fmt.Sprintf("type %s struct {\n", strcase.ToCamel(table)))
	// テーブルカラムの定義
	buf.WriteString("\t// column\n")
	for _, item := range p.Columns {
		data := item.DataType.String()
		// data := item.GetGoType()
		for _, enum := range p.Enums {
			if item.Name == enum.Column {
				data = enum.GoStructName(p)
			}
		}
		buf.WriteString(fmt.Sprintf(
			"\t%s %s `json:\"%s\"%s`%s\n",
			strcase.ToCamel(item.Name),
			data,
			item.Name,
			item.GetGoTag(p),    // gorm用タグの設定
			item.GetGoComment(), // コメント処理
		))
	}
	buf.WriteString("\n")
	// リレーションの定義
	buf.WriteString("\t// relation\n")
	for _, item := range p.Relations {
		buf.WriteString(fmt.Sprintf(
			"\t%s `gorm:\"%s\"`\n",
			item.GetRelation(),
			item.GetReference(p),
		))
	}
	buf.WriteString("}\n\n")
	// テーブル名が複数形でない場合の処理
	if con.IsSingular(p.Name) {
		buf.WriteString(fmt.Sprintf("func (c *%s) TableName() string {\n", strcase.ToCamel(table)))
		buf.WriteString(fmt.Sprintf("\treturn \"%s\"\n", p.Name))
		buf.WriteString("}\n\n")
	}
	// json型のmap変換メソッド
	for _, item := range p.Columns {
		if item.DataType == DataTypeJson {
			buf.WriteString(fmt.Sprintf("func (c *%s) %sMap() map[string]interface{} {\n", strcase.ToCamel(table), strcase.ToCamel(item.Name)))
			buf.WriteString("\tvar data map[string]interface{}\n")
			buf.WriteString(fmt.Sprintf("\t\tif err := json.Unmarshal([]byte(c.%s), &data); err != nil {\n", strcase.ToCamel(item.Name)))
			buf.WriteString("\t\t\treturn nil\n")
			buf.WriteString("\t\t}\n")
			buf.WriteString("\treturn data\n")
			buf.WriteString("}\n\n")
		}
	}

	// ファイル出力
	filename := fmt.Sprintf("%s/%s.go", path, table)
	fp, err := os.Create(filename)
	if err != nil {
		fmt.Println(err)
		return false
	}
	defer fp.Close()
	if _, err := fp.Write(buf.Bytes()); err != nil {
		fmt.Printf("WriteString Error: %s\n", err.Error())
		return false
	}
	return true
}

func (p *Table) CreateGoFixture(path string) bool {
	buf := bytes.NewBufferString("package fixtures\n\n")
	con := pluralize.NewClient()
	table := con.Singular(p.Name)
	name := strcase.ToCamel(table)
	uid := ""
	for _, col := range p.Columns {
		if col.Name == "id" && col.DataType == DataTypeUUID {
			uid = "\t\tId: uuid.NewString(),"
		}
	}
	// 雛形
	buf.WriteString(fmt.Sprintf(`import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

func %s(setters ...func(model *models.%s)) *models.%s {
	model := &models.%s{
%s	}
	for _, setter := range setters {
		setter(model)
	}
	return model
}

func Create%s(db *gorm.DB, setters ...func(model *models.%s)) *models.%s {
	model := %s(setters...)
	if err := db.Create(model).Error; err != nil {
		return nil
	}
	return model
}`, name, name, name, name, uid, name, name, name, name))
	// ファイル出力
	filename := fmt.Sprintf("%s/%s.go", path, table)
	if _, err := os.Stat(filename); err != nil { // ファイルの上書きはしない
		fp, err := os.Create(filename)
		if err != nil {
			fmt.Println(err)
			return false
		}
		defer fp.Close()
		if _, err := fp.Write(buf.Bytes()); err != nil {
			fmt.Printf("WriteString Error: %s\n", err.Error())
			return false
		}
	}
	return true
}

func (p *Table) CreateSchemaFile(path string) bool {
	// con := pluralize.NewClient()
	buf := bytes.NewBuffer([]byte{})
	buf.WriteString(`package schema

import (
	"github.com/h-nosaka/catwalk"
	"github.com/h-nosaka/catwalk/base"
)

`)
	buf.WriteString(fmt.Sprintf("func %s() catwalk.Table {\n", strcase.ToCamel(p.Name)))
	comment := ""
	if p.Comment != nil {
		comment = *p.Comment
	}
	buf.WriteString(fmt.Sprintf("return catwalk.NewTable(\"%s\", \"%s\", catwalk.JsonCaseSnake, \"%s\").SetColumns(\n", p.Schema, p.Name, comment))
	for _, col := range p.Columns {
		// catwalk.NewColumn("price", catwalk.DataTypeString, 32, false, "価格").Done(),
		comment := ""
		if col.Comment != nil {
			comment = *col.Comment
		}
		buf.WriteString(fmt.Sprintf("\t\t\tcatwalk.NewColumn(\"%s\", catwalk.%s, %d, %t, \"%s\")", col.Name, col.DataType.Go(), col.Count, col.Null, comment))
		if col.Default != nil && *col.Default != "" {
			buf.WriteString(fmt.Sprintf(".SetDefault(\"%s\")", *col.Default))
		}
		if col.Extra != nil && *col.Extra != "" {
			buf.WriteString(fmt.Sprintf(".SetExtra(\"%s\")", *col.Extra))
		}
		buf.WriteString(".Done(),\n")
	}
	buf.WriteString("		)")
	if len(p.Indexes) > 0 {
		buf.WriteString(".SetIndexes(\n")
		for _, idx := range p.Indexes {
			// catwalk.NewIndex("items_price_idx", catwalk.IndexTypeNotUnique, "price"),
			buf.WriteString(fmt.Sprintf("\t\t\tcatwalk.NewIndex(\"%s\", catwalk.%s, \"%s\"),\n", idx.Name, IndexType(idx.Type).Go(), strings.Join(idx.Columns, `", "`)))
		}
		buf.WriteString("		)")
	}
	if len(p.Relations) > 0 {
		buf.WriteString(".SetRelations(\n")
		for _, relation := range p.Relations {
			// catwalk.NewRelation("", "id", "account_devices", "account_id", false, false),
			buf.WriteString(fmt.Sprintf("\t\t\tcatwalk.NewRelation(\"%s\", \"%s\", \"%s\", \"%s\", %t, true),\n", relation.Name, relation.Column, relation.RefTable, relation.RefColumn, relation.HasOne))
		}
		buf.WriteString("		)")
	}
	buf.WriteString(".SetEnums(\n\t).Done()\n}\n")
	// ファイル出力
	filename := fmt.Sprintf("%s/%s.go", path, p.Name)
	if _, err := os.Stat(filename); err != nil { // ファイルの上書きはしない
		fp, err := os.Create(filename)
		if err != nil {
			fmt.Println(err)
			return false
		}
		defer fp.Close()
		if _, err := fp.Write(buf.Bytes()); err != nil {
			fmt.Printf("WriteString Error: %s\n", err.Error())
			return false
		}
	}
	return true
}

func (p *Column) GetGoTag(table *Table) string {
	ok := false
	for _, index := range table.Indexes {
		if index.Name == "PRIMARY KEY" {
			if !slices.Contains(index.Columns, p.Name) {
				ok = true
			}
		}
	}
	if ok {
		return ` gorm:"primarykey"`
	}
	switch p.DataType {
	case DataTypeArrayInt8:
		return ` gorm:"type:integer[]"`
	case DataTypeArrayString, DataTypeArrayMacaddr:
		return ` gorm:"type:text[]"`
	}
	return ""
}

func (p *Column) GetGoComment() string {
	if p.Comment != nil && len(*p.Comment) > 0 {
		return fmt.Sprintf(" // %s", *p.Comment)
	}
	return ""
}

func (p *Relation) GetReference(t *Table) string {
	con := pluralize.NewClient()
	if !p.HasOne && strcase.ToCamel(p.Column) == "Id" && p.RefColumn == fmt.Sprintf("%s_id", con.Singular(t.Name)) {
		return fmt.Sprintf("foreignKey:%s", strcase.ToCamel(p.Column))
	}
	return fmt.Sprintf("foreignKey:%s;references:%s", strcase.ToCamel(p.Column), strcase.ToCamel(p.RefColumn))
}

func (p *Relation) GetRelation() string {
	con := pluralize.NewClient()
	model := strcase.ToCamel(con.Singular(p.RefTable))
	if !p.HasOne {
		return fmt.Sprintf("%s []%s", strcase.ToCamel(p.RefTable), model)
	}
	return fmt.Sprintf("%s *%s", model, model)
}

func (p *Enum) GoCreate(table *Table) string {
	buf := bytes.NewBuffer([]byte{})
	name := p.GoStructName(table)
	dataType := p.Type.String()
	for _, col := range table.Columns {
		if p.Column == col.Name {
			dataType = col.DataType.String()
		}
	}
	buf.WriteString(fmt.Sprintf("type %s %s\n", name, dataType))
	buf.WriteString("const (\n")
	switch p.Type {
	case EnumTypeString:
		for _, value := range p.Values {
			buf.WriteString(fmt.Sprintf("\t%s%s = %s(\"%s\")\n", name, value.Key, name, value.Value.(string)))
		}
	case EnumTypeUint:
		for _, value := range p.Values {
			buf.WriteString(fmt.Sprintf("\t%s%s = %s(%d)\n", name, value.Key, name, value.Value.(int)))
		}
	case EnumTypeBitfield:
		ok := true
		for _, value := range p.Values {
			if ok {
				buf.WriteString(fmt.Sprintf("\t%s%s %s = 1 << iota\n", name, value.Key, name))
				ok = false
			} else {
				buf.WriteString(fmt.Sprintf("\t%s%s\n", name, value.Key))
			}
		}
	}
	buf.WriteString(")\n\n")

	switch p.Type {
	case EnumTypeString:
		buf.WriteString(fmt.Sprintf("func (p %s) String() string {\n", name))
		buf.WriteString("\treturn string(p)\n")
		buf.WriteString("}\n\n")
	case EnumTypeUint:
		buf.WriteString(fmt.Sprintf("func (p %s) String() string {\n", name))
		buf.WriteString("\tswitch p {\n")
		for _, value := range p.Values {
			buf.WriteString(fmt.Sprintf("\tcase %s%s:\n", name, value.Key))
			buf.WriteString(fmt.Sprintf("\t\treturn \"%s\"\n", value.Key))
		}
		buf.WriteString("\t}\n")
		buf.WriteString("\treturn \"\"\n")
		buf.WriteString("}\n\n")

		con := pluralize.NewClient()
		names := con.Plural(name)
		buf.WriteString(fmt.Sprintf("func %s(key string) %s {\n", names, name))
		buf.WriteString("\tswitch key {\n")
		for _, value := range p.Values {
			buf.WriteString(fmt.Sprintf("\tcase \"%s\":\n", value.Key))
			buf.WriteString(fmt.Sprintf("\t\treturn %s%s\n", name, value.Key))
		}
		buf.WriteString("\t}\n")
		buf.WriteString("\treturn 0\n")
		buf.WriteString("}\n\n")

		buf.WriteString(fmt.Sprintf("func (p %s) MarshalJSON() ([]byte, error) {\n", name))
		buf.WriteString("\treturn json.Marshal(p.String())")
		buf.WriteString("}\n\n")

		buf.WriteString(fmt.Sprintf("func (p *%s) UnmarshalJSON(data []byte) error {\n", name))
		buf.WriteString("\tvar value string\n")
		buf.WriteString("\tif err := json.Unmarshal(data, &value); err != nil {\n")
		buf.WriteString("\t\treturn err")
		buf.WriteString("\t}\n")
		buf.WriteString(fmt.Sprintf("\t*p = %s(value)\n", names))
		buf.WriteString("\treturn nil\n")
		buf.WriteString("}\n\n")
	case EnumTypeBitfield:
		buf.WriteString(fmt.Sprintf("func (p %s) Check(flag %s) bool {\n", name, name))
		buf.WriteString("\treturn (p & flag) == flag\n")
		buf.WriteString("}\n\n")
	}
	return buf.String()
}

func (p *Enum) GoStructName(table *Table) string {
	con := pluralize.NewClient()
	return fmt.Sprintf("%s%s", strcase.ToCamel(con.Singular(table.Name)), strcase.ToCamel(p.Column))
}
