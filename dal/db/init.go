package db

import (
	"blinkable/pkg/viper"
	zlog "blinkable/pkg/zap"
	"fmt"
	"time"

	"go.uber.org/zap"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	DB  *gorm.DB
	cfg = viper.Load("db")
)

func getDBConnInfo() string {
	host := cfg.Viper.GetString("mysql.host")
	port := cfg.Viper.GetInt("mysql.port")
	user := cfg.Viper.GetString("mysql.user")
	password := cfg.Viper.GetString("mysql.password")
	dbname := cfg.Viper.GetString("mysql.dbname")
	return fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local", user, password, host, port, dbname)
}

func init() {
	zlog.Init()

	db, err := gorm.Open(mysql.Open(getDBConnInfo()), &gorm.Config{
		PrepareStmt:            true, //开启预编译
		SkipDefaultTransaction: true, //禁用默认事务
	})

	if err != nil {
		zap.S().Panicf("数据库连接错误 ===>  %v", err)
	}

	sqldb, err := db.DB()
	if err != nil {
		zap.S().Fatal("get sqldb error")
	}

	sqldb.SetMaxOpenConns(1000)                //设置最大连接数
	sqldb.SetMaxIdleConns(20)                  //设置最大空闲连接数
	sqldb.SetConnMaxLifetime(60 * time.Minute) //设置最大连接周期

	DB = db
}
