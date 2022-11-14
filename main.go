package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/freebitdx/ridgepole/mysql"
	"github.com/freebitdx/ridgepole/postgresql"
)

func main() {
	flag.Parse()
	switch flag.Arg(0) {
	case "help":
		help()
	case "dump":
		dump(flag.Args())
	case "diff":
		diff(flag.Args())
	case "run":
		run(flag.Args())
	case "model":
		model(flag.Args())
	case "init":
		initDatabase(flag.Args())
	default:
		help()
	}
}

func help() {
	fmt.Print("Usage:\n\nridgepole <command> [arguments]\n\n")
	fmt.Print("The commands are:\n\n")
	fmt.Print("\tdump      craete dump.sql and dump.yaml\n")
	fmt.Print("\tdiff      create *.sql and show\n")
	fmt.Print("\trun       run database migration\n")
	fmt.Print("\tmodel     create model\n")
	fmt.Print("\tinit      create database\n")
	fmt.Print("\n")
}

func dump(args []string) {
	if len(args) > 1 {
		switch args[1] {
		case "--help", "-h", "help":
			fmt.Print("Usage: ridgepole dump\n")
			return
		}
	}
	Dump()
}

func diff(args []string) {
	if len(args) > 1 {
		switch args[1] {
		case "--help", "-h", "help":
			fmt.Print("Usage: ridgepole diff [output filename]\n")
			return
		default:
			Diff(args[1])
			return
		}
	}
	fmt.Println(Diff())
}

func run(args []string) {
	if len(args) > 1 {
		switch args[1] {
		case "--help", "-h", "help":
			fmt.Print("Usage: ridgepole run [migrate filename]\n")
			return
		default:
			Run(args[1])
		}
	} else {
		Run()
	}
}

func model(args []string) {
	if len(args) > 1 {
		switch args[1] {
		case "--help", "-h", "help":
			fmt.Print("Usage: ridgepole model [output path]\n")
			return
		default:
			Model(args[1])
		}
	} else {
		Model()
	}
}

func initDatabase(args []string) {
	if len(args) > 1 {
		switch args[1] {
		case "--help", "-h", "help":
			fmt.Print("Usage: ridgepole init [database name]\n")
			return
		default:
			CreateDatabase(args[1])
		}
	} else {
		CreateDatabase()
	}
}

func GetEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}

func IsMySQL() bool {
	return GetEnv("RDB_TYPE", "") == "mysql"
}

func IsPostgres() bool {
	return GetEnv("RDB_TYPE", "") == "postgres"
}

func Dump() {
	if IsMySQL() {
		mysql.Dump()
	}
	if IsPostgres() {
		postgresql.Dump()
	}
}

func Diff(filename ...string) string {
	if IsMySQL() {
		return mysql.Diff(filename...)
	}
	if IsPostgres() {
		return postgresql.Diff(filename...)
	}
	return ""
}

func Run(filename ...string) {
	if IsMySQL() {
		mysql.Run(filename...)
	}
	if IsPostgres() {
		postgresql.Run(filename...)
	}
}

func Model(path ...string) {
	if IsMySQL() {
		mysql.Model(path...)
	}
	if IsPostgres() {
		postgresql.Model(path...)
	}
}

func CreateDatabase(name ...string) {
	if IsMySQL() {
		mysql.CreateDatabase(name...)
	}
	if IsPostgres() {
		postgresql.CreateDatabase(name...)
	}
}
