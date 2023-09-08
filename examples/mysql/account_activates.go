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
		Columns: db.DefaultColumn(
			db.NewColumn("account_id", "uuid", nil, nil, base.Bool(false), base.String("accounts.id"), nil),
			db.NewColumn("uuid", "varchar(64)", nil, nil, base.Bool(false), base.String("UUID"), nil),
			db.NewColumn("pincode_id", "bigint(20) unsigned", nil, nil, base.Bool(false), base.String("pincodes.id"), nil),
			db.NewColumn("expired_at", "timestamp", nil, nil, nil, base.String("PIN有効期限日時"), nil),
			db.NewColumn("activated_at", "timestamp", nil, nil, nil, base.String("アクティベート日時"), nil),
			db.NewColumn("last_login_at", "timestamp", nil, nil, nil, base.String("最終ログイン日時"), nil),
		),
		Indexes: db.DefaultIndex(
			db.NewIndex("account_activates_uuid_IDX", nil, "uuid"),
			db.NewIndex("account_activates_account_id_IDX", base.Bool(false), "account_id"),
			db.NewIndex("account_activates_pincode_id_IDX", base.Bool(false), "pincode_id"),
		),
		Foreignkeys: []db.IForeignkey{
			db.NewFK("account_activates_accounts_FK", "account_id", "accounts", "id", false, false),
			db.NewFK("account_activates_pincodes_FK", "pincode_id", "pincodes", "id", false, false),
		},
		Enums: []db.IEnum{},
		Methods: []db.IMethod{
			{
				Method: `func (p *AccountActivate) Active(db *gorm.DB, uuid string) error {
	return db.Joins("Account").Where("uuid = ? and activated_at is not NULL and Account.status = ?", uuid, AccountStatusActivated).First(&p).Error
}`,
			},
		},
	}
}
