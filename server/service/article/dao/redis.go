package dao

import "github.com/redis/go-redis/v9"

type RedisCli struct {
	redisClient *redis.Client
}

func NewRedisCli(client *redis.Client) *RedisCli {
	return &RedisCli{redisClient: client}
}
