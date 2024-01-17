package examples

import (
	"github.com/h-nosaka/catwalk"
)

func ActionLogs() catwalk.Table {
	return catwalk.NewTable("app", "action_logs", catwalk.JsonCaseSnake, "アクションログ").SetDefaultColumns(
		catwalk.DataTypeUUID,
		catwalk.NewColumn("uuid", catwalk.DataTypeString, 64, false, "UUID").Done(),
		catwalk.NewColumn("email", catwalk.DataTypeString, 256, false, "メールアドレス").Done(),
		catwalk.NewColumn("action_type", catwalk.DataTypeInt16, 6, false, "タイプ").Done(),
		catwalk.NewColumn("log", catwalk.DataTypeText16M, 0, false, "メッセージ").Done(),
		catwalk.NewColumn("recorded_at", catwalk.DataTypeTimestamp, 0, true, "実行日時").Done(),
	).SetEnums(
		catwalk.NewEnum("action_type", catwalk.EnumTypeUint,
			catwalk.EnumValue{Key: "RESUMED", Value: 1},
			catwalk.EnumValue{Key: "INACTIVE", Value: 2},
			catwalk.EnumValue{Key: "PAUSED", Value: 3},
			catwalk.EnumValue{Key: "DETACHED", Value: 4},
			catwalk.EnumValue{Key: "SAVEYOU", Value: 11},
			catwalk.EnumValue{Key: "KINGOFTIME", Value: 12},
			catwalk.EnumValue{Key: "KOTADMIN", Value: 13},
			catwalk.EnumValue{Key: "GAROON", Value: 14},
			catwalk.EnumValue{Key: "CLOUDMAIL", Value: 15},
			catwalk.EnumValue{Key: "SLACK", Value: 16},
		),
	).Done()
}
