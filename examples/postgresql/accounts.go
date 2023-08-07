package schema

import (
	"github.com/h-nosaka/catwalk/base"
	db "github.com/h-nosaka/catwalk/postgresql"
)

func Accounts() db.ITable {
	return db.ITable{
		Schema:    "app",
		Name:      "accounts",
		UseSchema: base.Bool(false),
		Comment:   base.String("アカウントマスタ"),
		Columns: db.DefaultColumn("accounts_seq",
			db.NewColumn("email", "varchar", 256, 0, nil, base.Bool(false), base.String("メールアドレス"), nil, nil),
			db.NewColumn("hashed_password", "varchar", 256, 0, nil, base.Bool(false), base.String("ハッシュ化済みパスワード"), nil, nil),
			db.NewColumn("salt", "varchar", 8, 0, nil, base.Bool(false), base.String("ソルト"), nil, nil),
			db.NewColumn("code", "varchar", 64, 0, nil, base.Bool(false), base.String("表示ID"), nil, nil),
			db.NewColumn("notification_id", "int8", 0, 0, nil, nil, base.String("notifications.id"), nil, nil),
			db.NewColumn("role", "int2", 0, 0, nil, nil, base.String("ロール"), nil, nil),
			db.NewColumn("status", "int2", 0, 0, nil, nil, base.String("ステータス"), nil, nil),
			db.NewColumn("flags", "int4", 0, 0, nil, nil, base.String("フラグ"), nil, nil),
			db.NewColumn("freezed_at", "timestamp", 0, 6, nil, nil, base.String("凍結日"), nil, nil),
			db.NewColumn("deleted_at", "timestamp", 0, 6, nil, nil, base.String("削除日"), nil, nil),
		),
		Indexes: db.DefaultIndex("accounts_primary_idx",
			db.NewIndex("accounts_email_IDX", nil, "email"),
			db.NewIndex("accounts_code_IDX", nil, "code"),
		),
		Foreignkeys: []db.IForeignkey{
			// db.NewFK("accounts_account_activates_FK", "id", "account_activates", "account_id", false, true),
		},
		Sequences: []db.ISequence{
			db.NewSeq("accounts_seq", 1, 1, 2147483647, 1),
		},
		Enums: []db.IEnum{
			db.NewEnum("role", db.EnumTypeBitfield, map[string]interface{}{
				"Viewer":  1,
				"Writer":  1,
				"Manager": 1,
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
