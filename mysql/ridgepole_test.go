package mysql

import (
	"fmt"
	"testing"

	"github.com/freebitdx/fbfiber/context"
)

func TestDump(t *testing.T) {
	c, err := context.New(true)
	if err != nil {
		panic(err)
	}
	defer c.Recover()
	Dump()
}

func TestDiff(t *testing.T) {
	c, err := context.New(true)
	if err != nil {
		panic(err)
	}
	defer c.Recover()
	t.Error(Diff("migrate_test"))
}

func TestRun(t *testing.T) {
	c, err := context.New(true)
	if err != nil {
		panic(err)
	}
	defer c.Recover()
	Run("migrate_test")
}

func TestModel(t *testing.T) {
	c, err := context.New(true)
	if err != nil {
		panic(err)
	}
	defer c.Recover()
	Model("models", fmt.Sprintf("%s/frontend/demo_flutter/lib/molecules", c.Setting.RootPath))
}

func TestCreateDatabase(t *testing.T) {
	c, err := context.New(false)
	if err != nil {
		panic(err)
	}
	defer c.Recover()
	CreateDatabase("app_test")
}
