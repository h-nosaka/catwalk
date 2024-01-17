package schema

import (
	"github.com/h-nosaka/catwalk"
)

func ActionLogs() catwalk.Table {
	return catwalk.NewTable("app", "action_logs", catwalk.JsonCaseSnake, "アクションログ").SetColumns(
		catwalk.NewColumn("id", catwalk.DataTypeUUID, 0, false, "ID").SetDefault("uuid()").Done(),
		catwalk.NewColumn("uuid", catwalk.DataTypeString, 64, false, "UUID").Done(),
		catwalk.NewColumn("email", catwalk.DataTypeString, 256, false, "メールアドレス").Done(),
		catwalk.NewColumn("action_type", catwalk.DataTypeInt16, 6, false, "タイプ").Done(),
		catwalk.NewColumn("log", catwalk.DataTypeText16M, 0, false, "メッセージ").Done(),
		catwalk.NewColumn("recorded_at", catwalk.DataTypeTimestamp, 0, true, "実行日時").SetDefault("NULL").Done(),
		catwalk.NewColumn("created_at", catwalk.DataTypeTimestamp, 0, false, "作成日").SetDefault("current_timestamp()").Done(),
		catwalk.NewColumn("updated_at", catwalk.DataTypeTimestamp, 0, false, "更新日").SetDefault("current_timestamp()").Done(),
	).SetEnums().Done()
}
