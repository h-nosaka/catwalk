package mysql

import (
	"bytes"
	"fmt"
	"os"
	"strings"

	"github.com/freebitdx/fbfiber/context"

	"github.com/gertd/go-pluralize"
	"github.com/iancoleman/strcase"
	"gorm.io/gorm"
)

type IMethod struct {
	Method  string
	Imports *[]string
}

type ITable struct {
	Schema      string        `gorm:"column:TABLE_SCHEMA"`
	Name        string        `gorm:"column:TABLE_NAME"`
	IsDB        *bool         `gorm:"->:false" yaml:"is_db,omitempty"`
	Comment     *string       `gorm:"column:TABLE_COMMENT" yaml:"comment,omitempty"`
	Rename      *string       `gorm:"->:false" yaml:"rename,omitempty"`
	Columns     []IColumn     `gorm:"->:false"`
	Indexes     []IIndex      `gorm:"->:false"`
	Foreignkeys []IForeignkey `gorm:"->:false"`
	Enums       []IEnum       `gorm:"->:false"`
	Methods     []IMethod     `gorm:"->:false"`
}

func (p *ITable) GetColumn(db *gorm.DB) {
	cols := []IColumn{}
	db.Raw(fmt.Sprintf(GetColumns, p.Schema, p.Name)).Scan(&cols)
	p.Columns = []IColumn{}
	for _, col := range cols {
		if col.Defaults != nil && *col.Defaults == "" {
			col.Defaults = nil
		}
		if col.Extra != nil && *col.Extra == "" {
			col.Extra = nil
		}
		col.Null = col.IsNull == "YES"
		if col.Comment != nil && *col.Comment == "" {
			col.Comment = nil
		}
		p.Columns = append(p.Columns, col)
	}
}

func (p *ITable) GetIndexes(db *gorm.DB) {
	p.Indexes = []IIndex{}
	var indexes []IIndex
	db.Raw(fmt.Sprintf(GetIndexes, p.Schema, p.Name)).Scan(&indexes)
	for _, index := range indexes {
		index.GetColumn(db, p)
		index.Unique = index.NonUnique == 0
		p.Indexes = append(p.Indexes, index)
	}
}

func (p *ITable) GetForeignkeys(db *gorm.DB) {
	db.Raw(fmt.Sprintf(GetForeignkeys, p.Schema, p.Name)).Scan(&p.Foreignkeys)
}

func (p *ITable) Create() string {
	if p.IsDB != nil && !*p.IsDB {
		return ""
	}
	buf := bytes.NewBuffer([]byte{})
	buf.WriteString(fmt.Sprintf("CREATE TABLE %s (\n", p.Name))
	for _, col := range p.Columns {
		buf.WriteString(col.Append())
		buf.WriteString(",\n")
	}
	for _, index := range p.Indexes {
		if index.Name == "PRIMARY" {
			buf.WriteString(index.Append())
			buf.WriteString("\n")
		}
	}
	if p.Comment != nil {
		buf.WriteString(fmt.Sprintf(") ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='%s';\n\n", *p.Comment))
	} else {
		buf.WriteString(") ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;\n\n")
	}
	for _, item := range p.Indexes {
		if item.Name != "PRIMARY" {
			buf.WriteString(item.Create(p))
		}
	}
	for _, item := range p.Foreignkeys {
		buf.WriteString(item.Create(p))
	}
	return buf.String()
}

func (p *ITable) Drop() string {
	return fmt.Sprintf(
		"DROP TABLE IF EXISTS %s RESTRICT;\n",
		p.Name,
	)
}

func (p *ITable) RenameTable() string {
	if p.Rename != nil {
		return fmt.Sprintf(
			"RENAME TABLE %s TO %s;\n",
			p.Name,
			*p.Rename,
		)
	}
	return ""
}

func (p ITable) Diff(src *[]ITable) string {
	buf := bytes.NewBuffer([]byte{})
	dest := ITable{}
	for _, item := range *src {
		if item.Name == p.Name {
			dest = item
		}
	}
	if dest.Name == "" {
		buf.WriteString(p.Create())
	} else if p.Create() != dest.Create() {
		for _, col := range p.Columns {
			buf.WriteString(col.Diff(&p, &dest))
		}
		for _, col := range dest.Columns {
			buf.WriteString(col.IsDrop(&p))
		}
		// インデックスの差分
		for _, index := range p.Indexes {
			buf.WriteString(index.Diff(&p, &dest))
		}
		// 制約の差分
		for _, key := range p.Foreignkeys {
			buf.WriteString(key.Diff(&p, &dest))
		}
		// カラムのリネーム
		for _, col := range p.Columns {
			if col.Rename != nil && col.Name != *col.Rename {
				buf.WriteString(col.RenameColumn(&p))
			}
		}
	}
	// テーブルのリネーム
	if p.Rename != nil && p.Name != *p.Rename {
		buf.WriteString(p.RenameTable())
	}
	return buf.String()
}

