package examples

import (
	"github.com/h-nosaka/catwalk"
)

func Accounts() catwalk.Table {
	return catwalk.NewTable("app", "accounts", catwalk.JsonCaseSnake, "アカウントマスタ").SetDefaultColumns(
		catwalk.DataTypeUUID,
		catwalk.NewColumn("email", catwalk.DataTypeString, 256, false, "メールアドレス").Done(),
		catwalk.NewColumn("hashed_password", catwalk.DataTypeString, 256, false, "ハッシュ化済みパスワード").Done(),
		catwalk.NewColumn("salt", catwalk.DataTypeString, 8, false, "ソルト").Done(),
		catwalk.NewColumn("notification_id", catwalk.DataTypeUUID, 0, false, "notifications.id").Done(),
		catwalk.NewColumn("role", catwalk.DataTypeUint32, 10, false, "ロール").Done(),
		catwalk.NewColumn("state", catwalk.DataTypeUint8, 3, false, "ステータス").Done(),
		catwalk.NewColumn("flags", catwalk.DataTypeUint64, 20, false, "フラグ").Done(),
		catwalk.NewColumn("freezed_at", catwalk.DataTypeTimestamp, 0, true, "凍結日").Done(),
		catwalk.NewColumn("deleted_at", catwalk.DataTypeTimestamp, 0, true, "削除日").Done(),
	).SetDefaultIndexes(
		catwalk.NewIndex("accounts_email_idx", catwalk.IndexTypeNotUnique, "email"),
		catwalk.NewIndex("accounts_state_idx", catwalk.IndexTypeNotUnique, "state"),
		catwalk.NewIndex("accounts_multi_idx", catwalk.IndexTypeUnique, "email", "deleted_at"),
	).SetRelations(
		catwalk.NewRelation("", "id", "account_devices", "account_id", false, false),
	).SetEnums(
		catwalk.NewEnum("role", catwalk.EnumTypeBitfield,
			catwalk.EnumValue{Key: "Viewer", Value: 1},
			catwalk.EnumValue{Key: "Writer", Value: 2},
			catwalk.EnumValue{Key: "Manager", Value: 3},
		),
		catwalk.NewEnum("state", catwalk.EnumTypeUint,
			catwalk.EnumValue{Key: "Created", Value: 1},
			catwalk.EnumValue{Key: "Activated", Value: 2},
			catwalk.EnumValue{Key: "Freezed", Value: 8},
			catwalk.EnumValue{Key: "Deleted", Value: 9},
		),
	).Done()
}
