package main

import (
	"context"
	user "example/cmd/user/kitex_gen/user"
)

// UserServiceImpl implements the last service interface defined in the IDL.
type UserServiceImpl struct{}

// UserLogin implements the UserServiceImpl interface.
func (s *UserServiceImpl) UserLogin(ctx context.Context, req *user.UserLoginRequests) (resp *user.UserLoginResponse, err error) {
	// TODO: Your code here...
	if req.Username == "114514" && req.Password == "114514" {
		resp = &user.UserLoginResponse{
			Msg: "success",
		}
		return
	}
	resp = &user.UserLoginResponse{
		Msg: "fail",
	}
	return
}
