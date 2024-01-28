package init

import (
	"blinkable/server/common/consts"
	"blinkable/server/service/homepage/config"
	"blinkable/server/service/homepage/model"
	"fmt"
	"time"

	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/redis/go-redis/v9"
	"gorm.io/driver/mysql"
	"gorm.io/gen"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
	"gorm.io/plugin/opentelemetry/logging/logrus"
	"gorm.io/plugin/opentelemetry/tracing"
)

type Querier interface {
	// SELECT * FROM @@table WHERE name = @name{{if role !=""}} AND role = @role{{end}}
	FilterWithNameAndRole(name, role string) ([]gen.T, error)
}

func InitDB() *gorm.DB {
	c := config.GlobalServerConfig.MysqlInfo
	dsn := fmt.Sprintf(consts.DSN, c.User, c.Password, c.Host, c.Port, c.Name)
	newLogger := logger.New(
		logrus.NewWriter(),
		logger.Config{
			SlowThreshold: time.Second,
			LogLevel:      logger.Silent,
			Colorful:      true,
		},
	)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		},
		Logger: newLogger,
	})
	if err != nil {
		klog.Fatalf("init gorm failed: %s", err)
	}

	if err := db.Use(tracing.NewPlugin()); err != nil {
		klog.Fatalf("use tracing plugin failed: %s", err)
	}

	g := gen.NewGenerator(gen.Config{
		OutPath: "./dao/query/",
		Mode:    gen.WithoutContext | gen.WithQueryInterface, // generate mode
	})
	g.UseDB(db)
	g.ApplyBasic(model.User{}, model.Guestbook{})
	g.ApplyInterface(func(Querier) {}, model.User{})
	g.Execute()

	return db
}
func InitRedis() *redis.Client {
	client := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%s", config.GlobalServerConfig.RedisInfo.Host, config.GlobalServerConfig.RedisInfo.Port),
		Password: config.GlobalServerConfig.RedisInfo.Password,
		DB:       consts.RedisUserDB,
	})
	return client
}
