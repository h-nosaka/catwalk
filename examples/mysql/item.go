package schema

import (
	"github.com/h-nosaka/catwalk/base"
	db "github.com/h-nosaka/catwalk/mysql"
)

func Item() db.ITable {
	return db.ITable{
		Schema:  "app",
		Name:    "items",
		Comment: base.String("アイテムマスタ"),
		Columns: db.DefaultUuidColumn(
			db.NewColumn("price", "varchar(32)", nil, nil, base.Bool(false), base.String("価格"), nil),
		),
		Indexes:     db.DefaultIndex(),
		Foreignkeys: []db.IForeignkey{},
		Enums:       []db.IEnum{},
		Methods:     []db.IMethod{},
	}
}
