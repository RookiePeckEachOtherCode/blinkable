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

// UserServiceImpl implements the last service interface defined in the IDL.
type UserServiceImpl struct{}

// UserLogin implements the UserServiceImpl interface.
func (s *UserServiceImpl) UserLogin(ctx context.Context, userLoginRequsts *user.UserLoginRequsts) (resp *user.UserLoginResponse, err error) {
	// TODO: Your code here...
	return
}

// UserRegister implements the UserServiceImpl interface.
func (s *UserServiceImpl) UserRegister(ctx context.Context, req *user.UserRegisterRequest) (resp *user.UseeRegisterResponse, err error) {
	resp = new(user.UseeRegisterResponse)
	query.SetDefault(db.DB)
	ud := query.User

	dbUser, err := ud.WithContext(ctx).Where(ud.Usernmae.Eq(req.Username)).Select().Find()

	if err != nil {
		resp = new(user.UseeRegisterResponse)
		resp.StatusCode = -1
		resp.StatusMsg = "查询用户失败"
		errno.HandleErrWithNormal("查询用户失败", err)
		return
	}

	if len(dbUser) > 0 {
		resp = new(user.UseeRegisterResponse)
		resp.StatusCode = -1
		resp.StatusMsg = "用户名已存在"
		errno.Error("用户名已存在")
		return
	}

	hashPassword, err := hash.HashPassword(req.Password)

	if err != nil {
		resp = new(user.UseeRegisterResponse)
		resp.StatusCode = -1
		resp.StatusMsg = "hash密码错误"
		errno.HandleErrWithNormal("hash密码错误", err)
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
		errno.HandleErrWithNormal(resp.StatusMsg, err)
		return
	}

	token, err := Jwt.NewJwt().ReleaseToken(Jwt.Claims{Id: int64(newUser.ID)})

	if err != nil {
		resp = new(user.UseeRegisterResponse)
		resp.StatusCode = -1
		resp.StatusMsg = "token颁发失败"
		errno.Error(resp.StatusMsg)
		return
	}

	resp.StatusMsg = "success"
	resp.UserId = int32(newUser.ID)
	resp.Token = token

	return
}
