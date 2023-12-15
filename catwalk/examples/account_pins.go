package examples

import (
	"github.com/h-nosaka/catwalk/catwalk"
)

func AccountPins() catwalk.Table {
	return catwalk.NewTable("app", "account_pins", catwalk.JsonCaseSnake, "アカウントとピンの紐付け").SetDefaultColumns(
		catwalk.DataTypeUUID,
		catwalk.NewColumn("account_id", catwalk.DataTypeUUID, 0, false, "accounts.id").Done(),
		catwalk.NewColumn("pin_id", catwalk.DataTypeUUID, 0, false, "pins.id").Done(),
		catwalk.NewColumn("expired_at", catwalk.DataTypeTimestamp, 0, false, "PIN有効期限日時").Done(),
		catwalk.NewColumn("deleted_at", catwalk.DataTypeTimestamp, 0, true, "使用済み日時").Done(),
	).SetDefaultIndexes(
		catwalk.NewIndex("account_pins_account_id_IDX", catwalk.IndexTypeNotUnique, "account_id"),
		catwalk.NewIndex("account_pins_pin_id_IDX", catwalk.IndexTypeNotUnique, "pin_id"),
	).SetRelations(
		catwalk.NewRelation("account_pins_accounts_FK", "account_id", "accounts", "id", false, true),
		catwalk.NewRelation("account_pins_pins_FK", "pin_id", "pins", "id", false, true),
	).Done()
}
