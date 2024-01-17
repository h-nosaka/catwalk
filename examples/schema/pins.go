package schema

import (
	"github.com/h-nosaka/catwalk"
)

func Pins() catwalk.Table {
	return catwalk.NewTable("app", "pins", catwalk.JsonCaseSnake, "ピンコードマスタ").SetColumns(
		catwalk.NewColumn("id", catwalk.DataTypeUUID, 0, false, "ID").SetDefault("uuid()").Done(),
		catwalk.NewColumn("pin", catwalk.DataTypeString, 128, false, "ピン").Done(),
		catwalk.NewColumn("usage", catwalk.DataTypeInt8, 3, false, "用途").Done(),
		catwalk.NewColumn("created_at", catwalk.DataTypeTimestamp, 0, false, "作成日").SetDefault("current_timestamp()").Done(),
		catwalk.NewColumn("updated_at", catwalk.DataTypeTimestamp, 0, false, "更新日").SetDefault("current_timestamp()").Done(),
	).SetIndexes(
		catwalk.NewIndex("PRIMARY", catwalk.IndexTypePrimary, "id"),
		catwalk.NewIndex("pins_pin_idx", catwalk.IndexTypeNotUnique, "pin"),
	).SetEnums().Done()
}
