package sources

import (
	"database/sql"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	gormopentracing "gorm.io/plugin/opentracing"
	"log"
	"tiktok-server/internal/model"
)

var MysqlSource = &mysqlSource{}

func init() {
	Context.RegisterSource(MysqlSource)
}


func initSQLDB(sqlDB *sql.DB) {
	// do nothing
}

func getDSN(config *model.Database) string {
	return fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?%s",
		config.Username,
		config.Password,
		config.Host,
		config.Port,
		config.DBName,
		config.Settings)
}


type mysqlSource struct {
	Db *gorm.DB
}

func (source *mysqlSource) ConfigAvailable(config *model.Config) bool {
	if config.Database != nil {
		return true
	}
	return false
}

func (source *mysqlSource) Setup(config *model.Config) {
	source.connect(config)
}

func (source *mysqlSource) connect(config *model.Config) {
	db, err := gorm.Open(mysql.Open(getDSN(config.Database)),
		&gorm.Config{
			PrepareStmt:            true,
			SkipDefaultTransaction: true,
		})
	if err != nil {
		log.Panicf("[redisdata.Conect]:%s", err)
	}
	log.Printf("[mysqldata.Setup]:Setup successfully")
	source.Db = db
	if err = db.Use(gormopentracing.New()); err != nil {
		panic(fmt.Errorf("register gorm plugin:%w", err))
	}
	sqlDB, err := db.DB()
	if err != nil {
		panic(fmt.Errorf(`get "sql.DB":%w`, err))
	}
	initSQLDB(sqlDB)
}

func (source *mysqlSource) Close() {}
