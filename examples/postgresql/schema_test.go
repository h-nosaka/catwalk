package schema

import (
	"fmt"
	"os"
	"os/exec"
	"testing"

	"github.com/h-nosaka/catwalk/base"
	db "github.com/h-nosaka/catwalk/postgresql"
)

func SetEnv() {
	os.Setenv("RDB_TYPE", os.Getenv("PRDB_TYPE"))
	os.Setenv("RDB_USER", os.Getenv("PRDB_USER"))
	os.Setenv("RDB_PASSWORD", os.Getenv("PRDB_PASSWORD"))
	os.Setenv("RDB_HOST", os.Getenv("PRDB_HOST"))
	os.Setenv("RDB_DATABASE", os.Getenv("PRDB_DATABASE"))
}

func TestDump(t *testing.T) {
	SetEnv()
	base.Init()
	Schema().Sql("./dump/dump.sql")
	db.NewSchemaFromDB().Sql("./dump/db.sql")
}

func TestModel(t *testing.T) {
	SetEnv()
	base.Init()
	schema := Schema()
	for _, table := range schema.Tables {
		table.CreateGoModel("./models")
	}
	if err := exec.Command("go", "fmt", "./models/...").Run(); err != nil {
		t.Error(err)
	}
}

func TestCreateDatabase(t *testing.T) {
	SetEnv()
	base.Init()
	Schema().CreateDatabase()
}

func TestDiff(t *testing.T) {
	SetEnv()
	base.Init()
	src := db.NewSchemaFromDB()
	diff := Schema().Diff(src)
	t.Log(diff)
}

func TestRun(t *testing.T) {
	SetEnv()
	base.Init()
	Schema().Run()
}

func TestSchema(t *testing.T) {
	SetEnv()
	base.Init()
	src := db.NewSchemaFromDB()
	src.CreateSchema("./dump")
}

func TestTruncateTable(t *testing.T) {
	SetEnv()
	base.Init()
	src := db.NewSchemaFromDB()
	for _, table := range src.Tables {
		fmt.Printf("テーブルを削除します: %s\n", table.Name)
		if err := base.DB.Exec(table.Drop()).Error; err != nil {
			fmt.Printf("削除エラー: %s\n", err)
		}
	}
	fmt.Print("DBのマイグレーションを実施します\n")
	Schema().Run()
}
