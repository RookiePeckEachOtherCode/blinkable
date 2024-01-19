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
	err = r.redisClient.Set(ctx, "user:"+user.Username, user.ID, 0).Err()
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
func (r *RedisCen) GetUserByUserName(ctx context.Context, name string) (*model.User, error) {
	userId, err := r.redisClient.Get(ctx, "user:"+name).Int64()
	if err != nil && err != redis.Nil {
		klog.Error("redis get user by name failed,", err)
		return nil, err
	}
	user, err := r.GetUserById(ctx, userId)
	if err != nil {
		klog.Error("redis get user by name by id failed,", err)
		return nil, err
	}
	return user, nil
}
