package mysql

import (
	"bytes"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/gertd/go-pluralize"
	"github.com/h-nosaka/catwalk/base"
	"github.com/iancoleman/strcase"
	"golang.org/x/exp/slices"
)

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
	Relations   []IRelation   `gorm:"->:false"`
	Partitions  *IPartition   `gorm:"->:false"`
}

type IPartition struct {
	Type   string
	Column string
	Keys   []IPartitionKey
}

type IPartitionKey struct {
	Key   string
	Value string
}

func (p *ITable) GetColumn() {
	cols := []IColumn{}
	base.DB.Raw(fmt.Sprintf(GetColumns, p.Schema, p.Name)).Scan(&cols)
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

func (p *ITable) GetIndexes() {
	p.Indexes = []IIndex{}
	var indexes []IIndex
	base.DB.Raw(fmt.Sprintf(GetIndexes, p.Schema, p.Name)).Scan(&indexes)
	for _, index := range indexes {
		index.GetColumn(p)
		index.Unique = index.NonUnique == 0
		p.Indexes = append(p.Indexes, index)
	}
}

func (p *ITable) GetForeignkeys() {
	base.DB.Raw(fmt.Sprintf(GetForeignkeys, p.Schema, p.Name)).Scan(&p.Foreignkeys)
}

func (p *ITable) GetPartitions() {
	rs := []map[string]interface{}{}
	base.DB.Raw(fmt.Sprintf(GetPartitions, p.Schema, p.Name)).Scan(&rs)
	if len(rs) > 0 && rs[0]["PARTITION_METHOD"] != nil {
		p.Partitions = &IPartition{
			Type:   rs[0]["PARTITION_METHOD"].(string),
			Column: rs[0]["PARTITION_EXPRESSION"].(string),
			Keys:   []IPartitionKey{},
		}
		for _, item := range rs {
			p.Partitions.Keys = append(p.Partitions.Keys, IPartitionKey{Key: item["PARTITION_NAME"].(string), Value: item["PARTITION_DESCRIPTION"].(string)})
		}
	}
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
	if p.Partitions != nil {
		buf.WriteString(p.Partitions.Create(p.Name))
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
		// パーティション
		if p.Partitions != nil {
			buf.WriteString(p.Partitions.Diff(p.Name, dest.Partitions))
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

func (p *IPartition) Create(table string) string {
	buf := bytes.NewBuffer([]byte{})
	buf.WriteString(fmt.Sprintf("ALTER TABLE %s PARTITION BY %s %s (\n", table, p.Type, p.Column))
	parts := []string{}
	for _, item := range p.Keys {
		switch p.Type {
		case "RANGE":
			parts = append(append, fmt.Sprintf("\tPARTITION %s VALUES LESS THAN (%s)", item.Key, item.Value))
		case "LIST":
			parts = append(append, fmt.Sprintf("\tPARTITION %s VALUES IN (%s),\n", item.Key, item.Value))
		}
	}
	buf.WriteString(strings.Join(parts, ",\n"))
	buf.WriteString("\n);\n\n")
	return buf.String()
}

func (p *IPartition) Diff(table string, src *IPartition) string {
	buf := bytes.NewBuffer([]byte{})
	// 新規パーティションテーブルの場合
	if p != nil && src == nil {
		buf.WriteString(p.Create(table))
		return buf.String()
	}
	// パーティション全体削除の場合
	if p == nil && src != nil {
		buf.WriteString(fmt.Sprintf("ALTER TABLE %s REMOVE PARTITIONING;\n\n", table))
		return buf.String()
	}
	// パーティション追加
	for _, item := range p.Keys {
		ok := false
		for _, dest := range src.Keys {
			if item.Key == dest.Key {
				ok = true
				break
			}
		}
		if !ok {
			switch p.Type {
			case "RANGE":
				buf.WriteString(fmt.Sprintf("ALTER TABLE %s ADD PARTITION (PARTITION %s VALUES LESS THAN (%s));\n", table, item.Key, item.Value))
			case "LIST":
				buf.WriteString(fmt.Sprintf("ALTER TABLE %s ADD PARTITION (PARTITION %s VALUES IN (%s));\n", table, item.Key, item.Value))
			}
		}
	}
	// パーティション削除
	for _, dest := range src.Keys {
		ok := false
		for _, item := range p.Keys {
			if dest.Key == item.Key {
				ok = true
				break
			}
		}
		if !ok {
			buf.WriteString(fmt.Sprintf("ALTER TABLE %s TRUNCATE PARTITION %s;\n", table, dest.Key)) // パーティションの中身をクリアする
			buf.WriteString(fmt.Sprintf("ALTER TABLE %s DROP PARTITION %s;\n", table, dest.Key))
		}
	}
	buf.WriteString("\n")
	return buf.String()
}

func (p *ITable) CreateGoModel(path string) {
	buf := bytes.NewBuffer([]byte{})
	buf.WriteString("package models\n\n")
	// インポート定義
	imports := []string{`"time"`}
	if p.IsDB == nil || (p.IsDB != nil && *p.IsDB) {
		imports = append(imports, `"gorm.io/gorm"`)
	}
	for _, item := range p.Columns {
		switch item.DataType {
		case "json":
			if !slices.Contains(imports, `"encoding/json"`) {
				imports = append(imports, `"encoding/json"`)
			}
		}
	}
	for _, item := range p.Methods {
		if item.Imports != nil {
			imports = append(imports, *item.Imports...)
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
	for _, item := range p.Relations {
		buf.WriteString(fmt.Sprintf(
			"\t%s `gorm:\"%s\"`\n",
			item.GetRelation(),
			item.GetReference(p),
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
	// Methodsで定義されたactionを追加
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

func (p *ITable) CreateSchemaFile() []byte {
	// con := pluralize.NewClient()
	buf := bytes.NewBuffer([]byte{})
	buf.WriteString(`package schema

import (
	"github.com/h-nosaka/catwalk/base"
	db "github.com/h-nosaka/catwalk/mysql"
)

`)
	buf.WriteString(fmt.Sprintf("func %s() db.ITable {\n", strcase.ToCamel(p.Name)))
	buf.WriteString("	return db.ITable{\n")
	buf.WriteString(fmt.Sprintf(`		Schema: "%s",
`, p.Schema))
	buf.WriteString(fmt.Sprintf(`		Name: "%s",
`, p.Name))
	if p.Comment != nil {
		buf.WriteString(fmt.Sprintf(`		Comment: base.String("%s"),
`, *p.Comment))
	}
	buf.WriteString("		Columns: []db.IColumn{\n")
	for _, col := range p.Columns {
		def := "nil"
		if col.Defaults != nil && *col.Defaults != "" {
			def = fmt.Sprintf(`base.String("%s")`, *col.Defaults)
		}
		null := fmt.Sprintf("base.Bool(%s)", strconv.FormatBool(col.Null))
		comment := "nil"
		if col.Comment != nil && *col.Comment != "" {
			comment = fmt.Sprintf(`base.String("%s")`, *col.Comment)
		}
		extra := "nil"
		if col.Extra != nil && *col.Extra != "" {
			extra = fmt.Sprintf(`base.String("%s")`, *col.Extra)
		}
		buf.WriteString(fmt.Sprintf(`			db.NewColumn("%s", "%s", %s, %s, %s, %s, nil),
`, col.Name, col.DataType, extra, def, null, comment))
	}
	buf.WriteString("		},\n")
	buf.WriteString("		Indexes: []db.IIndex{\n")
	for _, idx := range p.Indexes {
		buf.WriteString(fmt.Sprintf(`			db.NewIndex("%s", base.Bool(%t), "%s"),
`, idx.Name, idx.Unique, strings.Join(idx.Columns, `", "`)))
	}
	buf.WriteString("		},\n")
	buf.WriteString("		Foreignkeys: []db.IForeignkey{\n")
	for _, fk := range p.Foreignkeys {
		buf.WriteString(fmt.Sprintf(`			db.NewFK("%s", "%s", "%s", "%s", true, false),
`, fk.Name, fk.Column, fk.RefTable, fk.RefColumn))
	}
	buf.WriteString("		},\n")
	buf.WriteString(`		Enums: []db.IEnum{},
		Methods: []db.IMethod{},
		Relations: []db.IRelation{},
`)
	buf.WriteString("	}\n}\n")
	return buf.Bytes()
}
