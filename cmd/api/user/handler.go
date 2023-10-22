package userHandler

import (
	userpc "blinkable/cmd/api/user/rpc"
	"blinkable/common/errno"
	"blinkable/common/response"
	"blinkable/kitex_gen/user"
	"context"
	"net/http"

	"github.com/cloudwego/hertz/pkg/app"
)

func Login(ctx context.Context, c *app.RequestContext) {
	where := "api_user_login"

	username := c.Query("username")
	password := c.Query("password")

	if len(username) == 0 || len(password) == 0 {
		msg := "用户名或密码不能为空"
		c.JSON(http.StatusBadRequest, response.BuildBase(-1, msg))
		errno.HandleErrWithError(where, msg, nil)
		return
	}
	if len(username) > 40 || len(password) > 40 {
		msg := "用户名或密码过长"
		c.JSON(http.StatusBadRequest, response.BuildBase(-1, msg))
		errno.HandleErrWithError(where, msg, nil)
		return
	}

	req := &user.UserLoginRequsts{
		Username: username,
		Password: password,
	}

	resp, err := userpc.Login(ctx, req)

	if err != nil {
		c.JSON(http.StatusOK, response.BuildBase(-1, resp.StatusMsg))
		errno.HandleErrWithError(where, resp.StatusMsg, err)
		return
	}

	c.JSON(http.StatusOK, resp)
}

func Register(ctx context.Context, c *app.RequestContext) {
	where := "api_user_register"
	username := c.Query("username")
	password := c.Query("password")

	if len(username) == 0 || len(password) == 0 {
		msg := "用户名或密码不能为空"
		c.JSON(http.StatusBadRequest, response.BuildBase(-1, msg))
		errno.HandleErrWithError(where, msg, nil)
		return
	}
	if len(username) > 40 || len(password) > 40 {
		msg := "用户名或密码过长"
		c.JSON(http.StatusBadRequest, response.BuildBase(-1, msg))
		errno.HandleErrWithError(where, msg, nil)
		return
	}

	req := &user.UserRegisterRequest{
		Username: username,
		Password: password,
	}

	resp, err := userpc.Register(ctx, req)

	if err != nil {
		c.JSON(http.StatusOK, response.BuildBase(-1, resp.StatusMsg))
		errno.HandleErrWithError(where, resp.StatusMsg, err)
		return
	}

	c.JSON(http.StatusOK, resp)
}
