package model

import "time"

type Comment struct {
	CommentID  int32     `gorm:"column:comment_id;primaryKey;autoIncrement:true;comment:评论id" json:"comment_id"`        // 评论id
	UserID     int32     `gorm:"column:user_id;not null;comment:用户id" json:"user_id"`                                   // 用户id
	ArticleID  int32     `gorm:"column:article_id;not null;comment:文章id" json:"article_id"`                             // 文章id
	Content    string    `gorm:"column:content;not null;comment:评论内容" json:"content"`                                   // 评论内容
	CreateTime time.Time `gorm:"column:create_time;not null;default:CURRENT_TIMESTAMP;comment:创建时间" json:"create_time"` // 创建时间
}
