package main

import (
	"blinkable/common/errno"
	"blinkable/dal/db/model"
	"blinkable/pkg/viper"
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func getDBConnInfo() string {
	cfg := viper.Load("db")
	host := cfg.Viper.GetString("mysql.host")
	port := cfg.Viper.GetInt("mysql.port")
	user := cfg.Viper.GetString("mysql.user")
	password := cfg.Viper.GetString("mysql.password")
	dbname := cfg.Viper.GetString("mysql.dbname")
	return fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local", user, password, host, port, dbname)
}

func main() {
	where := "migrate"

	db, err := gorm.Open(mysql.Open(getDBConnInfo()), &gorm.Config{})
	if err != nil {
		errno.HandleErrWithPanic(where, "数据库连接错误", err)
	}

	// drop table
	err = db.Migrator().DropTable(&model.User{}, &model.Article{}, &model.Comment{})
	if err != nil {
		errno.HandleErrWithFatal(where, "drop table 失败", err)
	}

	err = db.AutoMigrate(&model.User{}, &model.Article{}, &model.Comment{})
	if err != nil {
		errno.HandleErrWithPanic(where, "automigate 失败", err)
	}
}
