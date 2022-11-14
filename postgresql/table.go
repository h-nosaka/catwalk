package postgresql

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

type PGTable struct {
	Schemaname string
	Tablename  string
	Comment    *string    `yaml:"comment,omitempty"`
	Rename     *string    `gorm:"->:false" yaml:"rename,omitempty"`
	Columns    []PGColumn `gorm:"->:false"`
}

func (p *PGTable) GetColumn(db *gorm.DB) {
	db.Raw(`
		SELECT columns.*, pg_description.description as comment
			FROM information_schema.columns
			JOIN pg_catalog.pg_class ON pg_class.relname = ?
			JOIN pg_catalog.pg_attribute ON pg_class.oid = pg_attribute.attrelid AND columns.column_name = pg_attribute.attname
			LEFT JOIN pg_catalog.pg_description ON pg_attribute.attrelid = pg_description.objoid AND pg_attribute.attnum = pg_description.objsubid
			WHERE table_name = ? order by ordinal_position;
	`, p.Tablename, p.Tablename).Scan(&p.Columns)
}

func (p *PGTable) Create() string {
	buf := bytes.NewBuffer([]byte{})
	buf.WriteString(fmt.Sprintf("CREATE TABLE %s.%s (\n", p.Schemaname, p.Tablename))
	for i, col := range p.Columns {
		buf.WriteString(col.Append())
		if i+1 < len(p.Columns) {
			buf.WriteString(",")
		}
		buf.WriteString("\n")
	}
	buf.WriteString(");\n\n")
	if p.Comment != nil {
		buf.WriteString(p.SetComment())
	}
	for _, col := range p.Columns {
		if col.Comment != nil {
			buf.WriteString(col.SetComment(p))
		}
	}
	return buf.String()
}

func (p *PGTable) Drop() string {
	return fmt.Sprintf(
		"DROP TABLE IF EXISTS %s.%s RESTRICT;\n",
		p.Schemaname,
		p.Tablename,
	)
}

func (p *PGTable) SetComment() string {
	if p.Comment == nil {
		return ""
	}
	return fmt.Sprintf("COMMENT ON TABLE %s.%s IS '%s';\n\n", p.Schemaname, p.Tablename, *p.Comment)
}

func (p *PGTable) DropComment() string {
	return fmt.Sprintf("COMMENT ON TABLE %s.%s IS NULL;\n\n", p.Schemaname, p.Tablename)
}

func (p *PGTable) RenameTable() string {
	if p.Rename != nil {
		return fmt.Sprintf(
			"ALTER TABLE %s.%s RENAME TO %s;\n",
			p.Schemaname,
			p.Tablename,
			*p.Rename,
		)
	}
	return ""
}

func (p PGTable) Diff(src *[]PGTable) string {
	buf := bytes.NewBuffer([]byte{})
	dest := PGTable{}
	for _, item := range *src {
		if item.Tablename == p.Tablename {
			dest = item
		}
	}
	if dest.Tablename == "" {
		buf.WriteString(p.Create())
	} else if p.Create() != dest.Create() {
		for _, col := range p.Columns {
			buf.WriteString(col.Diff(&dest.Columns, &p))
		}
		for _, col := range dest.Columns {
			buf.WriteString(col.IsDrop(&p.Columns, &p))
		}
	}
	// テーブルコメント
	if p.SetComment() != dest.SetComment() {
		if p.Comment != nil {
			buf.WriteString(p.SetComment())
		} else {
			buf.WriteString(p.DropComment())
		}
	}
	// カラムのリネーム
	for _, col := range p.Columns {
		if col.Rename != nil && col.ColumnName != *col.Rename {
			buf.WriteString(col.RenameColumn(&p))
		}
	}
	// テーブルのリネーム
	if p.Rename != nil && p.Tablename != *p.Rename {
		buf.WriteString(p.RenameTable())
	}
	return buf.String()
}

