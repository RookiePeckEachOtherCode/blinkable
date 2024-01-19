package init

import (
	"blinkable/server/common/consts"
	"blinkable/server/service/user/config"
	"fmt"
	"time"

	"github.com/redis/go-redis/v9"

	"github.com/cloudwego/kitex/pkg/klog"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
	"gorm.io/plugin/opentelemetry/logging/logrus"
	"gorm.io/plugin/opentelemetry/tracing"
)

func InitDB() *gorm.DB {
	c := config.GlobalServerConfig.MysqlInfo
	dsn := fmt.Sprintf(consts.DSN, c.User, c.Password, c.Host, c.Port, c.Name)
	newLogger := logger.New(
		logrus.NewWriter(),
		logger.Config{
			SlowThreshold: time.Second,
			LogLevel:      logger.Silent,
			// LogLevel:   logger.Info,
			Colorful: true,
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
