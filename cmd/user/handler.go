package main

import (
	"blinkable/common/consts"
	"blinkable/common/errno"
	"blinkable/dal/db"
	"blinkable/dal/db/model"
	"blinkable/dal/db/query"
	user "blinkable/kitex_gen/user"
	"blinkable/pkg/hash"
	Jwt "blinkable/pkg/jwt"
	"blinkable/pkg/minio"
	"context"
	"time"

	"go.uber.org/zap"
)

func init() {
	query.SetDefault(db.DB)
}

// UserServiceImpl implements the last service interface defined in the IDL.
type UserServiceImpl struct{}

// UserLogin implements the UserServiceImpl interface.
func (s *UserServiceImpl) UserLogin(ctx context.Context, req *user.UserLoginRequest) (resp *user.UserLoginResponse, err error) {
	ud := query.Q.User

	dbUser, _ := ud.WithContext(ctx).Where(ud.Usernmae.Eq(req.Username)).Select(ud.ID, ud.Password).First()

	if dbUser == nil {
		resp = &user.UserLoginResponse{
			StatusCode: -1,
			StatusMsg:  errno.ErrUserNotExist.Error(),
		}
		zap.S().Errorf("%v ===> %v", errno.ErrLogin, resp.StatusMsg)
		return resp, nil
	}

	if !hash.CheckPasswordHash(req.Password, dbUser.Password) {
		resp = &user.UserLoginResponse{
			StatusCode: -1,
			StatusMsg:  errno.ErrUserNameOrPasswordIsWrong.Error(),
		}
		zap.S().Errorf("%v ===> %v", errno.ErrLogin, resp.StatusMsg)
		return resp, nil
	}

	token, err := Jwt.NewJwt().ReleaseToken(Jwt.Claims{Id: int64(dbUser.ID)})

	if err != nil {
		resp = &user.UserLoginResponse{
			StatusCode: -1,
			StatusMsg:  errno.ErrTokenRelease.Error(),
		}
		zap.S().Errorf("%v ===> %v ===> %v", errno.ErrLogin, resp.StatusMsg, err)
		return resp, nil
	}

	resp = &user.UserLoginResponse{
		StatusCode: 0,
		StatusMsg:  "success",
		Token:      token,
		UserId:     dbUser.ID,
	}
	return
}

// UserRegister implements the UserServiceImpl interface.
func (s *UserServiceImpl) UserRegister(ctx context.Context, req *user.UserRegisterRequest) (resp *user.UserRegisterResponse, err error) {
	ud := query.Q.User

	dbUser, err := ud.WithContext(ctx).Where(ud.Usernmae.Eq(req.Username)).Select().Find()

	if err != nil {
		resp = &user.UserRegisterResponse{
			StatusCode: -1,
			StatusMsg:  errno.ErrInternalServerError.Error(),
		}
		zap.S().Errorf("%v ===> %v ===> %v", errno.ErrRegister, resp.StatusMsg, err)
		return resp, nil
	}

	if len(dbUser) > 0 {
		resp = &user.UserRegisterResponse{
			StatusCode: -1,
			StatusMsg:  errno.ErrUserIsExist.Error(),
		}
		zap.S().Errorf("%v ===> %v", errno.ErrRegister, resp.StatusMsg)
		return resp, nil
	}

	hashPassword, err := hash.HashPassword(req.Password)

	if err != nil {
		resp = &user.UserRegisterResponse{
			StatusCode: -1,
			StatusMsg:  errno.ErrHashPasswordIsWrong.Error(),
		}
		zap.S().Errorf("%v ===> %v ===> %v", errno.ErrRegister, resp.StatusMsg, err)
		return resp, nil
	}

	newUser := &model.User{
		Usernmae: req.Username,
		Password: hashPassword,
		Model: model.Model{
			CreateAt: time.Now(),
			UpdateAt: time.Now(),
		},
	}

	err = ud.WithContext(ctx).Create(newUser)

	if err != nil {
		resp = &user.UserRegisterResponse{
			StatusCode: -1,
			StatusMsg:  errno.ErrCreateUserIsWrong.Error(),
		}
		zap.S().Errorf("%v ===> %v ===> %v", errno.ErrRegister, resp.StatusMsg, err)
		return resp, nil
	}

	token, err := Jwt.NewJwt().ReleaseToken(Jwt.Claims{Id: int64(newUser.ID)})

	if err != nil {
		resp = &user.UserRegisterResponse{
			StatusCode: -1,
			StatusMsg:  errno.ErrTokenRelease.Error(),
		}
		zap.S().Errorf("%v ===> %v ===> %v", errno.ErrRegister, resp.StatusMsg, err)
		return resp, nil
	}

	resp = &user.UserRegisterResponse{
		StatusCode: 0,
		StatusMsg:  "success",
		Token:      token,
		UserId:     newUser.ID,
	}
	return resp, nil
}

