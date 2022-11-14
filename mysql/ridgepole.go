package mysql

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"strings"

	"github.com/freebitdx/fbfiber/context"
	"github.com/freebitdx/fbfiber/setting"

	"gorm.io/driver/mysql"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Load() (*context.Context, *ISchema, *ISchema) {
	c, err := context.New(false)
	if err != nil {
		panic(err)
	}
	schema := &ISchema{}
	// table取得
	schema.Tables = []ITable{}
	tables := []ITable{}
	c.Database.Raw(fmt.Sprintf(GetTables, c.Setting.DBDatabase)).Scan(&tables)
	for _, table := range tables {
		table.GetColumn(c.Database)
		table.GetIndexes(c.Database)
		table.GetForeignkeys(c.Database)
		schema.Tables = append(schema.Tables, table)
	}

	file := ISchema{}.Load(fmt.Sprintf("%s/schema/schema.yaml", c.Setting.RootPath))

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
	gopath := "models/dump"
	dartpath := fmt.Sprintf("%s/tmp", c.Setting.RootPath)
	if len(path) > 0 && len(path[0]) > 0 {
		gopath = path[0]
	}
	if len(path) > 1 && len(path[1]) > 0 {
		dartpath = path[1]
	}

	schema := ISchema{}.Load(fmt.Sprintf("%s/schema/schema.yaml", c.Setting.RootPath))

	buf := bytes.NewBuffer([]byte{})
	buf.WriteString("import 'package:json_annotation/json_annotation.dart';\n")
	buf.WriteString("part 'models.g.dart';\n\n")
	for _, table := range schema.Tables {
		table.GoModel(fmt.Sprintf("%s/%s", c.Setting.RootPath, gopath))
		buf.WriteString(table.DartModel(dartpath))
	}
	fp, err := os.Create(fmt.Sprintf("%s/models.dart", dartpath))
	if err != nil {
		panic(err)
	}
	defer fp.Close()
	if _, err := fp.Write(buf.Bytes()); err != nil {
		fmt.Printf("WriteString Error: %s\n", err.Error())
	}
	fp.Close()
	if err := exec.Command("go", "fmt", fmt.Sprintf("%s/%s/...", c.Setting.RootPath, gopath)).Run(); err != nil {
		panic(err)
	}
	fmt.Printf("Model comleted: %d models\n", len(schema.Tables))
}

func CreateDatabase(name ...string) {
	settings := setting.Setting(true)
	var db *gorm.DB
	switch settings.DBType {
	case "postgres":
		dsn := fmt.Sprintf("user=%s password=%s dbname=%s host=%s sslmode=disable", settings.DBUser, settings.DBPassword, settings.DBDatabase, settings.DBHost)
		gdb, err := gorm.Open(postgres.Open(dsn))
		if err != nil {
			panic(err)
		}
		db = gdb
	case "mysql":
		dsn := fmt.Sprintf("%s:%s@tcp(%s:3306)/?charset=utf8mb4&parseTime=True&loc=Local", settings.DBUser, settings.DBPassword, settings.DBHost)
		gdb, err := gorm.Open(mysql.Open(dsn))
		if err != nil {
			panic(err)
		}
		db = gdb
	}
	if err := db.Exec(fmt.Sprintf(GetCreateDatabase, name[0])).Error; err != nil {
		panic(err)
	}
}
