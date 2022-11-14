package postgresql

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"strings"

	"github.com/freebitdx/fbfiber/context"
	"github.com/freebitdx/fbfiber/setting"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Load() (*context.Context, *PGSchema, *PGSchema) {
	c, err := context.New(false)
	if err != nil {
		panic(err)
	}
	schema := &PGSchema{}
	// sequence取得
	c.Database.Raw(`
		SELECT * FROM pg_catalog.pg_sequences WHERE schemaname = 'public'
	`).Scan(&schema.Sequences)
	// table取得
	schema.Tables = []PGTable{}
	tables := []PGTable{}
	c.Database.Raw(`
		SELECT pg_tables.*,pg_description.description as comment
			FROM pg_catalog.pg_tables
			JOIN pg_catalog.pg_class ON pg_class.relname = pg_tables.tablename
			LEFT JOIN pg_catalog.pg_description ON pg_class.oid = pg_description.objoid AND pg_description.objsubid = 0
			WHERE pg_tables.schemaname = 'public'
	`).Scan(&tables)
	for _, table := range tables {
		table.GetColumn(c.Database)
		schema.Tables = append(schema.Tables, table)
	}
	// index取得
	schema.Indexes = []PGIndex{}
	indexes := []PGIndex{}
	c.Database.Raw(`
		SELECT pg_indexes.*, table_constraints.constraint_type
			FROM pg_catalog.pg_indexes
 			LEFT JOIN information_schema.table_constraints ON pg_indexes.indexname = table_constraints.constraint_name
			WHERE schemaname = 'public'
	`).Scan(&indexes)
	for _, index := range indexes {
		index.GetColumn(c.Database)
		schema.Indexes = append(schema.Indexes, index)
	}
	// foreignkey取得
	c.Database.Raw(`
		SELECT table_constraints.*, f.attname as columnname, t.relname as reftable, r.attname as refcolumn
			FROM information_schema.table_constraints
			JOIN pg_catalog.pg_constraint ON table_constraints.constraint_name = pg_constraint.conname
			JOIN pg_catalog.pg_class as i ON pg_constraint.conindid = i.oid
			JOIN pg_catalog.pg_class as t ON pg_constraint.confrelid = t.oid
			JOIN pg_catalog.pg_attribute as f ON pg_constraint.conrelid = f.attrelid AND f.attnum = any(pg_constraint.conkey)
			JOIN pg_catalog.pg_attribute as r ON pg_constraint.confrelid = r.attrelid AND r.attnum = any(pg_constraint.confkey)
			WHERE table_schema = 'public' AND constraint_type = 'FOREIGN KEY';
	 `).Scan(&schema.Foreignkeys)

	file := PGSchema{}.Load(fmt.Sprintf("%s/schema/schema.yaml", c.Setting.RootPath))
	return c, schema, file
}

func Dump() {
	c, schema, _ := Load()
	defer c.Recover()
	schema.Yaml(fmt.Sprintf("%s/schema/dump.yaml", c.Setting.RootPath))
	schema.Sql(fmt.Sprintf("%s/schema/dump.sql", c.Setting.RootPath))
}

func Diff(filename ...string) string {
	c, schema, file := Load()
	defer c.Recover()
	buf := bytes.NewBuffer([]byte{})
	if filename != nil && len(filename[0]) > 0 {
		buf.WriteString(file.Diff(schema))
		if buf.Len() > 0 {
			fp, err := os.Create(fmt.Sprintf("%s/schema/%s.sql", c.Setting.RootPath, filename[0]))
			if err != nil {
				panic(err)
			}
			defer fp.Close()
			if _, err := fp.WriteString(buf.String()); err != nil {
				fmt.Printf("WriteString Error: %s\n", err.Error())
			}
		} else {
			fmt.Printf("No difference: %v byte\n", buf.Len())
		}
	} else {
		buf.WriteString(file.Diff(schema))
	}
	return buf.String()
}

func Run(filename ...string) {
	c, err := context.New(false)
	defer c.Recover()
	if err != nil {
		panic(err)
	}
	diff := ""
	if filename != nil && len(filename[0]) > 0 {
		fp, err := os.Open(fmt.Sprintf("%s/schema/%s.sql", c.Setting.RootPath, filename[0]))
		if err != nil {
			panic(err)
		}
		buf, err := ioutil.ReadAll(fp)
		if err != nil {
			panic(err)
		}
		diff = string(buf)
	} else {
		diff = Diff()
	}
	lines := strings.Split(diff, ";\n")
	count := 0
	for _, sql := range lines {
		if len(sql) > 10 {
			if err := c.Database.Exec(sql).Error; err != nil {
				panic(err)
			}
			count++
		}
	}
	fmt.Printf("Run comleted: %d sql\n", count)
}

func Model(path ...string) {
	c, err := context.New(false)
	if err != nil {
		panic(err)
	}
	if path == nil || len(path[0]) == 0 {
		path = []string{"models/dump"}
	}
	schema := PGSchema{}.Load(fmt.Sprintf("%s/schema/schema.yaml", c.Setting.RootPath))
	mapping := Mappings{}.Load(fmt.Sprintf("%s/schema/mapping.yaml", c.Setting.RootPath))
	for _, table := range schema.Tables {
		table.GoModel(fmt.Sprintf("%s/%s", c.Setting.RootPath, path[0]), &schema.Indexes, &schema.Foreignkeys, mapping)
	}
	if err := exec.Command("go", "fmt", fmt.Sprintf("%s/%s/...", c.Setting.RootPath, path[0])).Run(); err != nil {
		panic(err)
	}
	fmt.Printf("Model comleted: %d models\n", len(schema.Tables))
}

func CreateDatabase(name ...string) {
	settings := setting.Setting(true)
	dsn := fmt.Sprintf("user=%s password=%s dbname=%s host=%s sslmode=disable", settings.DBUser, settings.DBPassword, "postgres", settings.DBHost)
	db, err := gorm.Open(postgres.Open(dsn))
	if err != nil {
		panic(err)
	}
	if err := db.Exec(fmt.Sprintf("CREATE DATABASE %s OWNER = %s ENCODING = 'UTF8' TABLESPACE = pg_default", name[0], settings.DBUser)).Error; err != nil {
		panic(err)
	}
}
