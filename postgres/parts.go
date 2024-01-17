package postgres

var PARTS = map[string]string{
	"TableRename":                 "ALTER TABLE %s RENAME TO \"%s\";\n",  // schema.table, rename
	"TableDrop":                   "DROP TABLE IF EXISTS %s RESTRICT;\n", // schema.table
	"TableCreateOpen":             "CREATE TABLE %s (\n",                 // schema.table
	"TableCreateClose":            ");\n\n",
	"TableCreateCloseWithComment": ");\nCOMMENT ON TABLE %s IS '%s';\n\n",                         // schema.table, comment
	"ColumnAppend":                "\"%s\" %s",                                                    // column, _append
	"ColumnCreate":                "ALTER TABLE %s ADD COLUMN %s;\n",                              // schema.table, Append
	"ColumnSetComment":            "COMMENT ON COLUMN %s.\"%s\" IS '%s';\n",                       // schema.table, column, comment
	"ColumnSetModify":             "ALTER TABLE %s ALTER COLUMN \"%s\" %s %s;\n",                  // schema.table, Type, Append
	"ColumnRename":                "ALTER TABLE %s RENAME COLUMN \"%s\" TO \"%s\";\n",             // schema.table, column, rename
	"ColumnDrop":                  "ALTER TABLE %s DROP COLUMN \"%s\";\n",                         // schema.table, column
	"IndexCreatePrimary":          "ALTER TABLE %s ADD CONSTRAINT \"%s\" PRIMARY KEY (\"%s\");\n", // schema.table, index, columns
	"IndexCreateUnique":           "ALTER TABLE %s ADD CONSTRAINT \"%s\" UNIQUE (\"%s\");\n",      // schema.table, index, columns
	"IndexCreate":                 "CREATE INDEX IF NOT EXISTS \"%s\" ON %s (\"%s\");\n",          // index, schema.table, columns
	"IndexDropPrimary":            "ALTER TABLE %s DROP CONSTRAINT %s;\n",
	"IndexDropUnique":             "DROP UNIQUE INDEX IF EXISTS %s RESTRICT;\n",
	"IndexDrop":                   "DROP INDEX IF EXISTS %s RESTRICT;\n",
	"RelationCreate":              "ALTER TABLE ONLY %s ADD CONSTRAINT \"%s\" FOREIGN KEY (\"%s\") REFERENCES %s(\"%s\");\n", // schema.table, foreignkey, column, ref.schema.table, ref.column
	"RelationDrop":                "ALTER TABLE %s DROP CONSTRAINT \"%s\";\n",                                                // schema.table, foreignkey
	"PartitionCreate":             "",
	"PartitionCreateRange":        "",
	"PartitionCreateList":         "",
	"PartitionDrop":               "",
	"PartitionAppend":             "",
	"SchemaCreate":                "CREATE SCHEMA IF NOT EXISTS %s;\n\n",
}
