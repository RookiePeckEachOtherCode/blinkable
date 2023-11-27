package handler

import (
	"blinkable/cmd/api/mainview/rpc"
	"blinkable/common/errno"
	"blinkable/common/response"
	mainview "blinkable/kitex_gen/Mainview"
	"context"
	"github.com/cloudwego/hertz/pkg/app"
	"go.uber.org/zap"
	"io"
	"net/http"
	"os"
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
		zap.S().Errorf("%v ===> %v", errno.ErrLike, resp.StatusMsg)
		return
	}
	c.JSON(http.StatusOK, resp)

}
func ChangeCard(ctx context.Context, c *app.RequestContext) {
	// 获取表单中的数据
	icon, err := c.FormFile("icon")
	if err != nil {

	}
	iconfile, _ := os.Open(icon.Filename)
	icondata, _ := io.ReadAll(iconfile)

	image, err := c.Request.FormFile("image")
	if err != nil {

	}
	imagefile, _ := os.Open(image.Filename)
	imagedata, _ := io.ReadAll(imagefile)
	adminID, _ := strconv.Atoi(c.Query("admin_id"))
	title := c.Query("title")
	signature := c.Query("signature")

	// 创建 ChangeCardRequest
	req := &mainview.ChangeCardRequest{
		Icon:      icondata,
		Image:     imagedata,
		AdminId:   int32(adminID),
		Title:     title,
		Signature: signature,
	}

	// 调用 RPC 方法以更改卡片
	resp, _ := rpc.ChangeCard(ctx, req)

	// 处理响应
	if resp.StatusCode == -1 {
		c.JSON(http.StatusOK, response.BuildBase(-1, resp.StatusMsg))
		zap.S().Errorf("%v ===> %v", errno.ErrChangeCard, resp.StatusMsg)
		return
	}

	c.JSON(http.StatusOK, resp)
}
