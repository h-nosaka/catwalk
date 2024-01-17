package examples

import (
	"github.com/h-nosaka/catwalk"
)

func Pins() catwalk.Table {
	return catwalk.NewTable("app", "pins", catwalk.JsonCaseSnake, "ピンコードマスタ").SetDefaultColumns(
		catwalk.DataTypeUUID,
		catwalk.NewColumn("pin", catwalk.DataTypeString, 128, false, "ピン").Done(),
		catwalk.NewColumn("usage", catwalk.DataTypeInt8, 3, false, "用途").Done(),
	).SetDefaultIndexes(
		catwalk.NewIndex("pins_pin_idx", catwalk.IndexTypeNotUnique, "pin"),
	).SetEnums(
		catwalk.NewEnum("usage", catwalk.EnumTypeUint,
			catwalk.EnumValue{Key: "Onetime", Value: 1},
			catwalk.EnumValue{Key: "Pin", Value: 2},
		),
	).Done()
}
