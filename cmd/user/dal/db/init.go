package db

import (
	"database/sql"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	gormopentracing "gorm.io/plugin/opentracing"
	"tiktok-server/internal/conf"
	"tiktok-server/internal/model"
)

var DB *gorm.DB

func Init() {

	db, err := gorm.Open(mysql.Open(getDSN(conf.Config.Database)),
		&gorm.Config{
			PrepareStmt:            true,
			SkipDefaultTransaction: true,
		})
	if err != nil {
		panic(fmt.Errorf("open db:%w", err))
	}

	if err = db.Use(gormopentracing.New()); err != nil {
		panic(fmt.Errorf("register gorm plugin:%w", err))
	}

	sqlDB, err := db.DB()
	if err != nil {
		panic(fmt.Errorf(`get "sql.DB":%w`, err))
	}
	initSQLDB(sqlDB)
	DB = db
}

func initSQLDB(sqlDB *sql.DB) {
	// do nothing
}

func getDSN(config model.Database) string {
	return fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?%s",
		config.Username,
		config.Password,
		config.Host,
		config.Port,
		config.DBName,
		config.Settings)
}
