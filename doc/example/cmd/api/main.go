package main

import (
	"example/cmd/api/user"

	"github.com/cloudwego/hertz/pkg/app/server"
)

func main() {
	h := server.Default(
		server.WithHostPorts(":8080"), //设置网关监听的端口
	) //创建网关

	h.POST("/login", user.UserLogin)

	h.Spin() //启动网关
}
