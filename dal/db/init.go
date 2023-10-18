package db

import (
	"blinkable/common/errs"
	"blinkable/pkg/viper"
	"fmt"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	DB     *gorm.DB
	cfg    = viper.Load("db")
	isInit = false
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
	db, err := gorm.Open(mysql.Open(getDBConnInfo()), &gorm.Config{
		PrepareStmt:            true, //开启预编译
		SkipDefaultTransaction: true, //禁用默认事务
	})

	errs.HandleErrWithPanic("数据库连接失败", err)

	sqldb, err := db.DB()
	errs.HandleErrWithFatal("", err)

	sqldb.SetMaxOpenConns(1000)                //设置最大连接数
	sqldb.SetMaxIdleConns(20)                  //设置最大空闲连接数
	sqldb.SetConnMaxLifetime(60 * time.Minute) //设置最大连接周期

	isInit = true
}

func GetDB() *gorm.DB {
	return DB
}
func GetIsInit() bool {
	return isInit
}
