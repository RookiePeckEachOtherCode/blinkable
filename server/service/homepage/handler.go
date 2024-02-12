package main

import (
	homepage "blinkable/server/kitex_gen/Homepage"
	"blinkable/server/kitex_gen/base"
	"blinkable/server/service/homepage/model"
	"context"
	"net/http"

	"github.com/cloudwego/kitex/pkg/klog"
)

type MysqlCli interface {
	AddGuestbook(ctx context.Context, con string, userid int64, fid int64) error
	AddLikeNumberById(ctx context.Context, id int64) error
	GetUserById(ctx context.Context, id int64) (*model.User, error)
	GetAllAdminUserInfo(ctx context.Context) ([]*model.User, error)
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
	resp = new(homepage.LikeResponse)

	err = s.MysqlCli.AddLikeNumberById(ctx, req.UserId)
	if err != nil {
		klog.Errorf("AddLikeNumberById is failed: %s", err)
		resp = &homepage.LikeResponse{
			StatusCode: 500,
			StatusMsg:  "like failed",
			Succed:     false,
		}
		return resp, err
	}
	user, err := s.MysqlCli.GetUserById(ctx, req.UserId)
	if err != nil {
		klog.Errorf("mysql find user by id failed: %s", err)
		resp = &homepage.LikeResponse{
			StatusCode: 500,
			StatusMsg:  "not find user in database",
			Succed:     false,
		}
		return resp, err
	}
	err = s.RedisCli.UpdateUserInfo(ctx, user)
	if err != nil {
		klog.Errorf("redis update user info failed: %s", err)
		resp = &homepage.LikeResponse{
			StatusCode: 500,
			StatusMsg:  "redis update user info failed",
			Succed:     false,
		}
		return resp, err
	}
	resp = &homepage.LikeResponse{
		StatusCode: http.StatusOK,
		StatusMsg:  "Accepted",
		Succed:     true,
	}
	return resp, nil
}

// GetMainview implements the HomepageServiceImpl interface.
func (s *HomepageServiceImpl) GetMainview(ctx context.Context, req *homepage.GetHomepageRequest) (resp *homepage.GetHomepageResponse, err error) {
	resp = new(homepage.GetHomepageResponse)

	res, err := s.MysqlCli.GetAllAdminUserInfo(ctx)
	if err != nil {
		klog.Errorf("mysql get all admin user is faled: %s", err)
		resp = &homepage.GetHomepageResponse{
			Users:      nil,
			StatusCode: 500,
			StatusMsg:  "find all admin user inf database failed",
			Succed:     false,
		}
		return resp, err
	}
	users := make([]*base.User, len(res))
	for i := 0; i < len(res); i++ {
		users[i] = &base.User{
			Id:            res[i].ID,
			Name:          res[i].Username,
			Avatar:        res[i].Avatar,
			ArticlesNum:   res[i].ArticlesNum,
			Level:         res[i].Level,
			Signature:     res[i].Signature,
			Experience:    res[i].Experience,
			BackgroundImg: res[i].BackgroundImage,
			LikeNum:       res[i].LikeNum,
			GithubUrl:     res[i].GithubUrl,
			Guestbooks:    make([]*base.Guestbook, len(res[i].Guestbooks)),
		}
		for j := 0; j < len(res[i].Guestbooks); j++ {
			users[i].Guestbooks[j] = &base.Guestbook{
				Id:         res[i].Guestbooks[j].ID,
				UserId:     res[i].Guestbooks[j].UserID,
				Context:    res[i].Guestbooks[j].Context,
				FromUserId: res[i].Guestbooks[j].FromUserID,
				CreateTime: res[i].Guestbooks[j].CreateTime.Format("2006-01-02 15:04:05"),
			}
		}
	}

	resp = &homepage.GetHomepageResponse{
		Users:      users,
		StatusCode: 0,
		StatusMsg:  "success",
		Succed:     true,
	}

	resp.StatusMsg = "Accept"
	resp.StatusCode = http.StatusOK
	resp.Succed = true

	return resp, nil
}

// AddGuestbook implements the HomepageServiceImpl interface.
func (s *HomepageServiceImpl) AddGuestbook(ctx context.Context, req *homepage.AddGuestbookRequest) (resp *homepage.AddGuestbookResponse, err error) {
	resp = new(homepage.AddGuestbookResponse)
	err = s.MysqlCli.AddGuestbook(ctx, req.Context, req.UserId, req.FromUserId)
	if err != nil {
		klog.Errorf("create guestbook failed in database")
		resp = &homepage.AddGuestbookResponse{
			StatusCode: http.StatusBadRequest,
			StatusMsg:  "create guestbook failed",
			Succed:     false,
		}
		return resp, err
	}

	user, err := s.MysqlCli.GetUserById(ctx, req.UserId)
	if err != nil {
		klog.Errorf("get user by id failed: %s", err)
		resp = &homepage.AddGuestbookResponse{
			StatusCode: http.StatusBadRequest,
			StatusMsg:  "get user by id failed",
			Succed:     false,
		}
		return resp, err
	}
	err = s.RedisCli.UpdateUserInfo(ctx, user)
	if err != nil {
		klog.Errorf("redis update user info failed: %s", err)
		resp = &homepage.AddGuestbookResponse{
			StatusCode: http.StatusBadRequest,
			StatusMsg:  "redis update user info failed",
			Succed:     false,
		}
		return resp, err
	}

	resp = &homepage.AddGuestbookResponse{
		StatusCode: http.StatusOK,
		StatusMsg:  "Accept",
		Succed:     true,
	}
	return resp, nil
}
