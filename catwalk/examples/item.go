package examples

import "github.com/h-nosaka/catwalk/catwalk"

func Items() catwalk.Table {
	return catwalk.NewTable("app", "items", catwalk.JsonCaseSnake, "データバリエーション").SetDefaultColumns(
		catwalk.DataTypeUUID,
		catwalk.NewColumn("price", catwalk.DataTypeString, 32, false, "価格").Done(),
	).SetDefaultIndexes(
		catwalk.NewIndex("items_price_IDX", catwalk.IndexTypeNotUnique, "price"),
	).SetEnums(
		catwalk.NewEnum("usage", catwalk.EnumTypeUint,
			catwalk.EnumValue{Key: "Onetime", Value: 1},
			catwalk.EnumValue{Key: "Pin", Value: 2},
		),
	).Done()
}
