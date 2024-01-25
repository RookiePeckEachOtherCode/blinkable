package main

import (
	"blinkable/server/common/consts"
	"blinkable/server/common/middleware"
	"blinkable/server/common/middleware/models"
	"blinkable/server/kitex_gen/base"
	"blinkable/server/service/user/config"
	"blinkable/server/service/user/model"
	"context"
	"errors"
	"fmt"
	"time"

	user "blinkable/server/kitex_gen/user"

	"github.com/bwmarrin/snowflake"
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/golang-jwt/jwt"
	"gorm.io/gorm"
)

type MysqlCen interface {
	CreateUser(ctx context.Context, user *model.User) error
	GetUserByUserName(ctx context.Context, Username string) (*model.User, error)
	GetUserById(ctx context.Context, id int64) (*model.User, error)
	UpdateUserInfo(ctx context.Context, user *model.User) error
}
type RedisCen interface {
	CreateUser(ctx context.Context, user *model.User) error
	GetUserById(ctx context.Context, id int64) (*model.User, error)
	GetUserByUserName(ctx context.Context, name string) (*model.User, error)
	UpdateUserInfo(ctx context.Context, user *model.User) error
}

// UserServiceImpl implements the last service interface defined in the IDL.
type UserServiceImpl struct {
	MysqlCen
	RedisCen
	JWT *middleware.JWT
}

// UserLogin implements the UserServiceImpl interface.
func (s *UserServiceImpl) UserLogin(ctx context.Context, req *user.UserLoginRequest) (resp *user.UserLoginResponse, err error) {
	resp = new(user.UserLoginResponse)

	user, err := s.MysqlCen.GetUserByUserName(ctx, req.Username)
	if user == nil {
		resp.BaseResp = &base.BaseResponse{
			StatusCode: 500,
			StatusMsg:  "user not found",
			Succed:     false,
		}
		return resp, nil
	}
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			resp.BaseResp = &base.BaseResponse{
				StatusCode: 500,
				StatusMsg:  "user not found",
				Succed:     false,
			}
			return resp, err
		} else {
			klog.Errorf("mysql find user by name failed: %s", err)
			resp.BaseResp = &base.BaseResponse{
				StatusCode: 500,
				StatusMsg:  "mysql find user by name failed",
				Succed:     false,
			}
			return resp, err
		}
	}

	if user.Password != middleware.Md5Crypt(
		req.Password,
		config.GlobalServerConfig.Salt,
	) {
		resp.BaseResp = &base.BaseResponse{
			StatusCode: 500,
			StatusMsg:  "WA password",
			Succed:     true,
		}
		return resp, nil
	}

	resp.UserId = user.ID
	resp.Token, err = s.JWT.CreateToken(models.CustomClaims{
		ID: user.ID,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Unix() + 60*60*24*30,
			Issuer:    "Blinkable",
			NotBefore: time.Now().Unix(),
		},
	})

	if err != nil {
		klog.Errorf("jwt token create failed: %s", err)
		resp.BaseResp = &base.BaseResponse{
			StatusCode: 500,
			StatusMsg:  "jwt token create failed",
			Succed:     false,
		}
		return resp, err
	}

	resp.BaseResp = &base.BaseResponse{
		StatusCode: 0,
		StatusMsg:  "AC login",
		Succed:     true,
	}
	return resp, nil
}

