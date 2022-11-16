package schema

import (
	"github.com/h-nosaka/catwalk/base"
	db "github.com/h-nosaka/catwalk/mysql"
)

func AccountActivates() db.ITable {
	return db.ITable{
		Schema:  "app",
		Name:    "account_activates",
		Comment: base.String("UUID管理マスタ"),
		Columns: []db.IColumn{
			db.NewColumn("id", "bigint(20) unsigned", base.String("auto_increment"), nil, base.Bool(false), base.String("primary key"), nil),
			db.NewColumn("account_id", "bigint(20) unsigned", nil, nil, base.Bool(false), base.String("accounts.id"), nil),
			db.NewColumn("uuid", "varchar(64)", nil, nil, base.Bool(false), base.String("UUID"), nil),
			db.NewColumn("pincode_id", "bigint(20) unsigned", nil, nil, base.Bool(false), base.String("pincodes.id"), nil),
			db.NewColumn("expired_at", "timestamp", nil, base.String("NULL"), base.Bool(true), base.String("PIN有効期限日時"), nil),
			db.NewColumn("activated_at", "timestamp", nil, base.String("NULL"), base.Bool(true), base.String("アクティベート日時"), nil),
			db.NewColumn("last_login_at", "timestamp", nil, base.String("NULL"), base.Bool(true), base.String("最終ログイン日時"), nil),
			db.NewColumn("created_at", "timestamp", nil, base.String("current_timestamp()"), base.Bool(true), base.String("作成日"), nil),
			db.NewColumn("updated_at", "timestamp", nil, base.String("current_timestamp()"), base.Bool(true), base.String("更新日"), nil),
		},
		Indexes: []db.IIndex{
			db.NewIndex("PRIMARY", base.Bool(true), "id"),
			db.NewIndex("account_activates_uuid_IDX", base.Bool(true), "uuid"),
			db.NewIndex("account_activates_account_id_IDX", base.Bool(false), "account_id"),
			db.NewIndex("account_activates_pincode_id_IDX", base.Bool(false), "pincode_id"),
		},
		Foreignkeys: []db.IForeignkey{},
		Enums:       []db.IEnum{},
		Methods:     []db.IMethod{},
		Relations:   []db.IRelation{},
	}
}
