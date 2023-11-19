package main

import (
	mainviewhandler "blinkable/cmd/api/mainview/handler"
	userhandler "blinkable/cmd/api/user/handler"
	"blinkable/pkg/viper"
	zlog "blinkable/pkg/zap"

	"fmt"

	"github.com/cloudwego/hertz/pkg/app/server"
	"github.com/hertz-contrib/gzip"
	"go.uber.org/zap"
)

var (
	cfg        = viper.Load("api")
	serverAddr = fmt.Sprintf("%s:%d", cfg.Viper.GetString("server.host"), cfg.Viper.GetInt("server.port"))
	// serverName = cfg.Viper.GetString("server.name")
)

func InitHertz() *server.Hertz {
	h := server.Default(
		server.WithHostPorts(serverAddr),
	)
	h.Use(gzip.Gzip(gzip.BestSpeed))
	Register(h)
	return h
}
func Register(h *server.Hertz) {
	blinkable := h.Group("/blinkable")
	{
		user := blinkable.Group("/user")
		{

			user.POST("/login", userhandler.Login)
			user.POST("/register", userhandler.Register)
			user.GET("/info", userhandler.Info)
			user.POST("/update", userhandler.UserInfoUpdate)

		}
		mainview := blinkable.Group("/Main")
		{
			mainview.GET("", mainviewhandler.Getmainview)
			mainview.POST("/like", mainviewhandler.LikeAction)
			mainview.POST("/addguestbook", mainviewhandler.AddGuestbook)
		}
	}
}

func main() {
	zlog.Init()

	zap.S().Debugf("Starting blinkable api server on %s", serverAddr)
	h := InitHertz()
	zap.S().Debugf("Starting blinkable api server on %s...OK", serverAddr)
	h.Spin()
}
