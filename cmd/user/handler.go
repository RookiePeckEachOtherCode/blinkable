package main

import (
	user "blinkable/cmd/user/kitex_gen/user"
	"context"
)

// UserServiceImpl implements the last service interface defined in the IDL.
type UserServiceImpl struct{}

// UserLogin implements the UserServiceImpl interface.
func (s *UserServiceImpl) UserLogin(ctx context.Context, userLoginRequsts *user.UserLoginRequsts) (resp *user.UserLoginResponse, err error) {

	return
}

// UserRegister implements the UserServiceImpl interface.
func (s *UserServiceImpl) UserRegister(ctx context.Context, userRegisterRequest *user.UserRegisterRequest) (resp *user.UseeRegisterResponse, err error) {
	return
}
