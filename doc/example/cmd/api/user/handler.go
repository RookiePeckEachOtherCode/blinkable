package user

import (
	"context"
	"example/cmd/user/kitex_gen/user"
	"example/cmd/user/kitex_gen/user/userservice"
	"net/http"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/kitex/client"
	etcd "github.com/kitex-contrib/registry-etcd"
)

var userClient userservice.Client //用户服务客户端

func init() {
	r, err := etcd.NewEtcdResolver([]string{"127.0.0.1:2379"})
	if err != nil {
		panic(err)
	}
	_userClient, err := userservice.NewClient(
		"user-service",
		client.WithResolver(r),
	)
	if err != nil {
		panic(err)
	}
	userClient = _userClient
}
func UserLogin(ctx context.Context, c *app.RequestContext) {
	username, _ := c.GetQuery("username") //获取请求参数
	password, _ := c.GetQuery("password") //获取请求参数

	resp, err := userClient.UserLogin(ctx, &user.UserLoginRequests{ //通过客户端调用服务端
		Username: username,
		Password: password,
	})

	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
	}
	c.JSON(http.StatusOK, resp)
}
