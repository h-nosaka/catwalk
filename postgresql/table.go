package postgresql

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

type IMethod struct {
	Method  string
	Imports *[]string
}

type JsonCase string

const (
	JsonCaseSnake  = JsonCase("snake")
	JsonCaseCamel  = JsonCase("camel")
	JsonCasePascal = JsonCase("pascal")
)

type ITable struct {
	Schema      string        `gorm:"column:schemaname"`
	Name        string        `gorm:"column:tablename"`
	IsDB        *bool         `gorm:"->:false" yaml:"is_db,omitempty"`
	UseSchema   *bool         `gorm:"->:false"`
	JsonCase    JsonCase      `gorm:"->:false"`
	Comment     *string       `gorm:"column:comment" yaml:"comment,omitempty"`
	Rename      *string       `gorm:"->:false" yaml:"rename,omitempty"`
	Columns     []IColumn     `gorm:"->:false"`
	Indexes     []IIndex      `gorm:"->:false"`
	Foreignkeys []IForeignkey `gorm:"->:false"`
	Sequences   []ISequence   `gorm:"->:false"`
	Enums       []IEnum       `gorm:"->:false"`
	Methods     []IMethod     `gorm:"->:false"`
	Relations   []IRelation   `gorm:"->:false"`
}

func NewMethod(method string, imports ...string) IMethod {
	return IMethod{
		Method:  method,
		Imports: &imports,
	}
}

func NewRelation(column string, refTable string, refColumn string, one bool) IRelation {
	return IRelation{
		Column:    column,
		RefTable:  refTable,
		RefColumn: refColumn,
		HasOne:    one,
	}
}

func (p *ITable) GetColumn() {
	cols := []IColumn{}
	base.DB.Raw(GetColumns, p.Name, p.Name).Scan(&cols)
	p.Sequences = []ISequence{}
	p.Columns = []IColumn{}
	for _, col := range cols {
		if col.Defaults != nil && strings.Contains(*col.Defaults, "nextval") {
			sequence := strings.Split(*col.Defaults, "'")
			if len(sequence) < 2 {
				continue
			}
			p.Sequences = append(p.Sequences, p.GetSequences(sequence[1]))
		}
		if col.Defaults != nil && *col.Defaults == "" {
			col.Defaults = nil
		}
		col.Null = col.IsNull == "YES"
		if col.Comment != nil && *col.Comment == "" {
			col.Comment = nil
		}
		p.Columns = append(p.Columns, col)
	}
}

func (p *ITable) GetSequences(name string) ISequence {
	rs := ISequence{}
	base.DB.Raw(GetSequences, name).Scan(&rs)
	return rs
}

func (p *ITable) GetIndexes() {
	p.Indexes = []IIndex{}
	var indexes []IIndex
	base.DB.Raw(GetIndexes, p.Name).Scan(&indexes)
	for _, index := range indexes {
		index.GetColumn()
		p.Indexes = append(p.Indexes, index)
	}
}

func (p *ITable) GetForeignkeys() {
	base.DB.Raw(GetForeignkeys, p.Name).Scan(&p.Foreignkeys)
}

func (p *ITable) SchemaName() string {
	if p.UseSchema == nil || !*p.UseSchema {
		return ""
	}
	return fmt.Sprintf("%s.", p.Schema)
}

func (p *ITable) Create() string {
	buf := bytes.NewBuffer([]byte{})
	for _, seq := range p.Sequences {
		buf.WriteString(seq.Create(p))
	}
	buf.WriteString(fmt.Sprintf("CREATE TABLE %s%s (\n", p.SchemaName(), p.Name))
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
	for _, item := range p.Indexes {
		buf.WriteString(item.Create(p))
	}
	for _, item := range p.Foreignkeys {
		buf.WriteString(item.Create(p))
	}
	return buf.String()
}

func (p *ITable) Drop() string {
	return fmt.Sprintf(
		"DROP TABLE IF EXISTS %s%s RESTRICT;\n",
		p.SchemaName(),
		p.Name,
	)
}

func (p *ITable) SetComment() string {
	if p.Comment == nil {
		return ""
	}
	return fmt.Sprintf("COMMENT ON TABLE %s%s IS '%s';\n\n", p.SchemaName(), p.Name, *p.Comment)
}

func (p *ITable) DropComment() string {
	return fmt.Sprintf("COMMENT ON TABLE %s%s IS NULL;\n\n", p.SchemaName(), p.Name)
}

func (p *ITable) RenameTable() string {
	if p.Rename != nil {
		return fmt.Sprintf(
			"ALTER TABLE %s%s RENAME TO %s;\n",
			p.SchemaName(),
			p.Name,
			*p.Rename,
		)
	}
	return ""
}

func (p *ITable) Diff(src *[]ITable) string {
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
			buf.WriteString(col.Diff(&dest.Columns, p))
		}
		for _, col := range dest.Columns {
			buf.WriteString(col.IsDrop(&p.Columns, p))
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
		if col.Rename != nil && col.Name != *col.Rename {
			buf.WriteString(col.RenameColumn(p))
		}
	}
	// テーブルのリネーム
	if p.Rename != nil && p.Name != *p.Rename {
		buf.WriteString(p.RenameTable())
	}
	return buf.String()
}

