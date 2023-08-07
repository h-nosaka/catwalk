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
		Columns: []db.IColumn{
			db.NewColumn("id", "bigint(20) unsigned", base.String("auto_increment"), nil, base.Bool(false), base.String("primary key"), nil),
			db.NewColumn("email", "varchar(256)", nil, nil, base.Bool(false), base.String("メールアドレス"), nil),
			db.NewColumn("hashed_password", "varchar(256)", nil, nil, base.Bool(false), base.String("ハッシュ化済みパスワード"), nil),
			db.NewColumn("salt", "varchar(8)", nil, nil, base.Bool(false), base.String("ソルト"), nil),
			db.NewColumn("code", "varchar(64)", nil, nil, base.Bool(false), base.String("表示ID"), nil),
			db.NewColumn("notification_id", "bigint(20) unsigned", nil, base.String("NULL"), base.Bool(true), base.String("notifications.id"), nil),
			db.NewColumn("role", "tinyint(3) unsigned", nil, base.String("NULL"), base.Bool(true), base.String("ロール"), nil),
			db.NewColumn("status", "tinyint(3) unsigned", nil, base.String("NULL"), base.Bool(true), base.String("ステータス"), nil),
			db.NewColumn("flags", "int(10) unsigned", nil, base.String("NULL"), base.Bool(true), base.String("フラグ"), nil),
			db.NewColumn("freezed_at", "timestamp", nil, base.String("NULL"), base.Bool(true), base.String("削除日"), nil),
			db.NewColumn("deleted_at", "timestamp", nil, base.String("NULL"), base.Bool(true), base.String("削除日"), nil),
			db.NewColumn("created_at", "timestamp", nil, base.String("current_timestamp()"), base.Bool(true), base.String("作成日"), nil),
			db.NewColumn("updated_at", "timestamp", nil, base.String("current_timestamp()"), base.Bool(true), base.String("更新日"), nil),
		},
		Indexes: []db.IIndex{
			db.NewIndex("PRIMARY", base.Bool(true), "id"),
			db.NewIndex("accounts_email_IDX", base.Bool(true), "email"),
			db.NewIndex("accounts_code_IDX", base.Bool(true), "code"),
		},
		Foreignkeys: []db.IForeignkey{},
		Enums:       []db.IEnum{},
		Methods:     []db.IMethod{},
		Relations:   []db.IRelation{},
	}
}
