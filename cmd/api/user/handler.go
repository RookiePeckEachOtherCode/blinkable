package userHandler

import (
	userpc "blinkable/cmd/api/user/rpc"
	"blinkable/common/response"
	"blinkable/kitex_gen/user"
	"context"
	"fmt"
	"net/http"

	"github.com/cloudwego/hertz/pkg/app"
)

func Login(ctx context.Context, c *app.RequestContext) {
}

func Register(ctx context.Context, c *app.RequestContext) {
	username := c.Query("username")
	password := c.Query("password")

	if len(username) == 0 || len(password) == 0 {
		c.JSON(http.StatusBadRequest, response.User{
			BaseResponse: response.BaseResponse{
				StatusCode: -1,
				StatusMsg:  "用户名或密码不能为空",
			},
		})
		return
	}
	if len(username) > 40 || len(password) > 40 {
		c.JSON(http.StatusBadRequest, response.User{
			BaseResponse: response.BaseResponse{
				StatusCode: -1,
				StatusMsg:  "用户名或者密码太长",
			},
		})
		return
	}

	req := &user.UserRegisterRequest{
		Username: username,
		Password: password,
	}

	resp, err := userpc.Register(ctx, req)

	if err != nil {
		c.JSON(http.StatusOK, response.BaseResponse{
			StatusCode: -1,
			StatusMsg:  fmt.Sprintf("register错误: %v", err),
		})
	}

	c.JSON(http.StatusOK, response.User{
		UserId: resp.UserId,
		Token:  resp.Token,
		BaseResponse: response.BaseResponse{
			StatusCode: resp.StatusCode,
			StatusMsg:  resp.StatusMsg,
		},
	})
}
