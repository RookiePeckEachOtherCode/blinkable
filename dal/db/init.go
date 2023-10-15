package db

import (
	"blinkable/errs"
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func init() {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", User, Password, Host, Port, DBName)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	errs.HandleErrWithPanic("数据库连接失败", err)

	sqldb, err := db.DB()

	sqldb.SetMaxIdleConns(10)
	sqldb.SetConnMaxIdleTime(10)
	sqldb.SetMaxOpenConns(100)
	sqldb.SetConnMaxLifetime(60)

	errs.HandleErrWithPanic("数据库连接失败", err)

	DB = db
}
