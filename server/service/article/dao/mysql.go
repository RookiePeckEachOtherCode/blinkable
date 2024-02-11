package dao

import (
	"blinkable/server/service/article/model"
	"context"
	"gorm.io/gorm"
	"time"
)
import "blinkable/server/service/article/dao/query"

type MysqlCen struct {
	q *query.Query
}

func NewUser(db *gorm.DB) *MysqlCen {
	qr := query.Use(db)
	return &MysqlCen{q: qr}
}

func (u MysqlCen) GetArticleSum(ctx context.Context) (sum int64, err error) {
	count, err := u.q.Article.WithContext(ctx).Count()
	if err != nil {
		return 0, err
	}
	return count, nil
}
func (u MysqlCen) GetArticleList(ctx context.Context, start int64, end int64) ([]*model.Article, error) {
	ArticleList, err := u.q.Article.WithContext(ctx).Select(u.q.Article.ID, u.q.Article.CreateTime, u.q.Article.CreaterID).Limit(int(start - end)).Offset(int(start)).Find()
	if err != nil {
		return nil, err
	}
	return ArticleList, nil
}
func (u MysqlCen) GetArticle(ctx context.Context, articleid int32) (*model.Article, error) {
	article, err := u.q.Article.WithContext(ctx).Where(u.q.Article.ID.Eq(int64(articleid))).First()
	if err != nil {
		return nil, err
	}
	return article, nil
}
func (u MysqlCen) AddArticle(ctx context.Context, url string, createrid int64) error {
	article := model.Article{
		CreateTime: time.Now(),
		Context:    url,
		CreaterID:  createrid,
		Comments:   nil,
	}
	err := u.q.Article.WithContext(ctx).Create(&article)
	return err
}
func (u MysqlCen) DeleteArticle(ctx context.Context, articleid int64) error {
	_, err := u.q.Comment.WithContext(ctx).Where(u.q.Comment.ArticleId.Eq(articleid)).Delete()
	if err != nil {
		return err
	}
	_, err = u.q.Article.WithContext(ctx).Where(u.q.Article.ID.Eq(articleid)).Delete()
	if err != nil {
		return err
	}
	return nil
}
func (u MysqlCen) AddComment(ctx context.Context, userid int64, articleid int64, contest string) error {
	comment := model.Comment{
		Context:    contest,
		CreateTime: time.Now(),
		UserID:     userid,
		ArticleId:  articleid,
	}
	err := u.q.Comment.WithContext(ctx).Create(&comment)
	return err
}
