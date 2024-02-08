package main

import (
	article "blinkable/server/kitex_gen/Article"
	"context"
)

// ArticleServiceImpl implements the last service interface defined in the IDL.
type ArticleServiceImpl struct{}

// GetArticleSum implements the ArticleServiceImpl interface.
func (s *ArticleServiceImpl) GetArticleSum(ctx context.Context, req *article.GetArticleListRequest) (resp *article.GetArticleSumResponse, err error) {
	// TODO: Your code here...
	return
}

// GetArticleList implements the ArticleServiceImpl interface.
func (s *ArticleServiceImpl) GetArticleList(ctx context.Context, req *article.GetArticleListRequest) (resp *article.GetArticleListResponse, err error) {
	// TODO: Your code here...
	return
}

// GetArticle implements the ArticleServiceImpl interface.
func (s *ArticleServiceImpl) GetArticle(ctx context.Context, req *article.GetArticleRequest) (resp *article.GetArticleResponse, err error) {
	// TODO: Your code here...
	return
}

// PublishArticle implements the ArticleServiceImpl interface.
func (s *ArticleServiceImpl) PublishArticle(ctx context.Context, req *article.PublishArticleRequest) (resp *article.PublishArticleResponse, err error) {
	// TODO: Your code here...
	return
}

// AddComment implements the ArticleServiceImpl interface.
func (s *ArticleServiceImpl) AddComment(ctx context.Context, req *article.AddCommentRequest) (resp *article.AddCommentResponse, err error) {
	// TODO: Your code here...
	return
}

// DeleteArticle implements the ArticleServiceImpl interface.
func (s *ArticleServiceImpl) DeleteArticle(ctx context.Context, req *article.DeleteArticleRequest) (resp *article.DeleteArticleResponse, err error) {
	// TODO: Your code here...
	return
}
