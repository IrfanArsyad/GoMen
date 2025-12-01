package config

import (
	"fmt"
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/driver/postgres"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var DB *gorm.DB

func ConnectDatabase() {
	var err error
	var dialector gorm.Dialector

	cfg := Get().Database

	switch cfg.Driver {
	case "mysql":
		dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
			cfg.Username,
			cfg.Password,
			cfg.Host,
			cfg.Port,
			cfg.Database,
		)
		dialector = mysql.Open(dsn)

	case "postgres":
		dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Jakarta",
			cfg.Host,
			cfg.Username,
			cfg.Password,
			cfg.Database,
			cfg.Port,
		)
		dialector = postgres.Open(dsn)

	case "sqlite":
		dialector = sqlite.Open(cfg.Database)

	default:
		log.Fatal("Unsupported database driver: " + cfg.Driver)
	}

	logLevel := logger.Silent
	if Get().App.Debug {
		logLevel = logger.Info
	}

	DB, err = gorm.Open(dialector, &gorm.Config{
		Logger: logger.Default.LogMode(logLevel),
	})

	if err != nil {
		log.Fatal("Failed to connect to database: " + err.Error())
	}

	log.Println("Database connected successfully")
}

func GetDB() *gorm.DB {
	return DB
}
