package schema

import (
	"github.com/h-nosaka/catwalk/base"
	db "github.com/h-nosaka/catwalk/mysql"
)

func ActionLogs() db.ITable {
	return db.ITable{
		Schema:  "app",
		Name:    "action_logs",
		IsDB:    base.Bool(false),
		Comment: base.String("アクションログ ESIDX"),
		Columns: db.DefaultColumn(
			db.NewColumn("_id", "varchar(64)", nil, nil, base.Bool(false), base.String("ID"), nil),
			db.NewColumn("uuid", "varchar(64)", nil, nil, base.Bool(false), base.String("UUID"), nil),
			db.NewColumn("email", "varchar(256)", nil, nil, base.Bool(true), base.String("メールアドレス"), nil),
			db.NewColumn("action_type", "int(3) unsigned", nil, nil, base.Bool(true), base.String("タイプ"), nil),
			db.NewColumn("message", "text", nil, nil, base.Bool(false), base.String("メッセージ"), nil),
			db.NewColumn("recorded_at", "timestamp", nil, nil, base.Bool(false), base.String("実行日時"), nil),
		)[1:],
		Enums: []db.IEnum{
			db.NewEnum("action_type", db.EnumTypeUint, map[string]interface{}{
				"RESUMED":    1,
				"INACTIVE":   2,
				"PAUSED":     3,
				"DETACHED":   4,
				"SAVEYOU":    11,
				"KINGOFTIME": 12,
				"KOTADMIN":   13,
				"GAROON":     14,
				"CLOUDMAIL":  15,
				"SLACK":      16,
			}),
		},
	}
}
