package base

import (
	"fmt"
	"strings"
	"time"

	"go.uber.org/zap"
	"gorm.io/driver/mysql"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"moul.io/zapgorm2"
)

var Log *zap.Logger
var DB *gorm.DB
var DBType string
var DBName string

func Init() {
	if DB != nil {
		return
	}

	switch GetEnv("APP_MODE", "develop") {
	case "production":
		Log, _ = zap.NewProduction()
	default:
		config := zap.NewDevelopmentConfig()
		config.Encoding = "console"
		Log, _ = config.Build()
	}
	gormLogger := zapgorm2.New(Log)
	gormLogger.SetAsDefault()

	DBType = GetEnv("RDB_TYPE", "mysql")
	dbUser := GetEnv("RDB_USER", "root")
	dbPassword := GetEnv("RDB_PASSWORD", "q1w2e3r4")
	dbHost := GetEnv("RDB_HOST", "db")
	DBName = GetEnv("RDB_DATABASE", "app")
	config := &gorm.Config{
		Logger: gormLogger,
		NowFunc: func() time.Time {
			return time.Now().Local().Truncate(time.Second)
		},
	}
	var db *gorm.DB
	var err error
	switch DBType {
	case "mysql":
		dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", dbUser, dbPassword, dbHost, DBName)
		db, err = gorm.Open(mysql.Open(dsn), config)
		if err != nil {
			dsn := fmt.Sprintf("%s:%s@tcp(%s)/?charset=utf8mb4&parseTime=True&loc=Local", dbUser, dbPassword, dbHost)
			db, err = gorm.Open(mysql.Open(dsn), config)
		}
	case "postgres":
		sslmode := "require"
		if GetEnv("APP_MODE", "production") == "develop" {
			sslmode = "disable"
		}
		host := strings.Split(dbHost, ":")
		dsn := fmt.Sprintf("user=%s password=%s dbname=%s host=%s port=%s sslmode=%s", dbUser, dbPassword, DBName, host[0], host[1], sslmode)
		db, err = gorm.Open(postgres.Open(dsn), config)
		if err != nil {
			dsn := fmt.Sprintf("user=%s password=%s host=%s port=%s sslmode=%s", dbUser, dbPassword, host[0], host[1], sslmode)
			db, err = gorm.Open(postgres.Open(dsn), config)
		}
	default:
		panic(fmt.Sprintf("未定義のデータベースタイプ: %s", DBType))
	}
	if err != nil {
		if strings.Contains(err.Error(), dbPassword) {
			panic(fmt.Errorf("DB接続エラー: host=%s, db=%s", dbHost, DBName))
		} else {
			panic(err)
		}
	}

	if GetEnvBool("DEBUG", false) {
		DB = db.Debug()
	} else {
		DB = db
	}
}
