package main

import (
	article "blinkable/server/kitex_gen/Article"
	"blinkable/server/kitex_gen/base"
	"blinkable/server/service/article/model"
	"context"
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/redis/go-redis/v9"
	"net/http"
	"time"
)

type MysqlCli interface {
	GetArticleSum(ctx context.Context) (sum int64, err error)
	GetArticleList(ctx context.Context, start int64, end int64) ([]*model.Article, error)
	GetArticle(ctx context.Context, articleid int32) (*model.Article, error)
	AddArticle(ctx context.Context, content string, createrid int64, title string) error
	DeleteArticle(ctx context.Context, articleid int64) error
	AddComment(ctx context.Context, userid int64, articleid int64, contest string) error
}
type RedisCli interface {
	CreateArticle(ctx context.Context, article *model.Article) error
	GetArticleById(ctx context.Context, ArticleId int64) (*model.Article, error)
	DeleteArtcileById(ctx context.Context, ArticleId int64)
}

// ArticleServiceImpl implements the last service interface defined in the IDL.
type ArticleServiceImpl struct {
	MysqlCli
	RedisCli
}

// GetArticleSum implements the ArticleServiceImpl interface.
func (s *ArticleServiceImpl) GetArticleSum(ctx context.Context, req *article.GetArticleSumRequest) (resp *article.GetArticleSumResponse, err error) {
	resp = new(article.GetArticleSumResponse)
	sum, err := s.MysqlCli.GetArticleSum(ctx)
	if err != nil {
		klog.Errorf("Get article sum failed:%s", err)
		resp = &article.GetArticleSumResponse{
			StatusCode: http.StatusBadRequest,
			StatusMsg:  "ERROR",
			Succed:     false,
			Sum:        0,
		}
		return resp, err
	}
	resp = &article.GetArticleSumResponse{
		StatusCode: http.StatusOK,
		StatusMsg:  "Accept",
		Succed:     false,
		Sum:        sum,
	}
	return resp, nil
}

// GetArticleList implements the ArticleServiceImpl interface.
func (s *ArticleServiceImpl) GetArticleList(ctx context.Context, req *article.GetArticleListRequest) (resp *article.GetArticleListResponse, err error) {
	resp = &article.GetArticleListResponse{
		StatusCode: 0,
		StatusMsg:  "",
		Succed:     false,
		Articles:   nil,
	}
	resp.Articles = make([]*base.ArticleMsg, 0)
	sum, _ := s.MysqlCli.GetArticleSum(ctx)
	articleList, err := s.MysqlCli.GetArticleList(ctx, int64(req.Start), int64(min(req.End, int32(sum))))
	if err != nil {
		klog.Errorf("Get articlelist failed:%s", err)
		resp = &article.GetArticleListResponse{
			StatusCode: http.StatusBadRequest,
			StatusMsg:  "get articlelist failed",
			Succed:     false,
			Articles:   nil,
		}
	}
	for i := 0; i < len(articleList); i++ {
		resp.Articles = append(resp.Articles, &base.ArticleMsg{
			CreateTime: articleList[i].CreateTime.String(),
			UpdateTime: "",
			CreaterId:  articleList[i].CreaterID,
			ArticleId:  articleList[i].ID,
			Title:      articleList[i].Title,
		})
	}
	resp.StatusMsg = "Accept"
	resp.Succed = true
	resp.StatusCode = 200
	return resp, nil
}

// GetArticle implements the ArticleServiceImpl interface.
func (s *ArticleServiceImpl) GetArticle(ctx context.Context, req *article.GetArticleRequest) (resp *article.GetArticleResponse, err error) {
	resp = &article.GetArticleResponse{
		StatusCode: 0,
		StatusMsg:  "",
		Succed:     false,
		Comments:   nil,
		Content:    "",
		CreaterId:  0,
	}
	resp.Comments = make([]*base.Comment, 0)
	result, err := s.RedisCli.GetArticleById(ctx, int64(req.ArticleId))
	if err != nil && err != redis.Nil {
		klog.Errorf("get article failed:%s", err)
		resp = &article.GetArticleResponse{
			StatusCode: http.StatusBadRequest,
			StatusMsg:  "failed get article in redis",
			Succed:     false,
			Comments:   nil,
			Content:    "",
			CreaterId:  0,
		}
		return resp, err
	}
	if err == redis.Nil {
		resl, err := s.MysqlCli.GetArticle(ctx, req.ArticleId)
		if err != nil {
			klog.Errorf("get article failed:%s", err)
			resp = &article.GetArticleResponse{
				StatusCode: http.StatusBadRequest,
				StatusMsg:  "failed get article in database",
				Succed:     false,
				Comments:   nil,
				Content:    "",
				CreaterId:  0,
			}
			return resp, err
		}
		err = s.RedisCli.CreateArticle(ctx, resl)
		if err != nil {
			klog.Errorf("create article in redis failed:%s", err)
			resp = &article.GetArticleResponse{
				StatusCode: http.StatusBadRequest,
				StatusMsg:  "create article failed",
				Succed:     false,
				Comments:   nil,
				Content:    "",
				CreaterId:  0,
			}
			return resp, err
		}
		for _, i := range resl.Comments {
			resp.Comments = append(resp.Comments, &base.Comment{
				CommentId:  int32(i.ID),
				UserId:     i.UserID,
				Context:    i.Context,
				CreateTime: i.CreateTime.String(),
			})
		}
		resp.StatusMsg = "Accept"
		resp.StatusCode = 200
		resp.CreaterId = resl.CreaterID
		resp.Content = resl.Context
		resp.Succed = true
	} else {
		for _, i := range result.Comments {
			resp.Comments = append(resp.Comments, &base.Comment{
				CommentId:  int32(i.ID),
				UserId:     i.UserID,
				Context:    i.Context,
				CreateTime: i.CreateTime.String(),
			})
		}
		resp.StatusMsg = "Accept"
		resp.StatusCode = 200
		resp.CreaterId = result.CreaterID
		resp.Content = result.Context
		resp.Succed = true
	}
	return resp, nil
}

