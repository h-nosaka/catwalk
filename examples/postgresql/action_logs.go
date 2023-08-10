package schema

import (
	"github.com/h-nosaka/catwalk/base"
	db "github.com/h-nosaka/catwalk/postgresql"
)

func ActionLogs() db.ITable {
	return db.ITable{
		Schema:    "app",
		Name:      "action_logs",
		UseSchema: base.Bool(false),
		IsDB:      base.Bool(false),
		Comment:   base.String("アクションログ ESIDX"),
		Columns: db.DefaultColumn(
			db.NewColumn("_id", "varchar", 64, 0, nil, base.Bool(false), base.String("ID"), nil, nil),
			db.NewColumn("uuid", "varchar", 64, 0, nil, base.Bool(false), base.String("UUID"), nil, nil),
			db.NewColumn("email", "varchar", 256, 0, nil, base.Bool(true), base.String("メールアドレス"), nil, nil),
			db.NewColumn("action_type", "int2", 0, 0, nil, base.Bool(true), base.String("タイプ"), nil, nil),
			db.NewColumn("message", "text", 0, 0, nil, base.Bool(false), base.String("メッセージ"), nil, nil),
			db.NewColumn("recorded_at", "timestamp", 0, 0, nil, base.Bool(false), base.String("実行日時"), nil, nil),
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
