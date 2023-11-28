package main

import (
	mainviewhandler "blinkable/cmd/api/mainview/handler"
	userhandler "blinkable/cmd/api/user/handler"
	"blinkable/pkg/viper"
	zlog "blinkable/pkg/zap"
	"time"

	"fmt"

	"github.com/cloudwego/hertz/pkg/app/server"
	"github.com/hertz-contrib/cors"
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
	h.Use(cors.New(cors.Config{
		//准许跨域请求网站,多个使用,分开,限制使用*
		AllowOrigins: []string{"*"},
		//准许使用的请求方式
		AllowMethods: []string{"PUT", "PATCH", "POST", "GET", "DELETE"},
		//准许使用的请求表头
		AllowHeaders: []string{"Origin", "Authorization", "Content-Type"},
		//显示的请求表头
		ExposeHeaders: []string{"Content-Type"},
		//凭证共享,确定共享
		AllowCredentials: true,
		//容许跨域的原点网站,可以直接return true就万事大吉了
		AllowOriginFunc: func(origin string) bool {
			return true
		},
		//超时时间设定
		MaxAge: 24 * time.Hour,
	}))
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
			mainview.POST("/changecard", mainviewhandler.ChangeCard)
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
