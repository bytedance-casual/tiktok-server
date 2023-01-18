package db

import (
	"database/sql"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"tiktok-server/internal/conf"
	"tiktok-server/internal/model"
)

var DB *gorm.DB

func Init() error {
	db, err := gorm.Open(mysql.Open(getDSN(conf.Config.Database)), &gorm.Config{
		PrepareStmt: true,
	})
	if err != nil {
		return fmt.Errorf("open db:%w", err)
	}
	sqlDB, err := db.DB()
	if err != nil {
		return fmt.Errorf(`get "sql.DB":%w`, err)
	}
	initSQLDB(sqlDB)
	DB = db
	return nil
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
