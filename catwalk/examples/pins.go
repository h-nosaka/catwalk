package examples

import (
	"github.com/h-nosaka/catwalk/catwalk"
)

func Pins() catwalk.Table {
	return catwalk.NewTable("app", "pins", catwalk.JsonCaseSnake, "ピンコードマスタ").SetDefaultColumns(
		catwalk.DataTypeUUID,
		catwalk.NewColumn("pin", catwalk.DataTypeString, 128, false, "ピン").Done(),
		catwalk.NewColumn("usage", catwalk.DataTypeInt8, 0, false, "用途").Done(),
	).SetDefaultIndexes(
		catwalk.NewIndex("pins_pin_IDX", catwalk.IndexTypeNotUnique, "pin"),
	).SetEnums(
		catwalk.NewEnum("usage", catwalk.EnumTypeUint,
			catwalk.EnumValue{Key: "Onetime", Value: 1},
			catwalk.EnumValue{Key: "Pin", Value: 2},
		),
	).Done()
}
