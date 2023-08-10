package schema

import (
	"github.com/h-nosaka/catwalk/base"
	db "github.com/h-nosaka/catwalk/postgresql"
)

func Pincodes() db.ITable {
	return db.ITable{
		Schema:    "app",
		Name:      "pincodes",
		UseSchema: base.Bool(false),
		Comment:   base.String("ピンコードマスタ"),
		Columns: db.DefaultColumn(
			db.NewColumn("pin", "varchar", 6, 0, nil, base.Bool(false), base.String("ピン"), nil, nil),
		),
		Indexes: db.DefaultIndex("pincodes_primary_idx",
			db.NewIndex("pincodes_pin_IDX", nil, "pin"),
		),
		Foreignkeys: []db.IForeignkey{},
		Enums:       []db.IEnum{},
		Methods:     []db.IMethod{},
	}
}
