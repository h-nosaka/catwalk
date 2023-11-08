package postgresql

const (
	GetSequences = "SELECT * FROM pg_catalog.pg_sequences WHERE schemaname = 'public' AND sequencename = ?;"
	GetTables    = `
		SELECT pg_tables.*,pg_description.description as comment
		FROM pg_catalog.pg_tables
		JOIN pg_catalog.pg_class ON pg_class.relname = pg_tables.tablename
		LEFT JOIN pg_catalog.pg_description ON pg_class.oid = pg_description.objoid AND pg_description.objsubid = 0
		WHERE pg_tables.schemaname NOT IN ('pg_catalog','information_schema');
	`
	GetIndexes = `
		SELECT pg_indexes.*, table_constraints.constraint_type
		FROM pg_catalog.pg_indexes
		LEFT JOIN information_schema.table_constraints ON pg_indexes.indexname = table_constraints.constraint_name
		WHERE schemaname NOT IN ('pg_catalog','information_schema') AND tablename = ?;
	`
	GetIndexColumn = `
		SELECT pg_attribute.attname
		FROM pg_catalog.pg_attribute
		LEFT JOIN pg_catalog.pg_class ON pg_attribute.attrelid = pg_class.oid
		WHERE pg_class.relname = ?
		ORDER BY pg_attribute.attnum;
	`
	GetForeignkeys = `
		SELECT table_constraints.*, f.attname as columnname, t.relname as reftable, r.attname as refcolumn
		FROM information_schema.table_constraints
		JOIN pg_catalog.pg_constraint ON table_constraints.constraint_name = pg_constraint.conname
		JOIN pg_catalog.pg_class as i ON pg_constraint.conindid = i.oid
		JOIN pg_catalog.pg_class as t ON pg_constraint.confrelid = t.oid
		JOIN pg_catalog.pg_attribute as f ON pg_constraint.conrelid = f.attrelid AND f.attnum = any(pg_constraint.conkey)
		JOIN pg_catalog.pg_attribute as r ON pg_constraint.confrelid = r.attrelid AND r.attnum = any(pg_constraint.confkey)
		WHERE table_schema NOT IN ('pg_catalog','information_schema') AND table_name = ? AND constraint_type = 'FOREIGN KEY';
	`
	GetCreateTable    = ""
	GetCreateDatabase = "CREATE DATABASE %s OWNER = %s ENCODING = 'UTF8' TABLESPACE = pg_default"
	GetColumns        = `
		SELECT columns.*, pg_description.description as comment
		FROM information_schema.columns
		JOIN pg_catalog.pg_class ON pg_class.relname = ?
		JOIN pg_catalog.pg_attribute ON pg_class.oid = pg_attribute.attrelid AND columns.column_name = pg_attribute.attname
		LEFT JOIN pg_catalog.pg_description ON pg_attribute.attrelid = pg_description.objoid AND pg_attribute.attnum = pg_description.objsubid
		WHERE table_schema NOT IN ('pg_catalog','information_schema') AND table_name = ? order by ordinal_position;
	`
)
