package schema

import (
	"os/exec"
	"testing"

	"github.com/h-nosaka/catwalk/base"
	mdb "github.com/h-nosaka/catwalk/mysql"
)

func TestDump(t *testing.T) {
	base.Init()
	Schema().Sql("./dump/dump.sql")
	mdb.NewSchemaFromDB().Sql("./dump/db.sql")
}

func TestModel(t *testing.T) {
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
	base.Init()
	Schema().CreateDatabase()
}

func TestDiff(t *testing.T) {
	base.Init()
	src := mdb.NewSchemaFromDB()
	diff := Schema().Diff(src)
	t.Log(diff)
}

func TestRun(t *testing.T) {
	base.Init()
	Schema().Run()
}

func TestSchema(t *testing.T) {
	base.Init()
	src := mdb.NewSchemaFromDB()
	src.CreateSchema("./dump")
}
