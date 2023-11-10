package schema

import (
	"github.com/h-nosaka/catwalk/base"
	db "github.com/h-nosaka/catwalk/mysql"
)

func Accounts() db.ITable {
	return db.ITable{
		Schema:  "app",
		Name:    "accounts",
		Comment: base.String("アカウントマスタ"),
		Columns: db.DefaultUuidColumn(
			db.NewColumn("email", "varchar(256)", nil, nil, base.Bool(false), base.String("メールアドレス"), nil),
			db.NewColumn("hashed_password", "varchar(256)", nil, nil, base.Bool(false), base.String("ハッシュ化済みパスワード"), nil),
			db.NewColumn("salt", "varchar(8)", nil, nil, base.Bool(false), base.String("ソルト"), nil),
			db.NewColumn("code", "varchar(64)", nil, nil, base.Bool(false), base.String("表示ID"), nil),
			db.NewColumn("notification_id", "bigint(20) unsigned", nil, nil, nil, base.String("notifications.id"), nil),
			db.NewColumn("role", "tinyint(3) unsigned", nil, nil, nil, base.String("ロール"), nil),
			db.NewColumn("status", "tinyint(3) unsigned", nil, nil, nil, base.String("ステータス"), nil),
			db.NewColumn("flags", "int(10) unsigned", nil, nil, nil, base.String("フラグ"), nil),
			db.NewColumn("freezed_at", "timestamp", nil, nil, nil, base.String("削除日"), nil),
			db.NewColumn("deleted_at", "timestamp", nil, nil, nil, base.String("削除日"), nil),
		),
		Indexes: db.DefaultIndex(
			db.NewIndex("accounts_email_IDX", nil, "email"),
			db.NewIndex("accounts_code_IDX", nil, "code"),
			db.NewIndex("accounts_multi_IDX", base.Bool(true), "code", "email"),
		),
		Foreignkeys: []db.IForeignkey{
			db.NewFK("accounts_account_activates_FK", "id", "account_activates", "account_id", false, true),
		},
		Enums: []db.IEnum{
			db.NewEnum("role", db.EnumTypeBitfield, map[string]interface{}{
				"Viewer":  1,
				"Writer":  2,
				"Manager": 3,
			}),
			db.NewEnum("status", db.EnumTypeUint, map[string]interface{}{
				"Created":   0,
				"Activated": 1,
				"Freezed":   8,
				"Deleted":   9,
			}),
		},
		Methods: []db.IMethod{
			{
				Method: `func (p *Account) Auth(db *gorm.DB, email string) bool {
  db.Where("Email = ? and status = ?", email, AccountStatusActivated).First(&p)
  return p.Email == email
}`,
			},
		},
	}
}
