package main

import (
	"blinkable/common/errno"
	"blinkable/dal/db"
	"blinkable/dal/db/model"
	"blinkable/dal/db/query"
	user "blinkable/kitex_gen/user"
	"blinkable/pkg/hash"
	Jwt "blinkable/pkg/middleware/jwt"
	"context"
	"time"
)

func init() {
	query.SetDefault(db.DB)
}

// UserServiceImpl implements the last service interface defined in the IDL.
type UserServiceImpl struct{}

// UserLogin implements the UserServiceImpl interface.
func (s *UserServiceImpl) UserLogin(ctx context.Context, req *user.UserLoginRequsts) (resp *user.UserLoginResponse, err error) {
	where := "server_user_login"

	ud := query.Q.User

	dbUser, err := ud.WithContext(ctx).Where(ud.Usernmae.Eq(req.Username)).Select(ud.ID, ud.Password).First()

	if err != nil || dbUser == nil {
		resp = new(user.UserLoginResponse)
		resp.StatusCode = -1
		resp.StatusMsg = "用户不存在"
		errno.HandleErrWithError(where, resp.StatusMsg, err)
		return
	}

	if !hash.CheckPasswordHash(req.Password, dbUser.Password) {
		resp = new(user.UserLoginResponse)
		resp.StatusCode = -1
		resp.StatusMsg = "密码错误"
		errno.HandleErrWithError(where, resp.StatusMsg, nil)
	}

	token, err := Jwt.NewJwt().ReleaseToken(Jwt.Claims{Id: int64(dbUser.ID)})

	if err != nil {
		resp = new(user.UserLoginResponse)
		resp.StatusCode = -1
		resp.StatusMsg = "token颁发失败"
		errno.HandleErrWithError(where, resp.StatusMsg, err)
	}

	resp = new(user.UserLoginResponse)
	resp.StatusCode = 0
	resp.StatusMsg = "success"
	resp.Token = token
	resp.UserId = int32(dbUser.ID)
	return
}

// UserRegister implements the UserServiceImpl interface.
func (s *UserServiceImpl) UserRegister(ctx context.Context, req *user.UserRegisterRequest) (resp *user.UseeRegisterResponse, err error) {
	where := "server_user_register"

	ud := query.Q.User

	dbUser, err := ud.WithContext(ctx).Where(ud.Usernmae.Eq(req.Username)).Select().Find()

	if err != nil {
		resp = new(user.UseeRegisterResponse)
		resp.StatusCode = -1
		resp.StatusMsg = "查询用户失败"
		errno.HandleErrWithError(where, resp.StatusMsg, err)
		return
	}

	if len(dbUser) > 0 {
		resp = new(user.UseeRegisterResponse)
		resp.StatusCode = -1
		resp.StatusMsg = "用户名已存在"
		errno.HandleErrWithError(where, resp.StatusMsg, err)
		return
	}

	hashPassword, err := hash.HashPassword(req.Password)

	if err != nil {
		resp = new(user.UseeRegisterResponse)
		resp.StatusCode = -1
		resp.StatusMsg = "hash密码错误"
		errno.HandleErrWithError(where, resp.StatusMsg, nil)
		return
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
		resp = new(user.UseeRegisterResponse)
		resp.StatusCode = -1
		resp.StatusMsg = "创建用户失败"
		errno.HandleErrWithError(where, resp.StatusMsg, err)
		return
	}

	token, err := Jwt.NewJwt().ReleaseToken(Jwt.Claims{Id: int64(newUser.ID)})

	if err != nil {
		resp = new(user.UseeRegisterResponse)
		resp.StatusCode = -1
		resp.StatusMsg = "token颁发失败"
		errno.HandleErrWithError(where, resp.StatusMsg, err)
		return
	}

	resp = new(user.UseeRegisterResponse)

	resp.StatusMsg = "success"
	resp.UserId = int32(newUser.ID)
	resp.Token = token

	return
}
