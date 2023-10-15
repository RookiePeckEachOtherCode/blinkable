package main

import (
	userHandler "blinkable/cmd/api/user"

	"github.com/cloudwego/hertz/pkg/app/server"
)

func main() {
	h := server.Default()

	blinkable := h.Group("/blinkable")

	// /blinkable/user
	user := blinkable.Group("/user")
	// /blinkable/user/login
	user.POST("/login", userHandler.Login)
	// /blinkable/user/register
	user.POST("/register", userHandler.Register)

	h.Spin()
}
