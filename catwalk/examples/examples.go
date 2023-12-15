package examples

import (
	"github.com/h-nosaka/catwalk/catwalk"
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
