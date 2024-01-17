package examples_test

import (
	"os"
	"testing"

	"github.com/google/uuid"
	"github.com/h-nosaka/catwalk"
	"github.com/h-nosaka/catwalk/base"
	"github.com/h-nosaka/catwalk/bps"
	"github.com/h-nosaka/catwalk/examples"
	"github.com/h-nosaka/catwalk/examples/fixtures"
	"github.com/h-nosaka/catwalk/examples/models"
	"github.com/h-nosaka/catwalk/money"
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

func TestCreateSqlMDB(t *testing.T) {
	SetEnvM()
	catwalk.NewSchemaFromDB().CreateSql("./dump/mysql_current.sql")
	examples.Schema().CreateSql("./dump/mysql.sql")
}
func TestCreateSqlPDB(t *testing.T) {
	SetEnvP()
	catwalk.NewSchemaFromDB().CreateSql("./dump/postgres_current.sql")
	examples.Schema().ForcePublic().CreateSql("./dump/postgres.sql")
}

func TestCreateYaml(t *testing.T) {
	SetEnvM()
	catwalk.NewSchemaFromDB().CreateYaml("./dump/schema_current.yml")
	examples.Schema().CreateYaml("./dump/schema.yml")
}

func TestModel(t *testing.T) {
	SetEnvM()
	examples.Schema().CreateModel("./models")
}

func TestFixture(t *testing.T) {
	SetEnvM()
	examples.Schema().CreateFixture("./fixtures")
}

func TestCreateSchema(t *testing.T) {
	SetEnvM()
	catwalk.NewSchemaFromDB().CreateSchema("./schema")
}

func TestCreateDatabase(t *testing.T) {
	SetEnvM()
	// examples.Schema().CreateDatabase()
}

func TestDiffMDB(t *testing.T) {
	SetEnvM()
	diff := examples.Schema().Diff(catwalk.NewSchemaFromDB())
	t.Log(diff)
	if len(diff) > 0 {
		t.Error("diff found")
	}
}

func TestDiffPDB(t *testing.T) {
	SetEnvP()
	diff := examples.Schema().ForcePublic().Diff(catwalk.NewSchemaFromDB())
	t.Log(diff)
	if len(diff) > 0 {
		t.Error("diff found")
	}
}

func TestRunMDB(t *testing.T) {
	SetEnvM()
	examples.Schema().Run()
}

func TestRunPDB(t *testing.T) {
	SetEnvP()
	examples.Schema().ForcePublic().Run()
}

func TestTruncateTableMDB(t *testing.T) {
	SetEnvM()
	for _, table := range catwalk.NewSchemaFromDB().Tables {
		if err := base.DB.Exec(table.Drop()).Error; err != nil {
			t.Error(err)
		}
	}
}

func TestTruncateTablePDB(t *testing.T) {
	SetEnvP()
	for _, table := range catwalk.NewSchemaFromDB().Tables {
		if err := base.DB.Exec(table.Drop()).Error; err != nil {
			t.Error(err)
		}
	}
}

func TestDataMDB(t *testing.T) {
	SetEnvM()
	item := fixtures.CreateItem(base.DB, func(model *models.Item) {
		model.Price = "1000"
		model.Enumuint = models.ItemEnumuintCreated
		model.Enumstring = models.ItemEnumstringActive
		model.Enumbitfield = models.ItemEnumbitfieldRead | models.ItemEnumbitfieldWrite
		model.BlobS = []byte{1}
		model.BlobM = []byte{2}
		model.BlobL = []byte{3}
		model.Bytes = []byte{4}
		model.Json = []byte("{}")
		model.Bps = bps.New(1000.123456789)
		model.Masked = "foo"
	})
	if item == nil {
		t.Error("作成エラー")
	}
	t.Log(base.ToPrettyJson(item))
	if item != nil && item.Bps != nil {
		t.Logf("bps: %s", item.Bps.Currency(money.VND, money.JaJP))
		t.Logf("masked: %s", item.Masked)
	}
	if err := base.DB.Find(&item).Error; err != nil {
		t.Error(err)
	}
	t.Log(base.ToPrettyJson(item))
	account := fixtures.CreateAccount(base.DB, func(model *models.Account) {
		model.Email = "test@example.com"
		model.NotificationId = uuid.NewString()
		model.AccountDevices = append(model.AccountDevices, *fixtures.AccountDevice(func(ad *models.AccountDevice) {
			ad.AccountId = model.Id
			ad.Uuid = "test"
		}))
	})
	if account == nil {
		t.Error("作成エラー")
	}
	t.Log(base.ToPrettyJson(account))
}

func TestDataPDB(t *testing.T) {
	SetEnvP()
	item := fixtures.CreateItem(base.DB, func(model *models.Item) {
		model.Price = "1000"
		model.Enumuint = models.ItemEnumuintCreated
		model.Enumstring = models.ItemEnumstringActive
		model.Enumbitfield = models.ItemEnumbitfieldRead | models.ItemEnumbitfieldWrite
		model.BlobS = []byte{1}
		model.BlobM = []byte{2}
		model.BlobL = []byte{3}
		model.Bytes = []byte{4}
		model.Json = []byte("{}")
		model.Bps = bps.New(1000.123456789)
		model.Masked = "foo"
	})
	if item == nil {
		t.Error("作成エラー")
	}
	t.Log(base.ToPrettyJson(item))
	if item != nil && item.Bps != nil {
		t.Logf("bps: %s", item.Bps.Currency(money.VND, money.JaJP))
		t.Logf("masked: %s", item.Masked)
	}
	if err := base.DB.Find(&item).Error; err != nil {
		t.Error(err)
	}
	t.Log(base.ToPrettyJson(item))
	account := fixtures.CreateAccount(base.DB, func(model *models.Account) {
		model.Email = "test@example.com"
		model.NotificationId = uuid.NewString()
		model.AccountDevices = append(model.AccountDevices, *fixtures.AccountDevice(func(ad *models.AccountDevice) {
			ad.AccountId = model.Id
			ad.Uuid = "test"
		}))
	})
	if account == nil {
		t.Error("作成エラー")
	}
	t.Log(base.ToPrettyJson(account))
}
