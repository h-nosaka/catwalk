//go:build mage
// +build mage

package main

import (
	"fmt"
	"github.com/h-nosaka/catwalk/base"
	examples "github.com/h-nosaka/catwalk/examples"
	mdb "github.com/h-nosaka/catwalk/mysql"
)

func CreateDatabase() {
	base.Init()
	name := base.GetEnv("RDB_DATABASE", "app")
	fmt.Printf("データベースを作成します: %s\n", name)
	schema := &mdb.ISchema{}
	schema.CreateDatabase(name)
}

func CreateSQL() {
	base.Init()
	path := "./examples/dump/dump.sql"
	fmt.Printf("SQLファイルを次のパスに作成します: %s\n", path)
	examples.Schema().Sql(path)
}

func Diff() {
	base.Init()
	fmt.Print("現在のDBの状態と比較したDiffを出力します\n")
	diff := examples.Schema().Diff(mdb.NewSchemaFromDB())
	fmt.Println(diff)
}

func Migrate() {
	base.Init()
	fmt.Print("DBのマイグレーションを実施します\n")
	examples.Schema().Run()
}

func CreateModel() {
	base.Init()
	path := "./examples/models"
	fmt.Printf("モデルファイルを次のディレクトリに作成します: %s\n", path)
	examples.Schema().Model(path)
}

func CreateSchema() {
	base.Init()
	path := "./examples/dump"
	fmt.Printf("現在のDBの状態のスキーマファイルを次のディレクトリに作成します: %s\n", path)
	mdb.NewSchemaFromDB().CreateSchema("./examples/dump")
}
