package main

import (
	userHandler "blinkable/cmd/api/user"
	"blinkable/pkg/viper"
	"blinkable/pkg/zap"
	"fmt"

	"github.com/cloudwego/hertz/pkg/app/server"
)

var (
	zlog       = zap.Init()
	cfg        = viper.Load("api")
	serverAddr = fmt.Sprintf("%s:%s", cfg.Viper.GetString("server.host"), cfg.Viper.GetString("server.port"))
)

func InitHertz() *server.Hertz {
	h := server.Default(
		server.WithHostPorts(serverAddr),
	)
	Register(h)
	return h
}
func Register(h *server.Hertz) {
	blinkable := h.Group("/blinkable")
	{
		user := blinkable.Group("/user")
		{
			user.POST("/login", userHandler.Login)
			user.POST("/register", userHandler.Register)

		}
	}
}

func main() {
	zlog.Debug("Starting blinkable api server...")
	h := InitHertz()
	zlog.Debug("Starting blinkable api server...OK")
	h.Spin()
}
