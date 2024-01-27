package main

import (
	homepage "blinkable/server/kitex_gen/Homepage"
	"blinkable/server/service/homepage/model"
	"context"
)

type MysqlCli interface {
}
type RedisCli interface {
	GetUserById(ctx context.Context, id int64) (*model.User, error)
	UpdateUserInfo(ctx context.Context, user *model.User) error
	CreateUser(ctx context.Context, user *model.User) error
}

// HomepageServiceImpl implements the last service interface defined in the IDL.
type HomepageServiceImpl struct {
	MysqlCli
	RedisCli
}

// LikeAction implements the HomepageServiceImpl interface.
func (s *HomepageServiceImpl) LikeAction(ctx context.Context, req *homepage.LikeRequest) (resp *homepage.LikeResponse, err error) {
	// TODO: Your code here...
	return
}

// GetMainview implements the HomepageServiceImpl interface.
func (s *HomepageServiceImpl) GetMainview(ctx context.Context, req *homepage.GetHomepageRequest) (resp *homepage.GetHomepageResponse, err error) {
	// TODO: Your code here...
	return
}

// AddGuestbook implements the HomepageServiceImpl interface.
func (s *HomepageServiceImpl) AddGuestbook(ctx context.Context, req *homepage.AddGuestbookRequest) (resp *homepage.AddGuestbookResponse, err error) {
	// TODO: Your code here...
	return
}
