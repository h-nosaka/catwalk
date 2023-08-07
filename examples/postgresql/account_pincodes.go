package schema

import (
	"github.com/h-nosaka/catwalk/base"
	db "github.com/h-nosaka/catwalk/postgresql"
)

func AccountPincodes() db.ITable {
	return db.ITable{
		Schema:    "app",
		Name:      "account_pincodes",
		UseSchema: base.Bool(false),
		Comment:   base.String("アカウントとピンコードの紐付け"),
		Columns: db.DefaultColumn("account_pincodes_seq",
			db.NewColumn("account_id", "int8", 0, 0, nil, base.Bool(false), base.String("accounts.id"), nil, nil),
			db.NewColumn("pincode_id", "int8", 0, 0, nil, base.Bool(false), base.String("pincodes.id"), nil, nil),
			db.NewColumn("expired_at", "timestamp", 0, 0, nil, nil, base.String("PIN有効期限日時"), nil, nil),
			db.NewColumn("deleted_at", "timestamp", 0, 0, nil, nil, base.String("使用済み日時"), nil, nil),
		),
		Indexes: db.DefaultIndex("account_pincodes_primary_idx",
			db.NewIndex("account_pincodes_account_id_IDX", nil, "account_id"),
			db.NewIndex("account_pincodes_pincode_id_IDX", nil, "pincode_id"),
		),
		Foreignkeys: []db.IForeignkey{
			db.NewFK("account_pincodes_accounts_FK", "account_id", "accounts", "id", false, false),
			db.NewFK("account_pincodes_pincodes_FK", "pincode_id", "pincodes", "id", false, false),
		},
		Sequences: []db.ISequence{
			db.NewSeq("account_pincodes_seq", 1, 1, 2147483647, 1),
		},
		Enums:   []db.IEnum{},
		Methods: []db.IMethod{},
	}
}