func (p *PGTable) GoModel(path string, indexes *[]PGIndex, foreignkeys *[]PGForeignkey, mapping *Mappings) {
	buf := bytes.NewBuffer([]byte{})
	buf.WriteString("package model\n\n")
	// インポート定義
	imports := []string{`"time"`, `"gorm.io/gorm"`}
	for _, item := range p.Columns {
		switch item.UdtName {
		case "_varchar", "_macaddr", "_int4":
			if context.ArrayInclude(imports, `"github.com/lib/pq"`) {
				imports = append(imports, `"github.com/lib/pq"`)
			}
		// TODO: pgtypeはいったん不使用
		// case "cidr", "inet":
		// 	if context.ArrayInclude(imports, `"github.com/jackc/pgtype"`) {
		// 		imports = append(imports, `"github.com/jackc/pgtype"`)
		// 	}
		case "json":
			if context.ArrayInclude(imports, `"encoding/json"`) {
				imports = append(imports, `"encoding/json"`)
			}
		}
	}
	for _, item := range mapping.Actions {
		if p.Tablename == item.Tablename {
			if item.Imports != nil {
				imports = append(imports, *item.Imports...)
			}
		}
	}
	buf.WriteString(fmt.Sprintf("import (\n\t%s\n)\n\n", strings.Join(imports, "\n\t")))
	// コメント処理
	if p.Comment != nil {
		buf.WriteString(fmt.Sprintf("// %s\n", *p.Comment))
	}
	con := pluralize.NewClient()
	table := con.Singular(p.Tablename)
	// Enumの定義
	for _, enum := range mapping.Enums {
		if p.Tablename == enum.Tablename {
			buf.WriteString(enum.GoCreate())
		}
	}
	// 構造体定義
	buf.WriteString(fmt.Sprintf("type %s struct {\n", strcase.ToCamel(table)))
	// テーブルカラムの定義
	buf.WriteString("\t// column\n")
	for _, item := range p.Columns {
		data := item.GetGoType()
		for _, enum := range mapping.Enums {
			if p.Tablename == enum.Tablename && item.ColumnName == enum.Columnname {
				data = enum.GoStructName()
			}
		}
		buf.WriteString(fmt.Sprintf(
			"\t%s %s `json:\"%s\"%s`%s\n",
			strcase.ToCamel(item.ColumnName),
			data,
			item.ColumnName,
			item.GetGoTag(p, indexes), // gorm用タグの設定
			item.GetGoComment(),       // コメント処理
		))
	}
	buf.WriteString("\n")
	// リレーションの定義
	buf.WriteString("\t// relation\n")
	for _, item := range *foreignkeys {
		if item.TableName == p.Tablename {
			model := strcase.ToCamel(con.Singular(item.Reftable))
			buf.WriteString(fmt.Sprintf(
				"\t%s `gorm:\"%s\"`\n",
				item.GetRelation(false, model, &mapping.Mappings),
				fmt.Sprintf("foreignKey:%s;references:%s", strcase.ToCamel(item.Columnname), strcase.ToCamel(item.Refcolumn)),
			))
		}
		if item.Reftable == p.Tablename {
			model := strcase.ToCamel(con.Singular(item.TableName))
			buf.WriteString(fmt.Sprintf(
				"\t%s `gorm:\"%s\"`\n",
				item.GetRelation(true, model, &mapping.Mappings),
				fmt.Sprintf("foreignKey:%s;references:%s", strcase.ToCamel(item.Refcolumn), strcase.ToCamel(item.Columnname)),
			))
		}
	}
	buf.WriteString("}\n\n")
	// テーブル名が複数形でない場合の処理
	if con.IsSingular(p.Tablename) {
		buf.WriteString(fmt.Sprintf("func (c *%s) TableName() string {\n", strcase.ToCamel(table)))
		buf.WriteString(fmt.Sprintf("\treturn \"%s\"\n", p.Tablename))
		buf.WriteString("}\n\n")
	}
	// 共通メソッドの定義
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
	// json型のmap変換メソッド
	for _, item := range p.Columns {
		if item.UdtName == "json" {
			buf.WriteString(fmt.Sprintf("func (c *%s) %sMap() map[string]interface{} {\n", strcase.ToCamel(table), strcase.ToCamel(item.ColumnName)))
			buf.WriteString("\tvar data map[string]interface{}\n")
			buf.WriteString(fmt.Sprintf("\t\tif err := json.Unmarshal([]byte(c.%s), &data); err != nil {\n", strcase.ToCamel(item.ColumnName)))
			buf.WriteString("\t\t\treturn nil\n")
			buf.WriteString("\t\t}\n")
			buf.WriteString("\treturn data\n")
			buf.WriteString("}\n\n")
		}
	}
	// mapping.yamlで定義されたactionを追加
	for _, item := range mapping.Actions {
		if p.Tablename == item.Tablename {
			buf.WriteString(fmt.Sprintf("%s\n\n", item.Action))
		}
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
