package handler

import (
	"blinkable/cmd/api/mainview/rpc"
	"blinkable/common/errno"
	"blinkable/common/response"
	mainview "blinkable/kitex_gen/Mainview"
	"context"
	"github.com/cloudwego/hertz/pkg/app"
	"go.uber.org/zap"
	"net/http"
	"strconv"
)

func Getmainview(ctx context.Context, c *app.RequestContext) {

	req := &mainview.GetMainvewRequest{}
	resp, _ := rpc.Getmainview(ctx, req)
	if resp.StatusCode == -1 {
		c.JSON(http.StatusOK, response.BuildBase(-1, resp.StatusMsg))
		zap.S().Errorf("%v ===> %v", errno.ErrRegister, resp.StatusMsg)
		return
	}
	c.JSON(http.StatusOK, resp)
}

func AddGuestbook(ctx context.Context, c *app.RequestContext) {
	userid := c.Query("user_id")
	user_id, _ := strconv.Atoi(userid)
	adminid := c.Query("admin_id")
	admin_id, _ := strconv.Atoi(adminid)
	context := c.Query("context")
	req := &mainview.AddGuestbookRequest{
		UserId:  int32(user_id),
		AdminId: int32(admin_id),
		Context: context,
	}
	resp, _ := rpc.AddGuestbook(ctx, req)
	if resp.StatusCode == -1 {
		c.JSON(http.StatusOK, response.BuildBase(-1, resp.StatusMsg))
		zap.S().Errorf("%v ===> %v", errno.ErrLogin, resp.StatusMsg)
		return
	}
	c.JSON(http.StatusOK, resp)

}

func LikeAction(ctx context.Context, c *app.RequestContext) {
	userid := c.Query("user_id")
	user_id, _ := strconv.Atoi(userid)
	adminid := c.Query("admin_id")
	admin_id, _ := strconv.Atoi(adminid)
	req := &mainview.LikeRequest{
		AdminId: int32(admin_id),
		UserId:  int32(user_id),
	}
	resp, _ := rpc.LikeAction(ctx, req)
	if resp.StatusCode == -1 {
		c.JSON(http.StatusOK, response.BuildBase(-1, resp.StatusMsg))
		zap.S().Errorf("%v ===> %v", errno.ErrLogin, resp.StatusMsg)
		return
	}
	c.JSON(http.StatusOK, resp)

}
