package models_test

import (
	"os"
	"testing"

	"github.com/google/uuid"
	"github.com/h-nosaka/catwalk/base"
	"github.com/h-nosaka/catwalk/bps"
	"github.com/h-nosaka/catwalk/examples/mysql/models"
)

func SetEnv() {
	os.Setenv("RDB_TYPE", os.Getenv("MRDB_TYPE"))
	os.Setenv("RDB_USER", os.Getenv("MRDB_USER"))
	os.Setenv("RDB_PASSWORD", os.Getenv("MRDB_PASSWORD"))
	os.Setenv("RDB_HOST", os.Getenv("MRDB_HOST"))
	os.Setenv("RDB_DATABASE", os.Getenv("MRDB_DATABASE"))
}

func TestCreate(t *testing.T) {
	SetEnv()
	base.Init()
	item := models.Item{
		Id:    uuid.NewString(),
		Price: bps.New(1234.1234),
	}
	if err := base.DB.Save(&item).Error; err != nil {
		t.Error(err)
	}
}
