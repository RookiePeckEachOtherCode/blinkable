package model

type CommentsArticle struct {
	CommentsID int32 `gorm:"column:comments_id;primaryKey;comment:评论id" json:"comments_id"` // 评论id
	ArticlesID int32 `gorm:"column:articles_id;primaryKey;comment:文章id" json:"articles_id"` // 文章id
}
