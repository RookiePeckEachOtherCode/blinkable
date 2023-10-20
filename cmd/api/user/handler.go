package userHandler

import (
	"blinkable/cmd/api/user/rpc"
	"blinkable/common/response"
	"blinkable/kitex_gen/user"
	"context"
	"net/http"

	"github.com/cloudwego/hertz/pkg/app"
)

func Login(ctx context.Context, c *app.RequestContext) {
	// username, ok1 := c.GetQuery("username")
	// password, ok2 := c.GetQuery("password")

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

	resp, _ := rpc.Register(ctx, req)

	if resp.StatusCode == -1 {
		c.JSON(http.StatusOK, response.User{
			BaseResponse: response.BaseResponse{
				StatusCode: 0,
				StatusMsg:  resp.StatusMsg,
			},
		})
		return
	}

	c.JSON(http.StatusOK, response.User{
		BaseResponse: response.BaseResponse{
			StatusCode: 0,
			StatusMsg:  resp.StatusMsg,
		},
		UserId: resp.UserId,
		Token:  resp.Token,
	})
}