// UserRegister implements the UserServiceImpl interface.
func (s *UserServiceImpl) UserRegister(
	ctx context.Context,
	req *user.UserRegisterRequest,
) (resp *user.UserRegisterResponse, err error) {
	resp = new(user.UserRegisterResponse)
	sf, err := snowflake.NewNode(consts.UserSnowflakeNode)
	if err != nil {
		klog.Errorf("generate user snowflake id failed: %s", err)
		resp.BaseResp = &base.BaseResponse{
			StatusCode: 500,
			StatusMsg:  err.Error(),
			Succed:     false,
		}
		return resp, nil
	}
	newUser := &model.User{
		ID:       sf.Generate().Int64(),
		Username: req.Username,
		Password: middleware.Md5Crypt(
			req.Password,
			config.GlobalServerConfig.Salt,
		),
		// TODO rand pic
		Avatar:          "",
		BackgroundImage: "",
		Signature:       "这个人很懒，什么都没有",
		CreateAt:        time.Now(),
		UpdateAt:        time.Now(),
		Title:           "",
	}
	err = s.MysqlCen.CreateUser(ctx, newUser)
	if err != nil {
		if err.Error() == consts.ErrUserIsExist {
			resp.BaseResp = &base.BaseResponse{
				StatusCode: 500,
				StatusMsg:  "user already exists",
				Succed:     true,
			}
			return resp, nil
		} else {
			klog.Errorf("mysql create usre failed: %s", err)
			resp.BaseResp = &base.BaseResponse{
				StatusCode: 500,
				StatusMsg:  "mysql create user failed",
				Succed:     false,
			}
			return resp, err
		}
	}
	err = s.RedisCen.CreateUser(ctx, newUser)
	if err != nil {
		klog.Errorf("redis create user failed: %s", err)
		resp.BaseResp = &base.BaseResponse{
			StatusCode: 500,
			StatusMsg:  "redis create user failed",
			Succed:     false,
		}
	}
	resp.UserId = newUser.ID
	resp.Token, err = s.JWT.CreateToken(models.CustomClaims{
		ID: newUser.ID,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Unix() + 60*60*26*30,
			Issuer:    "Blinkable",
			NotBefore: time.Now().Unix(),
		},
	})
	if err != nil {
		klog.Errorf("jwt token create failed: %s", err)
		resp.BaseResp = &base.BaseResponse{
			StatusCode: 500,
			StatusMsg:  "jwt token create failed",
			Succed:     false,
		}
		return resp, nil
	}
	resp.BaseResp = &base.BaseResponse{
		StatusCode: 0,
		StatusMsg:  "user register success",
		Succed:     true,
	}
	return resp, nil
}

// GetUserInfo implements the UserServiceImpl interface.
func (s *UserServiceImpl) GetUserInfo(
	ctx context.Context,
	req *user.GetUserInfoRequest,
) (resp *user.GetUserInfoResponse, err error) {
	resp = new(user.GetUserInfoResponse)

	if req.Tp == 0 {
		user, err := s.RedisCen.GetUserById(ctx, req.UserId)
		if err != nil {
			klog.Errorf("redis get user by id failed,", err)
			resp.BaseResp = &base.BaseResponse{
				StatusCode: 500,
				StatusMsg:  "redis get user by id failed",
				Succed:     false,
			}
			return nil, err
		}

		resp.BaseResp = &base.BaseResponse{
			StatusCode: 0,
			StatusMsg:  "get user info by id success",
			Succed:     true,
		}
		resp.UserInfo = &base.User{
			Id:            user.ID,
			Name:          user.Username,
			Avatar:        user.Avatar,
			ArticlesNum:   user.ArticlesNum,
			Experience:    user.Experience,
			BackgroundImg: user.BackgroundImage,
			Level:         user.Level,
			Signature:     user.Signature,
			Title:         user.Title,
		}
	} else if req.Tp == 1 {
		user, err := s.RedisCen.GetUserByUserName(ctx, req.UserName)
		if err != nil {
			klog.Errorf("redis get user by name failed,", err)
			resp.BaseResp = &base.BaseResponse{
				StatusCode: 500,
				StatusMsg:  "redis get user by name failed",
				Succed:     false,
			}
			return nil, err
		}
		resp.BaseResp = &base.BaseResponse{
			StatusCode: 0,
			StatusMsg:  "get user info by name success",
			Succed:     true,
		}
		resp.UserInfo = &base.User{
			Id:            user.ID,
			Name:          user.Username,
			Avatar:        user.Avatar,
			ArticlesNum:   user.ArticlesNum,
			Experience:    user.Experience,
			BackgroundImg: user.BackgroundImage,
			Level:         user.Level,
			Signature:     user.Signature,
			Title:         user.Title,
		}
	} else {
		err = errors.New("undefined type")
		klog.Errorf("get user info failed: %s", err)
		resp.BaseResp = &base.BaseResponse{
			StatusCode: 500,
			StatusMsg:  fmt.Sprintf("%s: %s", "get user inf failed", "undefined type"),
			Succed:     false,
		}
		return nil, err
	}

	return resp, nil
}

