package dao

import (
	"blinkable/server/service/article/model"
	"context"
	"github.com/bytedance/sonic"
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/redis/go-redis/v9"
	"strconv"
)

type RedisCli struct {
	redisClient *redis.Client
}

func NewRedisCli(client *redis.Client) *RedisCli {
	return &RedisCli{redisClient: client}
}

func (r *RedisCli) CreateArticle(ctx context.Context, article *model.Article) error {
	articleJson, err := sonic.Marshal(article)
	if err != nil {
		klog.Errorf("redis marshal article failed: %s", err)
		return err
	}
	err = r.redisClient.Set(ctx, "article:"+strconv.FormatInt(article.ID, 10), articleJson, 0).Err()
	if err != nil {
		klog.Error("redis create article failed %s", err)
		return err
	}
	return nil
}

func (r *RedisCli) GetArticleById(ctx context.Context, ArticleId int64) (*model.Article, error) {
	articleJson, err := r.redisClient.Get(ctx, "article:"+strconv.FormatInt(ArticleId, 10)).Bytes()
	if err != nil && err != redis.Nil {
		klog.Error("redis get article by id failed,", err)
		return nil, err
	}
	var article *model.Article
	err = sonic.Unmarshal(articleJson, &article)
	if err != nil {
		klog.Error("redis unmarshal article failed,", err)
		return nil, err
	}
	return article, nil
}

func (r *RedisCli) DeleteArtcileById(ctx context.Context, ArticleId int64) {
	r.redisClient.Del(ctx, "article:"+strconv.FormatInt(ArticleId, 10))

}
