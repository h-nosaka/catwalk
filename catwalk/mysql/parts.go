package mysql

var PARTS = map[string]string{
	"TableRename":                 "RENAME TABLE %s TO %s;\n",            // tablename, rename
	"TableDrop":                   "DROP TABLE IF EXISTS %s RESTRICT;\n", // tablename
	"TableCreateOpen":             "CREATE TABLE %s (\n",                 // tablename
	"TableCreateClose":            ") ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;\n\n",
	"TableCreateCloseWithComment": ") ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='%s';\n\n", // comment
	"ColumnAppend":                "`%s` %s",                                                                              // column, _append
	"ColumnCreate":                "ALTER TABLE %s ADD %s;\n",                                                             // schema.table, Append
	"ColumnSetComment":            "",
	"ColumnSetModify":             "ALTER TABLE %s MODIFY COLUMN %s;\n",                     // schema.table, Append
	"ColumnRename":                "ALTER TABLE %s CHANGE %s;\n",                            // schema.table, ColumnAppend([column, rename], Append)
	"ColumnDrop":                  "ALTER TABLE %s DROP COLUMN `%s`;\n",                     // schema.table, column
	"IndexCreatePrimary":          "ALTER TABLE %s ADD CONSTRAINT %s PRIMARY KEY (`%s`);\n", // index, schema.table, columns
	"IndexCreateUnique":           "CREATE UNIQUE INDEX %s USING BTREE ON %s (`%s`);\n",     // index, schema.table, columns
	"IndexCreate":                 "CREATE INDEX %s USING BTREE ON %s (`%s`);\n",            // index, schema.table, columns
	"IndexDropPrimary":            "ALTER TABLE %s DROP PRIMARY KEY %s;\n",
	"IndexDropUnique":             "ALTER TABLE %s DROP INDEX %s;\n",
	"IndexDrop":                   "ALTER TABLE %s DROP INDEX %s;\n",
	"RelationCreate":              "ALTER TABLE %s ADD CONSTRAINT %s FOREIGN KEY %s(`%s`) REFERENCES %s(`%s`);\n", // schema.table, foreignkey, schema.table, column, ref.schema.table, ref.column
	"RelationDrop":                "ALTER TABLE %s DROP FOREIGN KEY %s;\n",                                        // schema.table, foreignkey
	"PartitionCreate":             "ALTER TABLE %s PARTITION BY %s(%s) (\n",                                       // schema.table, type, columns,
	"PartitionCreateRange":        "\tPARTITION %s VALUES LESS THAN (%s)",                                         // partition, value
	"PartitionCreateList":         "\tPARTITION %s VALUES IN (%s)",                                                // partition, value
	"PartitionDrop":               "ALTER TABLE %s DROP PARTITION %s;\n",                                          // shcema.table, partition
	"PartitionAppend":             "",
	"SchemaCreate":                "CREATE DATABASE IF NOT EXISTS %s CHARACTER SET utf8mb4 COLLATE utf8mb4_bin;\n\n",
}
