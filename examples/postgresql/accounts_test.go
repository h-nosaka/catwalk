package schema

import (
	"testing"

	"github.com/h-nosaka/catwalk/base"
	db "github.com/h-nosaka/catwalk/postgresql"
)

func TestAccountsCreate(t *testing.T) {
	SetEnv()
	base.Init()
	tbl := Accounts()
	sql := tbl.Create()
	t.Log(sql)
}

func TestAccountsRun(t *testing.T) {
	SetEnv()
	base.Init()
	schema := &db.ISchema{
		Name: "app",
		Tables: []db.ITable{
			Accounts(),
		},
	}
	schema.Run()
	t.Error("test")
}