func (p *ITable) CreateGoModel(path string) {
	buf := bytes.NewBuffer([]byte{})
	buf.WriteString("package models\n\n")
	// インポート定義
	imports := []string{}
	if p.IsDB == nil || (p.IsDB != nil && *p.IsDB) {
		imports = append(imports, `"gorm.io/gorm"`)
	}
	for _, item := range p.Columns {
		switch item.DataType {
		case "_varchar", "_macaddr", "_int4":
			if !slices.Contains(imports, `"github.com/lib/pq"`) {
				imports = append(imports, `"github.com/lib/pq"`)
			}
		case "json":
			if !slices.Contains(imports, `"encoding/json"`) {
				imports = append(imports, `"encoding/json"`)
			}
		case "timestamp":
			if !slices.Contains(imports, `"time"`) {
				imports = append(imports, `"time"`)
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
		to := item.Name
		switch p.JsonCase {
		case JsonCaseCamel:
			to = strcase.ToLowerCamel(item.Name)
		case JsonCasePascal:
			to = strcase.ToCamel(item.Name)
		case JsonCaseSnake:
			to = strcase.ToSnake(item.Name)
		}
		buf.WriteString(fmt.Sprintf(
			"\t%s %s `json:\"%s\"%s`%s\n",
			strcase.ToCamel(item.Name),
			data,
			to,
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
	if p.UseSchema != nil && *p.UseSchema && p.Schema != "public" { // スキーマ名がpublic以外の場合
		buf.WriteString(fmt.Sprintf("func (c *%s) TableName() string {\n", strcase.ToCamel(table)))
		buf.WriteString(fmt.Sprintf("\treturn `%s.\"%s\"`\n", p.Schema, p.Name))
		buf.WriteString("}\n\n")
	} else if con.IsSingular(p.Name) { // テーブル名が複数形でない場合
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
		buf.WriteString("\tif err := tx.Where(p).First(p).Error; err != nil {\n")
		buf.WriteString("\t\treturn err\n")
		buf.WriteString("\t}\n")
		buf.WriteString("\treturn nil\n")
		buf.WriteString("}\n\n")
	}
	// json型のmap変換メソッド
	for _, item := range p.Columns {
		if item.DataType == "json" {
			buf.WriteString(fmt.Sprintf("func (c *%s) %sMap() interface{} {\n", strcase.ToCamel(table), strcase.ToCamel(item.Name)))
			buf.WriteString("\tvar data interface{}\n")
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

func (p *ITable) CreateSchemaFile() []byte {
	// con := pluralize.NewClient()
	buf := bytes.NewBuffer([]byte{})
	buf.WriteString(`package schema

import (
	"github.com/h-nosaka/catwalk/base"
	db "github.com/h-nosaka/catwalk/postgresql"
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
			def = fmt.Sprintf("base.String(`%s`)", *col.Defaults)
		}
		null := fmt.Sprintf("base.Bool(%s)", strconv.FormatBool(col.Null))
		comment := "nil"
		if col.Comment != nil && *col.Comment != "" {
			comment = fmt.Sprintf(`base.String("%s")`, *col.Comment)
		}
		buf.WriteString(fmt.Sprintf(`			db.NewColumn("%s", "%s", %d, %d, %s, %s, %s, nil, nil),
`, col.Name, col.DataType, col.CharaMax, col.NumMax, def, null, comment))
	}
	buf.WriteString("		},\n")
	buf.WriteString("		Indexes: []db.IIndex{\n")
	for _, idx := range p.Indexes {
		if idx.ConstraintType == nil {
			buf.WriteString(fmt.Sprintf(`			db.NewIndex("%s", nil, "%s"),
`, idx.Name, strings.Join(idx.Columns, `", "`)))
		} else {
			buf.WriteString(fmt.Sprintf(`			db.NewIndex("%s", base.String("%s"), "%s"),
`, idx.Name, *idx.ConstraintType, strings.Join(idx.Columns, `", "`)))
		}
	}
	buf.WriteString("		},\n")
	buf.WriteString("		Foreignkeys: []db.IForeignkey{\n")
	for _, fk := range p.Foreignkeys {
		buf.WriteString(fmt.Sprintf(`			db.NewFK("%s", "%s", "%s", "%s", true, false),
`, fk.Name, fk.Column, fk.RefTable, fk.RefColumn))
	}
	buf.WriteString("		},\n")
	buf.WriteString("		Sequences: []db.ISequence{\n")
	for _, seq := range p.Sequences {
		buf.WriteString(fmt.Sprintf(`			db.NewSeq("%s", %d, %d, %d, %d),
`, seq.Sequencename, seq.StartValue, seq.MinValue, seq.MaxValue, seq.IncrementBy))
	}
	buf.WriteString("		},\n")
	buf.WriteString(`		Enums: []db.IEnum{},
		Methods: []db.IMethod{},
		Relations: []db.IRelation{},
`)
	buf.WriteString("	}\n}\n")
	return buf.Bytes()
}
