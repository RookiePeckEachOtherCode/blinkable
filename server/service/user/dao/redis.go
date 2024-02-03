package dao

import (
	"blinkable/server/service/user/model"
	"context"
	"strconv"

	"github.com/bytedance/sonic"
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/redis/go-redis/v9"
)

type RedisCen struct {
	redisClient *redis.Client
}

func NewRedisCen(client *redis.Client) *RedisCen {
	return &RedisCen{redisClient: client}
}

func (r *RedisCen) CreateUser(ctx context.Context, user *model.User) error {
	userJson, err := sonic.Marshal(user)
	if err != nil {
		klog.Errorf("redis marshal user failed: %s", err)
		return err
	}
	err = r.redisClient.Set(ctx, "user:"+strconv.FormatInt(user.ID, 10), userJson, 0).Err()
	if err != nil {
		klog.Error("redis create user failed %s", err)
		return err
	}
	return nil
}
func (r *RedisCen) GetUserById(ctx context.Context, id int64) (*model.User, error) {
	userJson, err := r.redisClient.Get(ctx, "user:"+strconv.FormatInt(id, 10)).Bytes()
	if err != nil && err != redis.Nil {
		klog.Error("redis get user by id failed,", err)
		return nil, err
	}
	var user *model.User
	err = sonic.Unmarshal(userJson, &user)
	if err != nil {
		klog.Error("redis unmarshal user failed,", err)
		return nil, err
	}
	return user, nil
}
func (r *RedisCen) UpdateUserInfo(ctx context.Context, user *model.User) error {
	return r.CreateUser(ctx, user)
}

func (r *RedisCen) CreateGuestBooksByUserId(ctx context.Context, userId int64, guestbooks []*model.Guestbook) error {
	if guestbooks == nil {
		return nil
	}

	json, err := sonic.Marshal(guestbooks)
	if err != nil {
		klog.Errorf("redis marshal guestbooks failed: %s", err)
		return err
	}
	err = r.redisClient.Set(ctx, "guestbooks:"+strconv.FormatInt(userId, 10), json, 0).Err()
	if err != nil {
		klog.Error("redis create user failed %s", err)
		return err
	}
	return nil
}

func (r *RedisCen) GetGuestBooksByUserId(ctx context.Context, userId int64) ([]*model.Guestbook, error) {
	query := "guestbooks" + strconv.FormatInt(userId, 10)
	if r.redisClient.Exists(ctx, query).Val() == 0 {
		return nil, nil
	}

	json, err := r.redisClient.Get(ctx, query).Bytes()
	if err != nil {
		klog.Errorf("redis get guestbooks is failed: %s", err)
		return nil, err
	}

	var guestbooks []*model.Guestbook
	err = sonic.Unmarshal([]byte(json), &guestbooks)
	if err != nil {
		klog.Errorf("redis unmarshal guestbooks is failed: %s", err)
		return nil, err
	}
	return guestbooks, nil
}
