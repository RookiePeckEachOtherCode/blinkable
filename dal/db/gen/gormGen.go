package main

import (
	"blinkable/dal/db/model"
	"blinkable/pkg/viper"
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gen"
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

// Querier Dynamic SQL
type Querier interface {
	// FilterWithNameAndRole SELECT * FROM @@table WHERE name = @name{{if role !=""}} AND role = @role{{end}}
	FilterWithNameAndRole(name, role string) ([]gen.T, error)
}

func main() {
	g := gen.NewGenerator(gen.Config{
		OutPath: "../query",
		Mode:    gen.WithoutContext | gen.WithDefaultQuery | gen.WithQueryInterface, // generate mode
	})

	gormdb, _ := gorm.Open(mysql.Open(getDBConnInfo()))
	g.UseDB(gormdb) // reuse your gorm db

	// Generate basic type-safe DAO API for struct `model.User` following conventions
	g.ApplyBasic(model.User{}, model.Article{}, model.Comment{}, model.Guestbook{}, model.Admin{})

	// Generate Type Safe API with Dynamic SQL defined on Querier interface
	g.ApplyInterface(func(Querier) {}, model.User{}, model.Article{}, model.Comment{}, model.Guestbook{}, model.Admin{})

	// Generate the code
	g.Execute()
}
