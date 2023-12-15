package schema

import (
	db "github.com/h-nosaka/catwalk/mysql"
)

func Schema() *db.ISchema {
	return &db.ISchema{
		Name: "app",
		Tables: []db.ITable{
			Accounts(),
			Pincodes(),
			AccountActivates(),
			AccountPincodes(),
			ActionLogs(),
			Item(),
		},
	}
}
