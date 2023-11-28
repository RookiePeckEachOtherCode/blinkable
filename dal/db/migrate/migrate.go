package main

import (
	"blinkable/dal/db/model"
	"blinkable/pkg/viper"
	zlog "blinkable/pkg/zap"
	"fmt"

	"go.uber.org/zap"
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
	zlog.Init()

	db, err := gorm.Open(mysql.Open(getDBConnInfo()), &gorm.Config{})
	if err != nil {
		zap.S().Panicf("数据库连接错误 ===>  %v", err)
	}

	//drop table
	err = db.Migrator().DropTable(&model.User{}, &model.Article{}, &model.Comment{}, &model.Guestbook{}, &model.Admin{})
	if err != nil {
		zap.S().Fatalf("drop table 失败 ===>  %v", err)
	}

	err = db.AutoMigrate(&model.User{}, &model.Article{}, &model.Comment{}, &model.Guestbook{}, &model.Admin{})
	if err != nil {
		zap.S().Panicf("automigate 失败 ===>  %v", err)
	}
}