func (p *ITable) GoModel(path string) {
	buf := bytes.NewBuffer([]byte{})
	buf.WriteString("package model\n\n")
	// インポート定義
	imports := []string{`"time"`}
	if p.IsDB == nil || (p.IsDB != nil && *p.IsDB) {
		imports = append(imports, `"gorm.io/gorm"`)
	}
	for _, item := range p.Columns {
		switch item.DataType {
		case "json":
			if context.ArrayInclude(imports, `"encoding/json"`) {
				imports = append(imports, `"encoding/json"`)
			}
		}
	}
	for _, item := range p.Methods {
		if item.Imports != nil {
			imports = append(imports, *item.Imports...)
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
		data := item.GetGoType()
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
	for _, item := range p.Foreignkeys {
		buf.WriteString(fmt.Sprintf(
			"\t%s `gorm:\"%s\"`\n",
			item.GetRelation(),
			item.GetReference(),
			// fmt.Sprintf("foreignKey:%s;references:%s", strcase.ToCamel(item.Column), strcase.ToCamel(item.RefColumn)),
		))
	}
	buf.WriteString("}\n\n")
	// テーブル名が複数形でない場合の処理
	if con.IsSingular(p.Name) {
		buf.WriteString(fmt.Sprintf("func (c *%s) TableName() string {\n", strcase.ToCamel(table)))
		buf.WriteString(fmt.Sprintf("\treturn \"%s\"\n", p.Name))
		buf.WriteString("}\n\n")
	}
	// 共通メソッドの定義
	if p.IsDB == nil || (p.IsDB != nil && *p.IsDB) {
		buf.WriteString(fmt.Sprintf("func (p *%s) Find(db *gorm.DB, preloads ...string) error {\n", strcase.ToCamel(table)))
		buf.WriteString("\ttx := db\n")
		buf.WriteString("\tfor _, preload := range preloads {\n")
		buf.WriteString("\t\ttx = tx.Preload(preload)\n")
		buf.WriteString("\t}\n")
		buf.WriteString("\tif err := tx.First(p).Error; err != nil {\n")
		buf.WriteString("\t\treturn err\n")
		buf.WriteString("\t}\n")
		buf.WriteString("\treturn nil\n")
		buf.WriteString("}\n\n")
	}
	// json型のmap変換メソッド
	for _, item := range p.Columns {
		if item.DataType == "json" {
			buf.WriteString(fmt.Sprintf("func (c *%s) %sMap() map[string]interface{} {\n", strcase.ToCamel(table), strcase.ToCamel(item.Name)))
			buf.WriteString("\tvar data map[string]interface{}\n")
			buf.WriteString(fmt.Sprintf("\t\tif err := json.Unmarshal([]byte(c.%s), &data); err != nil {\n", strcase.ToCamel(item.Name)))
			buf.WriteString("\t\t\treturn nil\n")
			buf.WriteString("\t\t}\n")
			buf.WriteString("\treturn data\n")
			buf.WriteString("}\n\n")
		}
	}
	// mapping.yamlで定義されたactionを追加
	for _, item := range p.Methods {
		buf.WriteString(fmt.Sprintf("%s\n\n", item.Method))
	}

	// ファイル出力
	filename := fmt.Sprintf("%s/%s.go", path, table)
	fp, err := os.Create(filename)
	if err != nil {
		panic(err)
	}
	defer fp.Close()
	if _, err := fp.Write(buf.Bytes()); err != nil {
		fmt.Printf("WriteString Error: %s\n", err.Error())
	}
}

func (p *ITable) DartModel(path string) string {
	buf := bytes.NewBuffer([]byte{})
	init := bytes.NewBuffer([]byte{})

	// コメント処理
	if p.Comment != nil {
		buf.WriteString(fmt.Sprintf("// %s\n", *p.Comment))
	}
	con := pluralize.NewClient()
	table := con.Singular(p.Name)
	// Enumの定義
	for _, enum := range p.Enums {
		buf.WriteString(enum.DartCreate(p))
	}
	// 構造体定義
	buf.WriteString("@JsonSerializable()\n")
	buf.WriteString(fmt.Sprintf("class %s {\n", strcase.ToCamel(table)))
	// テーブルカラムの定義
	buf.WriteString("\t// column\n")
	for _, item := range p.Columns {
		data := item.GetDartType()
		// for _, enum := range mapping.Enums {
		// 	if p.Tablename == enum.Tablename && item.ColumnName == enum.Columnname {
		// 		data += fmt.Sprintf("|%s", enum.GoStructName())
		// 	}
		// }
		buf.WriteString(fmt.Sprintf(
			"\tlate %s %s;%s\n",
			data,
			strcase.ToLowerCamel(item.Name),
			item.GetGoComment(), // コメント処理
		))
		required := ""
		if !item.Null {
			required = "required "
		}
		init.WriteString(fmt.Sprintf(
			"\t\t%sthis.%s,\n",
			required,
			strcase.ToLowerCamel(item.Name),
		))
	}
	buf.WriteString("\n")
	// リレーションの定義
	buf.WriteString("\t// relation\n")
	for _, item := range p.Foreignkeys {
		buf.WriteString(fmt.Sprintf(
			"\t%s\n",
			item.GetDartRelation(),
		))
		property := strcase.ToLowerCamel(con.Singular(item.RefTable))
		if item.HasOne {
			property = strcase.ToLowerCamel(item.RefTable)
		}
		init.WriteString(fmt.Sprintf(
			"\t\tthis.%s,\n",
			property,
		))
	}
	// 初期化の定義
	buf.WriteString("\t// init\n")
	buf.WriteString(fmt.Sprintf("\t%s({\n", strcase.ToCamel(table)))
	buf.WriteString(init.String())
	buf.WriteString("\t});\n\n")
	// json化の定義
	buf.WriteString(fmt.Sprintf("\tfactory %s.fromJson(Map<String, dynamic> json) => _$%sFromJson(json);\n", strcase.ToCamel(table), strcase.ToCamel(table)))
	buf.WriteString(fmt.Sprintf("\tMap<String, dynamic> toJson() => _$%sToJson(this);\n", strcase.ToCamel(table)))

	buf.WriteString("}\n\n")
	return buf.String()
}
