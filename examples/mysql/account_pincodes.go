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
		Columns: db.DefaultColumn(
			db.NewColumn("account_id", "uuid", nil, nil, base.Bool(false), base.String("accounts.id"), nil),
			db.NewColumn("pincode_id", "bigint(20) unsigned", nil, nil, base.Bool(false), base.String("pincodes.id"), nil),
			db.NewColumn("expired_at", "timestamp", nil, nil, nil, base.String("PIN有効期限日時"), nil),
			db.NewColumn("deleted_at", "timestamp", nil, nil, nil, base.String("使用済み日時"), nil),
		),
		Indexes: db.DefaultIndex(
			db.NewIndex("account_pincodes_account_id_IDX", nil, "account_id"),
			db.NewIndex("account_pincodes_pincode_id_IDX", nil, "pincode_id"),
		),
		Foreignkeys: []db.IForeignkey{
			db.NewFK("account_pincodes_accounts_FK", "account_id", "accounts", "id", false, false),
			db.NewFK("account_pincodes_pincodes_FK", "pincode_id", "pincodes", "id", false, false),
		},
		Enums:   []db.IEnum{},
		Methods: []db.IMethod{},
	}
}
