package schema

import (
	"github.com/h-nosaka/catwalk"
)

func AccountDevices() catwalk.Table {
	return catwalk.NewTable("app", "account_devices", catwalk.JsonCaseSnake, "デバイス管理マスタ").SetColumns(
		catwalk.NewColumn("id", catwalk.DataTypeUUID, 0, false, "ID").SetDefault("uuid()").Done(),
		catwalk.NewColumn("account_id", catwalk.DataTypeUUID, 0, false, "accounts.id").Done(),
		catwalk.NewColumn("uuid", catwalk.DataTypeString, 64, false, "デバイスID").Done(),
		catwalk.NewColumn("activated_at", catwalk.DataTypeTimestamp, 0, true, "アクティベート日時").SetDefault("NULL").Done(),
		catwalk.NewColumn("last_login_at", catwalk.DataTypeTimestamp, 0, true, "最終ログイン日時").SetDefault("NULL").Done(),
		catwalk.NewColumn("created_at", catwalk.DataTypeTimestamp, 0, false, "作成日").SetDefault("current_timestamp()").Done(),
		catwalk.NewColumn("updated_at", catwalk.DataTypeTimestamp, 0, false, "更新日").SetDefault("current_timestamp()").Done(),
	).SetIndexes(
		catwalk.NewIndex("PRIMARY", catwalk.IndexTypePrimary, "id"),
		catwalk.NewIndex("account_devices_account_id_idx", catwalk.IndexTypeNotUnique, "account_id"),
		catwalk.NewIndex("account_devices_uuid_idx", catwalk.IndexTypeNotUnique, "uuid"),
	).SetRelations(
		catwalk.NewRelation("account_devices_accounts_fk", "account_id", "accounts", "id", false, true),
	).SetEnums().Done()
}
