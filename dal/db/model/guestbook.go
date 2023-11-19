package model

type Guestbook struct {
	Model
	// 名片id
	AdminID uint32 `gorm:"not null;comment: 管理员id"`
	// 评论内容
	Context string `gorm:"type: text;not null;comment: 留言内容"`
	// 发布者id
	UserID uint32 `gorm:"not null;comment: 用户id"`
}
