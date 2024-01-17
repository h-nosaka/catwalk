package schema

import (
	"github.com/h-nosaka/catwalk"
)

func Accounts() catwalk.Table {
	return catwalk.NewTable("app", "accounts", catwalk.JsonCaseSnake, "アカウントマスタ").SetColumns(
		catwalk.NewColumn("id", catwalk.DataTypeUUID, 0, false, "ID").SetDefault("uuid()").Done(),
		catwalk.NewColumn("email", catwalk.DataTypeString, 256, false, "メールアドレス").Done(),
		catwalk.NewColumn("hashed_password", catwalk.DataTypeString, 256, false, "ハッシュ化済みパスワード").Done(),
		catwalk.NewColumn("salt", catwalk.DataTypeString, 8, false, "ソルト").Done(),
		catwalk.NewColumn("notification_id", catwalk.DataTypeUUID, 0, false, "notifications.id").Done(),
		catwalk.NewColumn("role", catwalk.DataTypeUint32, 10, false, "ロール").Done(),
		catwalk.NewColumn("state", catwalk.DataTypeUint8, 3, false, "ステータス").Done(),
		catwalk.NewColumn("flags", catwalk.DataTypeUint64, 20, false, "フラグ").Done(),
		catwalk.NewColumn("freezed_at", catwalk.DataTypeTimestamp, 0, true, "凍結日").SetDefault("NULL").Done(),
		catwalk.NewColumn("deleted_at", catwalk.DataTypeTimestamp, 0, true, "削除日").SetDefault("NULL").Done(),
		catwalk.NewColumn("created_at", catwalk.DataTypeTimestamp, 0, false, "作成日").SetDefault("current_timestamp()").Done(),
		catwalk.NewColumn("updated_at", catwalk.DataTypeTimestamp, 0, false, "更新日").SetDefault("current_timestamp()").Done(),
	).SetIndexes(
		catwalk.NewIndex("PRIMARY", catwalk.IndexTypePrimary, "id"),
		catwalk.NewIndex("accounts_multi_idx", catwalk.IndexTypeUnique, "email", "deleted_at"),
		catwalk.NewIndex("accounts_email_idx", catwalk.IndexTypeNotUnique, "email"),
		catwalk.NewIndex("accounts_state_idx", catwalk.IndexTypeNotUnique, "state"),
	).SetEnums().Done()
}
