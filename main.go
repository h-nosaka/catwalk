package main

import (
	"flag"
	"fmt"

	"github.com/h-nosaka/catwalk/base"
	"github.com/h-nosaka/catwalk/mysql"
)

func main() {
	var (
		config string
		output string
		models string
		name   string
		help   bool
	)
	flag.StringVar(&config, "config", "./schema/schema.yaml", "schame yaml file path")
	flag.StringVar(&config, "c", "./schema/schema.yaml", "schame yaml file path")
	flag.StringVar(&output, "output", "./schema/dump", "dump file output directory path")
	flag.StringVar(&output, "o", "./schema/dump", "dump file output directory path")
	flag.StringVar(&models, "models", "./models", "models file output directory path")
	flag.StringVar(&models, "m", "./models", "models file output directory path")
	flag.StringVar(&name, "name", "app", "create databse name")
	flag.StringVar(&name, "n", "app", "create databse name")
	flag.BoolVar(&help, "help", false, "command help")
	flag.BoolVar(&help, "h", false, "command help")
	flag.Parse()
	if help {
		HelpCmd(flag.Arg(0))
		return
	}
	switch flag.Arg(0) {
	case "help":
		Help()
	case "dump":
		Dump(output)
	case "diff":
		Diff(config)
	case "run":
		Run(config)
	case "model":
		Model(config, models)
	case "init":
		CreateDatabase(name)
	default:
		Help()
	}
}

func Help() {
	fmt.Print("Usage:\n\ncatwalk <command> [arguments]\n\n")
	fmt.Print("The commands are:\n\n")
	fmt.Print("\tdump      craete dump.sql and dump.yaml\n")
	fmt.Print("\tdiff      show differential SQL\n")
	fmt.Print("\trun       run database migration\n")
	fmt.Print("\tmodel     create golang models\n")
	fmt.Print("\tinit      create database\n")
	fmt.Print("\n")
}

func HelpCmd(cmd string) {
	switch cmd {
	case "dump":
		fmt.Print("Usage: catwalk dump -output=dir\n")
	case "diff":
		fmt.Print("Usage: catwalk diff -config=filepath [output]\n")
	case "run":
		fmt.Print("Usage: catwalk run -config=filepath\n")
	case "model":
		fmt.Print("Usage: catwalk model -config=filepath -models=dir\n")
	case "init":
		fmt.Print("Usage: catwalk init -name=databasename\n")
	default:
		fmt.Print("command not found\n")
	}
}

func IsMySQL() bool {
	return base.GetEnv("RDB_TYPE", "") == "mysql"
}

func IsPostgres() bool {
	return base.GetEnv("RDB_TYPE", "") == "postgres"
}

func Dump(output string) {
	if IsMySQL() {
		schema := mysql.NewSchemaFromDB()
		schema.Sql(fmt.Sprintf("%s/dump.sql", output))
		schema.Yaml(fmt.Sprintf("%s/dump.yaml", output))
	}
	// if IsPostgres() {
	// 	postgresql.Dump()
	// }
}

func Diff(yamlpath string) string {
	if IsMySQL() {
		return mysql.NewSchema(yamlpath).Diff(mysql.NewSchemaFromDB())
	}
	// if IsPostgres() {
	// 	return postgresql.Diff(filename...)
	// }
	return ""
}

func Run(yamlpath string) {
	if IsMySQL() {
		mysql.NewSchema(yamlpath).Run()
	}
	// if IsPostgres() {
	// 	postgresql.Run(filename...)
	// }
}

func Model(yamlpath string, path ...string) {
	if IsMySQL() {
		mysql.NewSchema(yamlpath).Model(path...)
	}
	// if IsPostgres() {
	// 	postgresql.Model(path...)
	// }
}

func CreateDatabase(name ...string) {
	if IsMySQL() {
		(&mysql.ISchema{}).CreateDatabase(name...)
	}
	// if IsPostgres() {
	// 	postgresql.CreateDatabase(name...)
	// }
}
