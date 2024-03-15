package database

import (
	"database/sql"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
)

var db *gorm.DB

func Register(driverName, dataSource string) {
	var err error
	sqlDB, err := sql.Open(driverName, dataSource)
	db, err = gorm.Open(mysql.New(mysql.Config{Conn: sqlDB}), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		log.Panic(err)
	}
}

func GetDB() *gorm.DB {
	return db
}
