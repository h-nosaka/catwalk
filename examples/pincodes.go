package schema

import (
	"github.com/h-nosaka/catwalk/base"
	db "github.com/h-nosaka/catwalk/mysql"
)

func Pincodes() db.ITable {
	return db.ITable{
		Schema:  "app",
		Name:    "pincodes",
		Comment: base.String("ピンコードマスタ"),
		Columns: db.DefaultColumn(
			db.NewColumn("pin", "varchar(6)", nil, nil, base.Bool(false), base.String("ピン"), nil),
		),
		Indexes: db.DefaultIndex(
			db.NewIndex("pincodes_pin_IDX", nil, "pin"),
		),
		Foreignkeys: []db.IForeignkey{},
		Enums:       []db.IEnum{},
		Methods:     []db.IMethod{},
	}
}
