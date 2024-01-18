package main

import (
	"blinkable/server/common/consts"
	"blinkable/server/common/middleware"
	"blinkable/server/common/middleware/models"
	"blinkable/server/kitex_gen/base"
	user "blinkable/server/kitex_gen/user"
	"blinkable/server/service/user/config"
	"blinkable/server/service/user/model"
	"context"
	"time"

	"github.com/bwmarrin/snowflake"
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/golang-jwt/jwt"
	"gorm.io/gorm"
)

type MysqlCen interface {
	CreateUser(ctx context.Context, user *model.User) error
	GetUserByUserName(ctx context.Context, Username string) (*model.User, error)
}
type RedisCen interface {
	CreateUser(ctx context.Context, user *model.User) error
	GetUserById(ctx context.Context, id int64) (*model.User, error)
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

	if user.Password != middleware.Md5Crypt(req.Password, config.GlobalServerConfig.Salt) {
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
func (s *UserServiceImpl) UserRegister(ctx context.Context, req *user.UserRegisterRequest) (resp *user.UserRegisterResponse, err error) {
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
		Password: middleware.Md5Crypt(req.Password, config.GlobalServerConfig.Salt),
		//TODO rand pic
		Avatar:          "",
		BackgroundImage: "",
		Signature:       "这个人很懒，什么都没有",
	}
	err = s.MysqlCen.CreateUser(ctx, newUser)
	if err != nil {
		if err.Error() == consts.ErrUserIsExist {
			resp.BaseResp = &base.BaseResponse{
				StatusCode: 500,
				StatusMsg:  "user already exists",
				Succed:     false,
			}
			return resp, err
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
func (s *UserServiceImpl) GetUserInfo(ctx context.Context, req *user.GetUserInfoRequest) (resp *user.GetUserInfoResponse, err error) {
	resp = new(user.GetUserInfoResponse)
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
	}

	return resp, nil
}
