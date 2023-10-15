package model

import "time"

type Article struct {
	ArticleID  int32     `gorm:"column:article_id;primaryKey;autoIncrement:true;comment:文章id" json:"article_id"`        // 文章id
	AuthorID   int32     `gorm:"column:author_id;not null;comment:作者id" json:"author_id"`                               // 作者id
	Title      string    `gorm:"column:title;not null;comment:文章标题" json:"title"`                                       // 文章标题
	Label      string    `gorm:"column:label;not null;comment:文章标签" json:"label"`                                       // 文章标签
	Category   string    `gorm:"column:category;not null;comment:文章分类" json:"category"`                                 // 文章分类
	Content    string    `gorm:"column:content;not null;comment:文章内容" json:"content"`                                   // 文章内容
	CreateTime time.Time `gorm:"column:create_time;not null;default:CURRENT_TIMESTAMP;comment:创建时间" json:"create_time"` // 创建时间
	UpdateTime time.Time `gorm:"column:update_time;not null;comment:更新时间" json:"update_time"`                           // 更新时间
}
