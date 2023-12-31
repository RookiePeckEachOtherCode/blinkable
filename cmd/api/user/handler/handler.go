package handler

import (
	"blinkable/cmd/api/user/rpc"
	"blinkable/common/errno"
	"blinkable/common/response"
	"blinkable/kitex_gen/user"
	Jwt "blinkable/pkg/jwt"
	"bytes"
	"context"
	"io"
	"net/http"
	"path"
	"strconv"

	"github.com/cloudwego/hertz/pkg/app"
	"go.uber.org/zap"
)

func Login(ctx context.Context, c *app.RequestContext) {
	username := c.Query("username")
	password := c.Query("password")

	if len(username) == 0 || len(password) == 0 {
		err := errno.ErrUserNameOrPasswordIsEmpty
		c.JSON(http.StatusBadRequest, response.BuildBase(-1, err.Error()))
		zap.S().Errorf("%v ===> %v", errno.ErrRequestParamsIsWrong, err.Error())
		return
	}
	if len(username) > 40 || len(password) > 40 {
		err := errno.ErrUserNameOrPasswordIsLong
		c.JSON(http.StatusBadRequest, response.BuildBase(-1, err.Error()))
		zap.S().Errorf("%v ===> %v", errno.ErrRequestParamsIsWrong, err.Error())
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

func Info(ctx context.Context, c *app.RequestContext) {
	token := c.Query("token")
	_userId := c.Query("user_id")

	if len(token) == 0 || len(_userId) == 0 {
		c.JSON(http.StatusBadRequest, response.BaseResponse{
			StatusCode: -1,
			StatusMsg:  errno.ErrRequestParamsIsWrong.Error(),
		})
		zap.S().Error(errno.ErrRequestParamsIsWrong)
		return
	}

	userId, _ := strconv.Atoi(_userId)

	claims, err := Jwt.NewJwt().ParseToken(token)

	if err != nil {
		c.JSON(http.StatusBadRequest, response.BaseResponse{
			StatusCode: -1,
			StatusMsg:  err.Error(),
		})
		zap.S().Errorf("%v ===> %v", errno.ErrTokenParse, err.Error())
		return
	}

	if claims.Id != int64(userId) {
		c.JSON(http.StatusBadRequest, response.BaseResponse{
			StatusCode: -1,
			StatusMsg:  errno.ErrTokenInvalidId.Error(),
		})
		zap.S().Errorf("%v ===> %v", errno.ErrTokenInvalidId, err.Error())
		return
	}

	req := &user.UserInfoRequest{
		Token:  token,
		UserId: int32(userId),
	}

	resp, _ := rpc.Info(ctx, req)

	if resp.StatusCode == -1 {
		c.JSON(http.StatusOK, response.BuildBase(-1, resp.StatusMsg))
		zap.S().Errorf("%v ===> %v", errno.ErrUserInfo, resp.StatusMsg)
		return
	}

	c.JSON(http.StatusOK, resp)
}

func UserInfoUpdate(ctx context.Context, c *app.RequestContext) {
	_userId := c.Query("user_id")
	token := c.Query("token")
	signature := c.Query("signature")

	if len(_userId) == 0 || len(token) == 0 {
		c.JSON(http.StatusBadRequest, response.BaseResponse{
			StatusCode: -1,
			StatusMsg:  errno.ErrRequestParamsIsWrong.Error(),
		})
		zap.S().Error(errno.ErrRequestParamsIsWrong)
		return
	}

	userId, _ := strconv.Atoi(_userId)

	avatarImg, err := c.FormFile("avatar")
	if err != nil {
		zap.S().Errorf("%v ===> %v", errno.ErrUserUpdateInfo, err.Error())
		c.JSON(http.StatusBadRequest, response.BuildBase(-1, err.Error()))
		return
	}
	//获取头像的后缀名  eg. [.jpg]
	avatarImgSuffix := path.Ext(avatarImg.Filename)

	backgroundImg, err := c.FormFile("background_img")
	if err != nil {
		zap.S().Errorf("%v ===> %v", errno.ErrUserUpdateInfo, err.Error())
		c.JSON(http.StatusBadRequest, response.BuildBase(-1, err.Error()))
		return
	}
	//获取背景图片的后缀名  eg. [.jpg]
	backgroundImgSuffix := path.Ext(backgroundImg.Filename)
	zap.S().Debugf("backgroundImgSuffix: %v", backgroundImgSuffix)

	avatarImgData, err := avatarImg.Open()
	if err != nil {
		zap.S().Errorf("%v ===> %v", errno.ErrUserUpdateInfo, err.Error())
		c.JSON(http.StatusBadRequest, response.BuildBase(-1, err.Error()))
		return
	}

	avatarBuf := bytes.NewBuffer(nil)
	defer avatarImgData.Close()
	if _, err := io.Copy(avatarBuf, avatarImgData); err != nil {
		zap.S().Errorf("%v ===> %v", errno.ErrUserUpdateInfo, err.Error())
		c.JSON(http.StatusBadRequest, response.BuildBase(-1, err.Error()))
		return
	}

	backgroundImgData, err := backgroundImg.Open()
	if err != nil {
		zap.S().Errorf("%v ===> %v", errno.ErrUserUpdateInfo, err.Error())
		c.JSON(http.StatusBadRequest, response.BuildBase(-1, err.Error()))
		return
	}

	backgroundBuf := bytes.NewBuffer(nil)
	defer backgroundImgData.Close()
	if _, err := io.Copy(backgroundBuf, backgroundImgData); err != nil {
		zap.S().Errorf("%v ===> %v", errno.ErrUserUpdateInfo, err.Error())
		c.JSON(http.StatusBadRequest, response.BuildBase(-1, err.Error()))
		return
	}

	req := &user.UserInfoUpdateRequest{
		UserId:            int32(userId),
		Token:             token,
		Avatar:            avatarBuf.Bytes(),
		AvatarType:        avatarImgSuffix,
		BackgroundImg:     backgroundBuf.Bytes(),
		BackgroundImgType: backgroundImgSuffix,
		Signature:         signature,
	}

	resp, _ := rpc.UserInfoUpdate(ctx, req)

	if resp.StatusCode == -1 {
		c.JSON(http.StatusOK, response.BuildBase(-1, resp.StatusMsg))
		zap.S().Errorf("%v ===> %v", errno.ErrUserUpdateInfo, resp.StatusMsg)
		return
	}

	c.JSON(http.StatusOK, resp)
}
