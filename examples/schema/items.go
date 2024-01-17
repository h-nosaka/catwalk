package schema

import (
	"github.com/h-nosaka/catwalk"
)

func Items() catwalk.Table {
	return catwalk.NewTable("app", "items", catwalk.JsonCaseSnake, "データバリエーション").SetColumns(
		catwalk.NewColumn("id", catwalk.DataTypeUUID, 0, false, "ID").SetDefault("uuid()").Done(),
		catwalk.NewColumn("price", catwalk.DataTypeString, 32, false, "価格").Done(),
		catwalk.NewColumn("int8", catwalk.DataTypeInt8, 4, false, "").Done(),
		catwalk.NewColumn("int16", catwalk.DataTypeInt16, 6, false, "").Done(),
		catwalk.NewColumn("int32", catwalk.DataTypeInt32, 11, false, "").Done(),
		catwalk.NewColumn("int64", catwalk.DataTypeInt64, 20, false, "").Done(),
		catwalk.NewColumn("uint8", catwalk.DataTypeUint8, 3, false, "").Done(),
		catwalk.NewColumn("uint16", catwalk.DataTypeUint16, 5, false, "").Done(),
		catwalk.NewColumn("uint32", catwalk.DataTypeUint32, 10, false, "").Done(),
		catwalk.NewColumn("uint64", catwalk.DataTypeUint64, 20, false, "").Done(),
		catwalk.NewColumn("float64", catwalk.DataTypeFloat64, 0, false, "").Done(),
		catwalk.NewColumn("string", catwalk.DataTypeString, 64, false, "").Done(),
		catwalk.NewColumn("fixstring", catwalk.DataTypeFixString, 64, false, "").Done(),
		catwalk.NewColumn("text_s", catwalk.DataTypeText64K, 0, false, "").Done(),
		catwalk.NewColumn("text_m", catwalk.DataTypeText16M, 0, false, "").Done(),
		catwalk.NewColumn("text_l", catwalk.DataTypeText4G, 0, false, "").Done(),
		catwalk.NewColumn("blob_s", catwalk.DataTypeBlob64K, 0, false, "").Done(),
		catwalk.NewColumn("blob_m", catwalk.DataTypeBlob16M, 0, false, "").Done(),
		catwalk.NewColumn("blob_l", catwalk.DataTypeBlob4G, 0, false, "").Done(),
		catwalk.NewColumn("bytes", catwalk.DataTypeUUID, 1, false, "").Done(),
		catwalk.NewColumn("json", catwalk.DataTypeText4G, 0, false, "").Done(),
		catwalk.NewColumn("timestamp", catwalk.DataTypeTimestamp, 0, false, "").SetDefault("current_timestamp()").SetExtra("on update current_timestamp()").Done(),
		catwalk.NewColumn("datetime", catwalk.DataTypeDatetime, 0, false, "").Done(),
		catwalk.NewColumn("enumuint", catwalk.DataTypeUint8, 3, false, "").Done(),
		catwalk.NewColumn("enumstring", catwalk.DataTypeString, 64, false, "").Done(),
		catwalk.NewColumn("enumbitfield", catwalk.DataTypeUint64, 20, false, "").Done(),
		catwalk.NewColumn("created_at", catwalk.DataTypeTimestamp, 0, false, "作成日").SetDefault("current_timestamp()").Done(),
		catwalk.NewColumn("updated_at", catwalk.DataTypeTimestamp, 0, false, "更新日").SetDefault("current_timestamp()").Done(),
		catwalk.NewColumn("bps", catwalk.DataTypeString, 64, false, "").Done(),
		catwalk.NewColumn("masked", catwalk.DataTypeString, 256, false, "").SetDefault("''").Done(),
	).SetIndexes(
		catwalk.NewIndex("PRIMARY", catwalk.IndexTypePrimary, "id"),
		catwalk.NewIndex("items_price_idx", catwalk.IndexTypeNotUnique, "price"),
	).SetEnums().Done()
}
