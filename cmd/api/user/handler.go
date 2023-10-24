package userHandler

import (
	userpc "blinkable/cmd/api/user/rpc"
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
		zap.S().Error(err)
		return
	}
	if len(username) > 40 || len(password) > 40 {
		err := errno.ErrUserNameOrPasswordIsLong
		c.JSON(http.StatusBadRequest, response.BuildBase(-1, err.Error()))
		zap.S().Error(err)
		return
	}

	req := &user.UserLoginRequest{
		Username: username,
		Password: password,
	}

	resp, err := userpc.Login(ctx, req)

	if err != nil {
		c.JSON(http.StatusOK, response.BuildBase(-1, err.Error()))
		return
	}
	if resp == nil {
		c.JSON(http.StatusOK, response.BuildBase(-1, "resp is nil"))
	}

	// if resp == nil {
	// 	c.JSON(http.StatusOK, response.BuildBase(-1, "resp is nil"))
	// 	return
	// } else if err != nil {
	// 	c.JSON(http.StatusOK, response.BuildBase(-1, err.Error()))
	// 	return
	// } else if resp == nil && err == nil {
	// 	c.JSON(http.StatusOK, response.BuildBase(-1, "resp is nil and err is nil"))
	// }

	c.JSON(http.StatusOK, resp)
}

func Register(ctx context.Context, c *app.RequestContext) {
	username := c.Query("username")
	password := c.Query("password")

	if len(username) == 0 || len(password) == 0 {
		err := errno.ErrUserNameOrPasswordIsEmpty
		c.JSON(http.StatusBadRequest, response.BuildBase(-1, err.Error()))
		zap.S().Error(err)
		return
	}
	if len(username) > 40 || len(password) > 40 {
		err := errno.ErrUserNameOrPasswordIsLong
		c.JSON(http.StatusBadRequest, response.BuildBase(-1, err.Error()))
		zap.S().Error(err)
		return
	}

	req := &user.UserRegisterRequest{
		Username: username,
		Password: password,
	}

	resp, err := userpc.Register(ctx, req)

	if err != nil {
		c.JSON(http.StatusOK, response.BuildBase(-1, resp.StatusMsg))
		zap.S().Errorf("%v ===> %v", resp.StatusMsg, err)
		return
	}

	c.JSON(http.StatusOK, resp)
}