// PublishArticle implements the ArticleServiceImpl interface.
func (s *ArticleServiceImpl) PublishArticle(ctx context.Context, req *article.PublishArticleRequest) (resp *article.PublishArticleResponse, err error) {
	err = s.MysqlCli.AddArticle(ctx, req.Content, req.UserId, req.Title)
	if err != nil {
		klog.Errorf("publishArticle failed:%s", err)
		resp = &article.PublishArticleResponse{
			StatusCode: http.StatusBadRequest,
			StatusMsg:  "Publish to database failed ",
			Succed:     false,
		}
		return resp, err
	}
	err = s.RedisCli.CreateArticle(ctx, &model.Article{
		ID:         req.UserId,
		CreateTime: time.Now(),
		Context:    req.Content,
		CreaterID:  req.UserId,
		Title:      req.Title,
		Comments:   nil,
	})
	if err != nil {
		klog.Errorf("publishArticle failed:%s", err)
		resp = &article.PublishArticleResponse{
			StatusCode: http.StatusBadRequest,
			StatusMsg:  "Publish to redis failed ",
			Succed:     false,
		}
		return resp, err
	}
	resp = &article.PublishArticleResponse{
		StatusCode: http.StatusOK,
		StatusMsg:  "Accept",
		Succed:     true,
	}
	return resp, nil
}

// AddComment implements the ArticleServiceImpl interface.
func (s *ArticleServiceImpl) AddComment(ctx context.Context, req *article.AddCommentRequest) (resp *article.AddCommentResponse, err error) {
	err = s.MysqlCli.AddComment(ctx, req.UserId, int64(req.ArticleId), req.Context)
	if err != nil {
		klog.Errorf("addcomment upload to database failed:%s", err)
		resp = &article.AddCommentResponse{
			StatusCode: http.StatusBadRequest,
			StatusMsg:  "addcomment upload to database failed",
			Succed:     false,
		}
		return resp, err
	}
	NewArticle, err := s.MysqlCli.GetArticle(ctx, req.ArticleId)
	if err != nil {
		klog.Errorf("get artcile after add comment falied:%s", err)
		resp = &article.AddCommentResponse{
			StatusCode: http.StatusBadRequest,
			StatusMsg:  "Comment upload to database failed",
			Succed:     false,
		}
		return resp, err
	}
	err = s.RedisCli.CreateArticle(ctx, NewArticle)
	if err != nil {
		klog.Errorf("upload article after add comment to redis falied:%s", err)
		resp = &article.AddCommentResponse{
			StatusCode: http.StatusBadRequest,
			StatusMsg:  "Comment upload to database failed",
			Succed:     false,
		}
		return resp, err
	}
	resp = &article.AddCommentResponse{
		StatusCode: http.StatusOK,
		StatusMsg:  "Accept",
		Succed:     true,
	}
	return resp, nil
}

// DeleteArticle implements the ArticleServiceImpl interface.
func (s *ArticleServiceImpl) DeleteArticle(ctx context.Context, req *article.DeleteArticleRequest) (resp *article.DeleteArticleResponse, err error) {
	err = s.MysqlCli.DeleteArticle(ctx, int64(req.ArticleId))
	if err != nil {
		klog.Errorf("delete article in databse failed:%s", err)
		resp = &article.DeleteArticleResponse{
			StatusCode: http.StatusBadRequest,
			StatusMsg:  "delete article in database failed",
			Succed:     false,
		}
		return resp, err

	}
	s.RedisCli.DeleteArtcileById(ctx, int64(req.ArticleId))
	resp = &article.DeleteArticleResponse{
		StatusCode: 200,
		StatusMsg:  "Accept",
		Succed:     true,
	}
	return resp, nil
}
