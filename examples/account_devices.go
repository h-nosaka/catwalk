package examples

import (
	"github.com/h-nosaka/catwalk"
)

func AccountDevices() catwalk.Table {
	return catwalk.NewTable("app", "account_devices", catwalk.JsonCaseSnake, "デバイス管理マスタ").SetDefaultColumns(
		catwalk.DataTypeUUID,
		catwalk.NewColumn("account_id", catwalk.DataTypeUUID, 0, false, "accounts.id").Done(),
		catwalk.NewColumn("uuid", catwalk.DataTypeString, 64, false, "デバイスID").Done(),
		catwalk.NewColumn("activated_at", catwalk.DataTypeTimestamp, 0, true, "アクティベート日時").Done(),
		catwalk.NewColumn("last_login_at", catwalk.DataTypeTimestamp, 0, true, "最終ログイン日時").Done(),
	).SetDefaultIndexes(
		catwalk.NewIndex("account_devices_uuid_idx", catwalk.IndexTypeNotUnique, "uuid"),
		catwalk.NewIndex("account_devices_account_id_idx", catwalk.IndexTypeNotUnique, "account_id"),
	).SetRelations(
		catwalk.NewRelation("account_devices_accounts_fk", "account_id", "accounts", "id", true, true),
	).Done()
}