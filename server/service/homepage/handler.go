package main

import (
	"blinkable/server/common/consts"
	homepage "blinkable/server/kitex_gen/Homepage"
	model "blinkable/server/service/homepage/model"
	"context"
	"github.com/cloudwego/kitex/pkg/klog"
	"net/http"
	"time"
)

type MysqlCli interface {
	GetUsersByIds(ctx context.Context, ids []int64) ([]*model.User, error)
	GetGuestbooksById(ctx context.Context, id int64) ([]*model.Guestbook, error)
	AddGuestbook(ctx context.Context, con string, userid int64, fid int64) error
	AddLikeNumberById(ctx context.Context, id int64) error
	GetUserById(ctx context.Context, id int64) (*model.User, error)
}
type RedisCli interface {
	GetUserById(ctx context.Context, id int64) (*model.User, error)
	UpdateUserInfo(ctx context.Context, user *model.User) error
	CreateUser(ctx context.Context, user *model.User) error
	GetGuestbooksByUserID(ctx context.Context, userID int64) ([]*model.Guestbook, error)
	CreateGuestbook(ctx context.Context, guestbook *model.Guestbook) error
}

// HomepageServiceImpl implements the last service interface defined in the IDL.
type HomepageServiceImpl struct {
	MysqlCli
	RedisCli
}

// LikeAction implements the HomepageServiceImpl interface.
func (s *HomepageServiceImpl) LikeAction(ctx context.Context, req *homepage.LikeRequest) (resp *homepage.LikeResponse, err error) {
	resp = new(homepage.LikeResponse)

	err = s.MysqlCli.AddLikeNumberById(ctx, int64(req.AdminId))
	if err != nil {
		klog.Errorf("AddLikeNumberById painc: %s", err)
		resp = &homepage.LikeResponse{
			StatusCode: 500,
			StatusMsg:  "like failed",
			Succed:     false,
		}
		return resp, err
	}
	user, err := s.MysqlCli.GetUserById(ctx, int64(req.AdminId))
	if err != nil {
		klog.Errorf("mysql find user by id failed: %s", err)
		resp = &homepage.LikeResponse{
			StatusCode: 500,
			StatusMsg:  "not find user in database",
			Succed:     false,
		}
		return resp, err
	}
	err = s.RedisCli.CreateUser(ctx, user)
	if err != nil {
		klog.Errorf("redis create user failed: %s", err)
		resp = &homepage.LikeResponse{
			StatusCode: 500,
			StatusMsg:  "redis create user failed",
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
	for _, id := range consts.AdminIds {
		user, err := s.RedisCli.GetUserById(ctx, int64(id))
		if err != nil {
			klog.Errorf("redis get homepage users panic: %s", err)
			resp = &homepage.GetHomepageResponse{
				Users:      nil,
				StatusCode: 500,
				StatusMsg:  "find user in database panic",
				Succed:     false,
			}
			return resp, err
		}
		if user == nil {
			klog.Errorf("redis cant find homepage users: %s", err)
			resp = &homepage.GetHomepageResponse{
				Users:      nil,
				StatusCode: 500,
				StatusMsg:  "find user in database failed",
				Succed:     false,
			}
			return resp, err
		}
		us := &homepage.Users{
			AdminId:    int32(id),
			Signature:  user.Signature,
			Likes:      int32(user.Like),
			Title:      user.Title,
			Comments:   0,
			Guestbooks: nil,
			IconUrl:    user.Avatar,
			ImageUrl:   user.BackgroundImage,
			GitUrl:     user.Git,
		}
		gb, err := s.RedisCli.GetGuestbooksByUserID(ctx, int64(id))
		if err != nil {
			klog.Errorf("redis get guestbooks panic: %s", err)
			resp = &homepage.GetHomepageResponse{
				Users:      nil,
				StatusCode: 500,
				StatusMsg:  "find user in redis panic",
				Succed:     false,
			}
			return resp, err
		}
		if gb == nil {
			gb, err = s.MysqlCli.GetGuestbooksById(ctx, int64(id))
			if err != nil {
				klog.Errorf("query guestbook in database panic:%s", err)
				resp = &homepage.GetHomepageResponse{
					Users:      nil,
					StatusCode: 500,
					StatusMsg:  "find user in database panic",
					Succed:     false,
				}
				return resp, err
			}
			for _, i := range gb {
				err := s.RedisCli.CreateGuestbook(ctx, i)
				if err != nil {
					klog.Errorf("redis CreateGuestbook panic")
				}
				k := homepage.Guestbook{
					BookId:     int32(i.ID),
					UserId:     int32(i.UserID),
					Context:    i.Context,
					FormuserId: int32(i.FromuserID),
					CreateTime: i.CreateTime.String(),
				}
				us.Guestbooks = append(us.Guestbooks, &k)
			}
		} else {
			for _, i := range gb {
				k := homepage.Guestbook{
					BookId:     int32(i.ID),
					UserId:     int32(i.UserID),
					Context:    i.Context,
					FormuserId: int32(i.FromuserID),
					CreateTime: i.CreateTime.String(),
				}
				us.Guestbooks = append(us.Guestbooks, &k)
			}
		}
		resp.Users = append(resp.Users, us)
	}
	resp.StatusMsg = "Accept"
	resp.StatusCode = http.StatusOK
	resp.Succed = true

	return resp, nil
}

// AddGuestbook implements the HomepageServiceImpl interface.
func (s *HomepageServiceImpl) AddGuestbook(ctx context.Context, req *homepage.AddGuestbookRequest) (resp *homepage.AddGuestbookResponse, err error) {
	err = s.MysqlCli.AddGuestbook(ctx, req.Context, int64(req.UserId), int64(req.AdminId))
	if err != nil {
		klog.Errorf("create guestbook failed in database")
		resp = &homepage.AddGuestbookResponse{
			StatusCode: http.StatusBadRequest,
			StatusMsg:  "create guestbook failed",
			Succed:     false,
		}
		return resp, err
	}
	gbk := model.Guestbook{
		CreateTime: time.Now(),
		Context:    req.Context,
		FromuserID: int64(req.UserId),
		UserID:     int64(req.AdminId),
	}
	err = s.RedisCli.CreateGuestbook(ctx, &gbk)
	if err != nil {
		klog.Errorf("create guestbook failed in redis")
		resp = &homepage.AddGuestbookResponse{
			StatusCode: http.StatusBadRequest,
			StatusMsg:  "create guestbook failed",
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
