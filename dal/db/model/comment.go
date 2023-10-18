package model

type Comment struct {
	Model
	AuthorID  uint32 `gorm:"not null;comment: 作者id"` // 作者id
	ArticleID uint32 `gorm:"not null;comment: 文章id"` // 文章id
	Content   string `gorm:"type: text;not null;comment: 评论内容"`
}
