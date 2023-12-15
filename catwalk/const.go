package catwalk

import (
	"encoding/json"
)

type ISchema interface {
	CreateDatabase(dbname ...string) // データベースの作成
	CreateSchema(path ...string)     // 現行のDBからデータを取得してschemaファイルを作成
	CreateModel(path ...string)      // schemaの各テーブルからgolang用modelを作成
	CreateFixture(path ...string)    // schemaの各テーブルからgolang用fixtureを作成
	CreateSql(filename string)       // SQLファイルを作成
	CreateYaml(filename string)      // YAMLファイルを作成
	Type() string                    // データベースのタイプ表示
	Diff() string                    // schemaの定義と現行のDBの差分をSQLで表示
	Run()                            // 現行のDBに対してschemaと等価になるようにMigrationを実行
}

type SchemaMode int

const (
	SchemaModeMySQL    SchemaMode = iota // MySQL
	SchemaModePostgres                   // Postgres
	SchemaModeES                         // Elasticsearch
)

type Schema struct {
	Name   string     // DBName
	Mode   SchemaMode // RDB, NoSQL
	Tables []Table    // テーブル定義
}

type Table struct {
	Schema     string
	Name       string
	JsonCase   JsonCase `yaml:"json_case,omitempty"`
	Comment    *string  `yaml:"comment,omitempty"`
	Rename     *string  `yaml:"rename,omitempty"`
	Columns    []Column
	Indexes    []Index
	Relations  []Relation
	Enums      []Enum
	Partitions *Partition
}

type JsonCase string

const (
	JsonCaseSnake  = JsonCase("snake")
	JsonCaseCamel  = JsonCase("camel")
	JsonCasePascal = JsonCase("pascal")
)

type Column struct {
	Name     string
	DataType DataType
	Count    int     `yaml:"count,omitempty"`
	Extra    *string `yaml:"extra,omitempty"`
	Default  *string `yaml:"default,omitempty"`
	Null     bool    `yaml:"nullable,omitempty"`
	Comment  *string `yaml:"comment,omitempty"`
	Rename   *string `yaml:"rename,omitempty"`
}

type Index struct {
	Name    string
	Type    string
	Columns []string
}

type IndexType string

const (
	IndexTypePrimary   IndexType = "PRIMARY KEY"
	IndexTypeUnique    IndexType = "UNIQUE"
	IndexTypeNotUnique IndexType = ""
)

// type Sequence struct {
// 	Sequencename string
// 	StartValue   int
// 	MinValue     int
// 	MaxValue     int
// 	IncrementBy  int
// }

type Relation struct {
	Name         string
	Column       string
	RefTable     string
	RefColumn    string
	HasOne       bool
	HasAny       bool
	IsForeignKey bool
}

type EnumType uint

const (
	EnumTypeString   = EnumType(0)
	EnumTypeUint     = EnumType(1)
	EnumTypeBitfield = EnumType(2)
	EnumTypeUnkown   = EnumType(255)
)

func (p EnumType) String() string {
	switch p {
	case EnumTypeString:
		return "string"
	case EnumTypeUint:
		return "uint"
	case EnumTypeBitfield:
		return "uint64"
	}
	return "unknown"
}

func EnumTypes(key string) EnumType {
	switch key {
	case "string":
		return EnumTypeString
	case "uint":
		return EnumTypeUint
	case "uint64":
		return EnumTypeBitfield
	}
	return EnumTypeUnkown
}

func (p EnumType) MarshalJSON() ([]byte, error) {
	return json.Marshal(p.String())
}

func (p *EnumType) UnmarshalJSON(data []byte) error {
	var value string
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*p = EnumTypes(value)
	return nil
}

type Enum struct {
	Column string
	Type   EnumType
	Values []EnumValue
}

type EnumValue struct {
	Key   string
	Value interface{}
}

type Partition struct {
	Type   string
	Column string
	Keys   []PartitionKey
}

type PartitionKey struct {
	Key   string
	Value string
}
