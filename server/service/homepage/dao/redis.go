package dao

import (
	"blinkable/server/service/homepage/model"
	"context"
	"github.com/bytedance/sonic"
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/redis/go-redis/v9"
	"strconv"
	"time"
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

func (r *RedisCli) CreateGuestbook(ctx context.Context, guestbook *model.Guestbook) error {
	guestbookJson, err := sonic.Marshal(guestbook)
	if err != nil {
		klog.Errorf("redis marshal guestbook failed: %s", err)
		return err
	}

	key := "guestbook:" + strconv.FormatInt(guestbook.UserID, 10)

	err = r.redisClient.HSet(ctx, key, strconv.FormatInt(guestbook.ID, 10), guestbookJson).Err()
	if err != nil {
		klog.Error("redis create guestbook failed %s", err)
		return err
	}
	err = r.redisClient.Expire(ctx, key, 114514*time.Second).Err()
	if err != nil {
		klog.Error("redis set guestbook exceed time failed %s", err)
		return err
	}
	return nil
}

func (r *RedisCli) GetGuestbooksByUserID(ctx context.Context, userID int64) ([]*model.Guestbook, error) {

	key := "guestbook:" + strconv.FormatInt(userID, 10)

	result, err := r.redisClient.HGetAll(ctx, key).Result()
	if err != nil {
		klog.Errorf("redis get guestbooks by user ID failed: %s", err)
		return nil, err
	}

	var guestbooks []*model.Guestbook
	for _, value := range result {
		var guestbook model.Guestbook
		if err := sonic.Unmarshal([]byte(value), &guestbook); err != nil {
			klog.Errorf("redis unmarshal guestbook failed: %s", err)
			return nil, err
		}
		guestbooks = append(guestbooks, &guestbook)
	}

	return guestbooks, nil
}

func (r *RedisCli) UpdateUserInfo(ctx context.Context, user *model.User) error {
	return r.CreateUser(ctx, user)
}
