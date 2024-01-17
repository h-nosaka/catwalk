package examples

import (
	"github.com/h-nosaka/catwalk"
	"github.com/h-nosaka/catwalk/base"
)

func Schema() *catwalk.Schema {
	return catwalk.NewSchema(
		"app", catwalk.SchemaModeMySQL,
		Accounts(),
		Pins(),
		AccountDevices(),
		AccountPins(),
		ActionLogs(),
		Items(),
	)
}

type Item struct {
	// column
	Id     string `json:"id"` // ID
	Append string `json:"append"`
}

func Migration() error {
	defer base.Recover()
	// not nullのカラムを追加する
	beforeItem := Items()
	// 事前にnullのカラムを追加する
	beforeItem.Columns = append(beforeItem.Columns, catwalk.NewColumn("append", catwalk.DataTypeArrayString, 64, true, "").Done())
	schema := catwalk.NewSchema(
		"app", catwalk.SchemaModeMySQL,
		Accounts(),
		Pins(),
		AccountDevices(),
		AccountPins(),
		ActionLogs(),
		beforeItem,
	)
	schema.Run()
	// 追加したカラムに対して適切なデータを入れる
	items := []Item{}
	if err := base.DB.Find(&items).Error; err != nil {
		return err
	}
	for _, item := range items {
		item.Append = "foo" // 仮
		if err := base.DB.Save(&item).Error; err != nil {
			return err
		}
	}
	// not nullに制約を更新
	afterItem := Items()
	afterItem.Columns = append(beforeItem.Columns, catwalk.NewColumn("append", catwalk.DataTypeArrayString, 64, false, "").Done())
	schema = catwalk.NewSchema(
		"app", catwalk.SchemaModeMySQL,
		Accounts(),
		Pins(),
		AccountDevices(),
		AccountPins(),
		ActionLogs(),
		afterItem,
	)
	schema.Run()
	return nil
}
