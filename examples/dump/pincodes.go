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
		Columns: []db.IColumn{
			db.NewColumn("id", "bigint(20) unsigned", base.String("auto_increment"), nil, base.Bool(false), base.String("primary key"), nil),
			db.NewColumn("pin", "varchar(6)", nil, nil, base.Bool(false), base.String("ピン"), nil),
			db.NewColumn("created_at", "timestamp", nil, base.String("current_timestamp()"), base.Bool(true), base.String("作成日"), nil),
			db.NewColumn("updated_at", "timestamp", nil, base.String("current_timestamp()"), base.Bool(true), base.String("更新日"), nil),
		},
		Indexes: []db.IIndex{
			db.NewIndex("PRIMARY", base.Bool(true), "id"),
			db.NewIndex("pincodes_pin_IDX", base.Bool(true), "pin"),
		},
		Foreignkeys: []db.IForeignkey{},
		Enums:       []db.IEnum{},
		Methods:     []db.IMethod{},
		Relations:   []db.IRelation{},
	}
}
