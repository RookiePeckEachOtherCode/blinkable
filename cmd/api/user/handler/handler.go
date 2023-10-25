package handler

import (
	"blinkable/cmd/api/user/rpc"
	"blinkable/common/errno"
	"blinkable/common/response"
	"blinkable/kitex_gen/user"
	"context"
	"net/http"

	"github.com/cloudwego/hertz/pkg/app"
	"go.uber.org/zap"
)

func Login(ctx context.Context, c *app.RequestContext) {
	username := c.Query("username")
	password := c.Query("password")

	if len(username) == 0 || len(password) == 0 {
		err := errno.ErrUserNameOrPasswordIsEmpty
		c.JSON(http.StatusBadRequest, response.BuildBase(-1, err.Error()))
		zap.S().Errorf("%v ===> %v", errno.ErrLogin, err.Error())
		return
	}
	if len(username) > 40 || len(password) > 40 {
		err := errno.ErrUserNameOrPasswordIsLong
		c.JSON(http.StatusBadRequest, response.BuildBase(-1, err.Error()))
		zap.S().Errorf("%v ===> %v", errno.ErrLogin, err.Error())
		return
	}

	req := &user.UserLoginRequest{
		Username: username,
		Password: password,
	}

	resp, _ := rpc.Login(ctx, req)

	if resp.StatusCode == -1 {
		c.JSON(http.StatusOK, response.BuildBase(-1, resp.StatusMsg))
		zap.S().Errorf("%v ===> %v", errno.ErrLogin, resp.StatusMsg)
		return
	}

	c.JSON(http.StatusOK, resp)
}

func Register(ctx context.Context, c *app.RequestContext) {
	username := c.Query("username")
	password := c.Query("password")

	if len(username) == 0 || len(password) == 0 {
		err := errno.ErrUserNameOrPasswordIsEmpty
		c.JSON(http.StatusBadRequest, response.BuildBase(-1, err.Error()))
		zap.S().Errorf("%v ===> %v", errno.ErrRegister, err.Error())
		return
	}
	if len(username) > 40 || len(password) > 40 {
		err := errno.ErrUserNameOrPasswordIsLong
		c.JSON(http.StatusBadRequest, response.BuildBase(-1, err.Error()))
		zap.S().Errorf("%v ===> %v", errno.ErrRegister, err.Error())
		return
	}

	req := &user.UserRegisterRequest{
		Username: username,
		Password: password,
	}

	resp, _ := rpc.Register(ctx, req)

	if resp.StatusCode == -1 {
		c.JSON(http.StatusOK, response.BuildBase(-1, resp.StatusMsg))
		zap.S().Errorf("%v ===> %v", errno.ErrRegister, resp.StatusMsg)
		return
	}

	c.JSON(http.StatusOK, resp)
}
