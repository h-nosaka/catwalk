package schema

import (
	"github.com/h-nosaka/catwalk/base"
	db "github.com/h-nosaka/catwalk/mysql"
)

func AccountPincodes() db.ITable {
	return db.ITable{
		Schema:  "app",
		Name:    "account_pincodes",
		Comment: base.String("アカウントとピンコードの紐付け"),
		Columns: []db.IColumn{
			db.NewColumn("id", "bigint(20) unsigned", base.String("auto_increment"), nil, base.Bool(false), base.String("primary key"), nil),
			db.NewColumn("account_id", "bigint(20) unsigned", nil, nil, base.Bool(false), base.String("accounts.id"), nil),
			db.NewColumn("pincode_id", "bigint(20) unsigned", nil, nil, base.Bool(false), base.String("pincodes.id"), nil),
			db.NewColumn("expired_at", "timestamp", nil, base.String("NULL"), base.Bool(true), base.String("PIN有効期限日時"), nil),
			db.NewColumn("deleted_at", "timestamp", nil, base.String("NULL"), base.Bool(true), base.String("使用済み日時"), nil),
			db.NewColumn("created_at", "timestamp", nil, base.String("current_timestamp()"), base.Bool(true), base.String("作成日"), nil),
			db.NewColumn("updated_at", "timestamp", nil, base.String("current_timestamp()"), base.Bool(true), base.String("更新日"), nil),
		},
		Indexes: []db.IIndex{
			db.NewIndex("PRIMARY", base.Bool(true), "id"),
			db.NewIndex("account_pincodes_account_id_IDX", base.Bool(true), "account_id"),
			db.NewIndex("account_pincodes_pincode_id_IDX", base.Bool(true), "pincode_id"),
		},
		Foreignkeys: []db.IForeignkey{},
		Enums:       []db.IEnum{},
		Methods:     []db.IMethod{},
		Relations:   []db.IRelation{},
	}
}
