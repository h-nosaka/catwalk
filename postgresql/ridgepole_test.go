package postgresql

import (
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
	Diff("migrate_test")
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
	Model("models")
}
