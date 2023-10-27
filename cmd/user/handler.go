package main

import (
	"blinkable/common/errno"
	"blinkable/dal/db"
	"blinkable/dal/db/model"
	"blinkable/dal/db/query"
	user "blinkable/kitex_gen/user"
	"blinkable/pkg/hash"
	Jwt "blinkable/pkg/jwt"
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
func (s *UserServiceImpl) UserRegister(ctx context.Context, req *user.UserRegisterRequest) (resp *user.UseeRegisterResponse, err error) {
	resp = new(user.UseeRegisterResponse)
	ud := query.Q.User

	dbUser, err := ud.WithContext(ctx).Where(ud.Usernmae.Eq(req.Username)).Select().Find()

	if err != nil {
		resp = &user.UseeRegisterResponse{
			StatusCode: -1,
			StatusMsg:  errno.ErrInternalServerError.Error(),
		}
		zap.S().Errorf("%v ===> %v ===> %v", errno.ErrRegister, resp.StatusMsg, err)
		return resp, nil
	}

	if len(dbUser) > 0 {
		resp = &user.UseeRegisterResponse{
			StatusCode: -1,
			StatusMsg:  errno.ErrUserIsExist.Error(),
		}
		zap.S().Errorf("%v ===> %v", errno.ErrRegister, resp.StatusMsg)
		return resp, nil
	}

	hashPassword, err := hash.HashPassword(req.Password)

	if err != nil {
		resp = &user.UseeRegisterResponse{
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
		resp = &user.UseeRegisterResponse{
			StatusCode: -1,
			StatusMsg:  errno.ErrCreateUserIsWrong.Error(),
		}
		zap.S().Errorf("%v ===> %v ===> %v", errno.ErrRegister, resp.StatusMsg, err)
		return resp, nil
	}

	token, err := Jwt.NewJwt().ReleaseToken(Jwt.Claims{Id: int64(newUser.ID)})

	if err != nil {
		resp = &user.UseeRegisterResponse{
			StatusCode: -1,
			StatusMsg:  errno.ErrTokenRelease.Error(),
		}
		zap.S().Errorf("%v ===> %v ===> %v", errno.ErrRegister, resp.StatusMsg, err)
		return resp, nil
	}

	resp = &user.UseeRegisterResponse{
		StatusCode: 0,
		StatusMsg:  "success",
		Token:      token,
		UserId:     newUser.ID,
	}
	return resp, nil
}

// UserInfo implements the UserServiceImpl interface.
func (s *UserServiceImpl) UserInfo(ctx context.Context, req *user.UserInfoRequest) (resp *user.UserInfoResponse, err error) {
	ud := query.Q.User

	userInfo, err := ud.WithContext(ctx).Where(ud.ID.Eq(uint32(req.UserId))).Select().First()

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