// UpdateUserInfo implements the UserServiceImpl interface.
func (s *UserServiceImpl) UpdateUserInfo(ctx context.Context, req *user.UpdateUserInfoRequest) (resp *user.UpdateUserInfoResponse, err error) {
	resp = new(user.UpdateUserInfoResponse)
	dataUser, err := s.MysqlCen.GetUserById(ctx, req.UserId)
	if err != nil {
		klog.Errorf("get user info from mysql failed: %s", err)
		resp.BaseResp = &base.BaseResponse{
			StatusCode: 500,
			StatusMsg:  "get user info from mysql failed",
			Succed:     false,
		}
		return nil, err
	}

	if dataUser.Signature != req.Signature {
		dataUser.Signature = req.Signature
	}
	if dataUser.Username != req.Username {
		dataUser.Username = req.Username
	}
	if dataUser.Title != req.Title {
		dataUser.Title = req.Title
	}

	if err := s.MysqlCen.UpdateUserInfo(ctx, dataUser); err != nil {
		klog.Errorf("update userinfo to mysql failed: %s", err)
		resp.BaseResp = &base.BaseResponse{
			StatusCode: 500,
			StatusMsg:  "update userinfo to mysql failed",
			Succed:     false,
		}
		return resp, err
	}
	if err := s.RedisCen.UpdateUserInfo(ctx, dataUser); err != nil {
		klog.Errorf("update userinfo to redis failed: %s", err)
		resp.BaseResp = &base.BaseResponse{
			StatusCode: 500,
			StatusMsg:  "update userinfo to redis failed",
			Succed:     false,
		}
		return resp, err
	}

	resp.BaseResp = &base.BaseResponse{
		StatusCode: 0,
		StatusMsg:  "update userinfo to mysql success",
		Succed:     true,
	}

	return resp, nil
}

// UpdateUserPassword implements the UserServiceImpl interface.
func (s *UserServiceImpl) UpdateUserPassword(ctx context.Context, req *user.UpdateUserPasswordRequest) (resp *user.UpdateUserPasswordResponse, err error) {
	req.Oldpassword = middleware.Md5Crypt(req.Oldpassword, config.GlobalServerConfig.Salt)
	req.Newpassword_ = middleware.Md5Crypt(req.Newpassword_, config.GlobalServerConfig.Salt)
	resp = new(user.UpdateUserPasswordResponse)
	dataUser, err := s.RedisCen.GetUserById(ctx, req.UserId)
	if err != nil {
		klog.Errorf("get user info from redis is failed: %s", err)
		resp.BaseResp = &base.BaseResponse{
			StatusCode: 500,
			StatusMsg:  "get user info from redis is failed",
			Succed:     false,
		}
		return resp, err
	}
	if dataUser.Password != req.Oldpassword {
		klog.Errorf("user old_password incorrect")
		resp.BaseResp = &base.BaseResponse{
			StatusCode: 0,
			StatusMsg:  "user old_password incorrect",
			Succed:     false,
		}
		return resp, nil
	}
	dataUser.Password = req.Newpassword_
	if err := s.MysqlCen.UpdateUserInfo(ctx, dataUser); err != nil {
		klog.Errorf("update user password to mysql is failed: %s", err)
		resp.BaseResp = &base.BaseResponse{
			StatusCode: 500,
			StatusMsg:  "update user password to mysql is failed",
			Succed:     false,
		}
		return resp, err
	}
	if err := s.RedisCen.UpdateUserInfo(ctx, dataUser); err != nil {
		klog.Errorf("update user password to redis is failed: %s", err)
		resp.BaseResp = &base.BaseResponse{
			StatusCode: 500,
			StatusMsg:  "update user password to redis is failed",
			Succed:     false,
		}
		return resp, err
	}

	resp.BaseResp = &base.BaseResponse{
		StatusMsg:  "success",
		StatusCode: 0,
		Succed:     true,
	}
	return resp, nil
}