// UserInfo implements the UserServiceImpl interface.
func (s *UserServiceImpl) UserInfo(ctx context.Context, req *user.UserInfoRequest) (resp *user.UserInfoResponse, err error) {
	claims, err := Jwt.NewJwt().ParseToken(req.Token)

	if err != nil {
		resp = &user.UserInfoResponse{
			StatusCode: -1,
			StatusMsg:  errno.ErrTokenParse.Error(),
		}
		zap.S().Errorf("%v ===> %v ===> %v", errno.ErrTokenParse, resp.StatusMsg, err)
		return resp, nil
	}
	if claims.Id != int64(req.UserId) {
		resp = &user.UserInfoResponse{
			StatusCode: -1,
			StatusMsg:  errno.ErrTokenInvalidId.Error(),
		}
		zap.S().Errorf("%v ===> %v", errno.ErrTokenInvalidId, resp.StatusMsg)
		return resp, nil
	}

	ud := query.Q.User

	userInfo, err := ud.WithContext(ctx).Where(ud.ID.Eq(req.UserId)).Select().First()

	if err != nil {
		resp = &user.UserInfoResponse{
			StatusCode: -1,
			StatusMsg:  errno.ErrInternalServerError.Error(),
		}
		zap.S().Errorf("%v ===> %v ===> %v", errno.ErrUserNotFound, resp.StatusMsg, err)
		return resp, nil
	}

	resp = &user.UserInfoResponse{
		StatusCode:    0,
		StatusMsg:     "success",
		Username:      userInfo.Usernmae,
		UserId:        userInfo.ID,
		Avatar:        userInfo.Avatar,
		BackgroundImg: userInfo.BackgroundImage,
		Signature:     userInfo.Signature,
		Level:         userInfo.Level,
		Experience:    userInfo.Experience,
		ArticlesNum:   userInfo.ArticlesNum,
	}

	return resp, nil
}

// UserInfoUpdate implements the UserServiceImpl interface.
func (s *UserServiceImpl) UserInfoUpdate(ctx context.Context, req *user.UserInfoUpdateRequest) (resp *user.UserInfoUpdateResponse, err error) {
	claims, err := Jwt.NewJwt().ParseToken(req.Token)
	if err != nil {
		resp = &user.UserInfoUpdateResponse{
			StatusCode: -1,
			StatusMsg:  errno.ErrTokenParse.Error(),
		}
		zap.S().Errorf("%v ===> %v", resp.StatusMsg, err)
		return resp, nil
	}

	if claims.Id != int64(req.UserId) {
		resp = &user.UserInfoUpdateResponse{
			StatusCode: -1,
			StatusMsg:  errno.ErrTokenParse.Error(),
		}
		zap.S().Errorf("%v ===> %v", resp.StatusMsg, err)
		return resp, nil
	}

	avatarUrl, err := minio.UpLoadImg(consts.AvatarBucketName, req.Avatar, req.AvatarType)
	if err != nil {
		resp = &user.UserInfoUpdateResponse{
			StatusCode: -1,
			StatusMsg:  errno.ErrUploadImgTypeIsWrong.Error(),
		}
		zap.S().Errorf("%v ===> %v", resp.StatusMsg, err)
		return resp, nil
	}

	backgroundImgUrl, err := minio.UpLoadImg(consts.BackgroundBucketName, req.BackgroundImg, req.BackgroundImgType)
	if err != nil {
		resp = &user.UserInfoUpdateResponse{
			StatusCode: -1,
			StatusMsg:  errno.ErrUploadImgTypeIsWrong.Error(),
		}
		zap.S().Errorf("%v ===> %v", resp.StatusMsg, err)
		return resp, nil
	}

	ud := query.Q.User

	_, err = ud.WithContext(ctx).Where(ud.ID.Eq(req.UserId)).UpdateSimple(ud.Avatar.Value(avatarUrl), ud.BackgroundImage.Value(backgroundImgUrl), ud.Signature.Value(req.Signature))

	if err != nil {
		resp = &user.UserInfoUpdateResponse{
			StatusCode: -1,
			StatusMsg:  errno.ErrInternalServerError.Error(),
		}
		zap.S().Errorf("%v ===> %v", resp.StatusMsg, err)
		return resp, nil
	}

	resp = &user.UserInfoUpdateResponse{
		StatusCode: 0,
		StatusMsg:  "success",
	}

	return resp, nil
}
