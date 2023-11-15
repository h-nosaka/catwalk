package schema

import (
	"os"
	"os/exec"
	"testing"

	"github.com/h-nosaka/catwalk/base"
	"github.com/h-nosaka/catwalk/examples/postgresql/models"
	mdb "github.com/h-nosaka/catwalk/mysql"
)

func SetEnv() {
	os.Setenv("RDB_TYPE", os.Getenv("MRDB_TYPE"))
	os.Setenv("RDB_USER", os.Getenv("MRDB_USER"))
	os.Setenv("RDB_PASSWORD", os.Getenv("MRDB_PASSWORD"))
	os.Setenv("RDB_HOST", os.Getenv("MRDB_HOST"))
	os.Setenv("RDB_DATABASE", os.Getenv("MRDB_DATABASE"))
}

func TestDump(t *testing.T) {
	SetEnv()
	base.Init()
	// Schema().Sql("./dump/dump.sql")
	mdb.NewSchemaFromDB().Sql("./dump/db.sql")
	t.Error("test")
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

func TestFixture(t *testing.T) {
	SetEnv()
	base.Init()
	schema := Schema()
	schema.Fixture("./fixtures")
}

func TestCreateDatabase(t *testing.T) {
	SetEnv()
	base.Init()
	Schema().CreateDatabase()
}

func TestDiff(t *testing.T) {
	SetEnv()
	base.Init()
	src := mdb.NewSchemaFromDB()
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
	src := mdb.NewSchemaFromDB()
	src.CreateSchema("./dump")
}

func TestTruncateTable(t *testing.T) {
	SetEnv()
	base.Init()
	for _, table := range mdb.NewSchemaFromDB().Tables {
		if err := base.DB.Exec(table.Drop()).Error; err != nil {
			t.Error(err)
		}
	}
}

func TestData(t *testing.T) {
	SetEnv()
	base.Init()
	account := models.Account{Email: "test@example.com"}
	if err := base.DB.Save(&account).Error; err != nil {
		t.Error(err)
	}
	t.Log(account)
}
