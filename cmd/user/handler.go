package main

import (
	"blinkable/dal/db"
	user "blinkable/kitex_gen/user"
	"context"
)

// UserServiceImpl implements the last service interface defined in the IDL.
type UserServiceImpl struct{}

// UserLogin implements the UserServiceImpl interface.
func (s *UserServiceImpl) UserLogin(ctx context.Context, userLoginRequsts *user.UserLoginRequsts) (resp *user.UserLoginResponse, err error) {
	// TODO: Your code here...
	return
}

// UserRegister implements the UserServiceImpl interface.
func (s *UserServiceImpl) UserRegister(ctx context.Context, userRegisterRequest *user.UserRegisterRequest) (resp *user.UseeRegisterResponse, err error) {
	// zlog := zap.Init()
	resp = new(user.UseeRegisterResponse)

	username := userRegisterRequest.Username
	password := userRegisterRequest.Password

	if db.GetIsExistByName(ctx, username) {
		resp.StatusCode = -1
		resp.StatusMsg = "用户名已存在"

		return resp, nil
	}

	return
}
