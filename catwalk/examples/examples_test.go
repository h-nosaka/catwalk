package examples_test

import (
	"os"
	"os/exec"
	"testing"

	"github.com/h-nosaka/catwalk/base"
	"github.com/h-nosaka/catwalk/catwalk"
	"github.com/h-nosaka/catwalk/catwalk/examples"
)

func SetEnvM() {
	os.Setenv("RDB_TYPE", os.Getenv("MRDB_TYPE"))
	os.Setenv("RDB_USER", os.Getenv("MRDB_USER"))
	os.Setenv("RDB_PASSWORD", os.Getenv("MRDB_PASSWORD"))
	os.Setenv("RDB_HOST", os.Getenv("MRDB_HOST"))
	os.Setenv("RDB_DATABASE", os.Getenv("MRDB_DATABASE"))
	base.Init()
}

func SetEnvP() {
	os.Setenv("RDB_TYPE", os.Getenv("PRDB_TYPE"))
	os.Setenv("RDB_USER", os.Getenv("PRDB_USER"))
	os.Setenv("RDB_PASSWORD", os.Getenv("PRDB_PASSWORD"))
	os.Setenv("RDB_HOST", os.Getenv("PRDB_HOST"))
	os.Setenv("RDB_DATABASE", os.Getenv("PRDB_DATABASE"))
	base.Init()
}

func TestDumpMysql(t *testing.T) {
	SetEnvM()
	catwalk.NewSchemaFromDB().CreateSql("./dump/mysql_current.sql")
}
func TestDumpPostgres(t *testing.T) {
	SetEnvP()
	catwalk.NewSchemaFromDB().CreateSql("./dump/postgres_current.sql")
}

func TestCreateSqlMysql(t *testing.T) {
	SetEnvM()
	examples.Schema().CreateSql("./dump/mysql.sql")
}
func TestCreateSqlPostgres(t *testing.T) {
	SetEnvP()
	examples.Schema().CreateSql("./dump/postgres.sql")
}

func TestCreateYaml(t *testing.T) {
	SetEnvM()
	examples.Schema().CreateYaml("./dump/schema.yml")
}

func TestModel(t *testing.T) {
	SetEnvM()
	// for _, table := range examples.Schema().Tables {
	// 	table.CreateGoModel("./models")
	// }
	if err := exec.Command("go", "fmt", "./models/...").Run(); err != nil {
		t.Error(err)
	}
}

func TestFixture(t *testing.T) {
	SetEnvM()
	// examples.Schema().Fixture("./fixtures")
}

func TestCreateDatabase(t *testing.T) {
	SetEnvM()
	// examples.Schema().CreateDatabase()
}

func TestDiffMysql(t *testing.T) {
	SetEnvM()
	diff := examples.Schema().Diff(catwalk.NewSchemaFromDB())
	t.Log(diff)
	if len(diff) > 0 {
		t.Error("diff found")
	}
}

func TestDiffPostgres(t *testing.T) {
	SetEnvP()
	diff := examples.Schema().ForcePublic().Diff(catwalk.NewSchemaFromDB())
	t.Log(diff)
	if len(diff) > 0 {
		t.Error("diff found")
	}
}

func TestRunMysql(t *testing.T) {
	SetEnvM()
	examples.Schema().Run()
}

func TestRunPostgres(t *testing.T) {
	SetEnvP()
	examples.Schema().ForcePublic().Run()
}

func TestCreateSchema(t *testing.T) {
	SetEnvM()
	// catwalk.NewSchemaFromDB().CreateSchema("./dump")
}

func TestTruncateTable(t *testing.T) {
	SetEnvM()
	for _, table := range catwalk.NewSchemaFromDB().Tables {
		if err := base.DB.Exec(table.Drop()).Error; err != nil {
			t.Error(err)
		}
	}
}

// func TestData(t *testing.T) {
//  SetEnvM()
// 	account := models.Account{Email: "test@example.com"}
// 	if err := base.DB.Save(&account).Error; err != nil {
// 		t.Error(err)
// 	}
// 	t.Log(account)
// }
