package dao

import (
	"blinkable/server/service/homepage/model"
	"context"
	"strconv"

	"github.com/bytedance/sonic"
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/redis/go-redis/v9"
)

type RedisCli struct {
	redisClient *redis.Client
}

func NewRedisCli(client *redis.Client) *RedisCli {
	return &RedisCli{redisClient: client}
}

func (r *RedisCli) GetUserById(ctx context.Context, id int64) (*model.User, error) {
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

func (r *RedisCli) CreateUser(ctx context.Context, user *model.User) error {
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

func (r *RedisCli) UpdateUserInfo(ctx context.Context, user *model.User) error {
	return r.CreateUser(ctx, user)
}
