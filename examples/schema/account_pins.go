package schema

import (
	"github.com/h-nosaka/catwalk"
)

func AccountPins() catwalk.Table {
	return catwalk.NewTable("app", "account_pins", catwalk.JsonCaseSnake, "アカウントとピンの紐付け").SetColumns(
		catwalk.NewColumn("id", catwalk.DataTypeUUID, 0, false, "ID").SetDefault("uuid()").Done(),
		catwalk.NewColumn("account_id", catwalk.DataTypeUUID, 0, false, "accounts.id").Done(),
		catwalk.NewColumn("pin_id", catwalk.DataTypeUUID, 0, false, "pins.id").Done(),
		catwalk.NewColumn("expired_at", catwalk.DataTypeTimestamp, 0, true, "PIN有効期限日時").SetDefault("NULL").Done(),
		catwalk.NewColumn("deleted_at", catwalk.DataTypeTimestamp, 0, true, "使用済み日時").SetDefault("NULL").Done(),
		catwalk.NewColumn("created_at", catwalk.DataTypeTimestamp, 0, false, "作成日").SetDefault("current_timestamp()").Done(),
		catwalk.NewColumn("updated_at", catwalk.DataTypeTimestamp, 0, false, "更新日").SetDefault("current_timestamp()").Done(),
	).SetIndexes(
		catwalk.NewIndex("PRIMARY", catwalk.IndexTypePrimary, "id"),
		catwalk.NewIndex("account_pins_account_id_idx", catwalk.IndexTypeNotUnique, "account_id"),
		catwalk.NewIndex("account_pins_pin_id_idx", catwalk.IndexTypeNotUnique, "pin_id"),
	).SetRelations(
		catwalk.NewRelation("account_pins_accounts_fk", "account_id", "accounts", "id", false, true),
		catwalk.NewRelation("account_pins_pins_fk", "pin_id", "pins", "id", false, true),
	).SetEnums().Done()
}
