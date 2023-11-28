package model

type Admin struct {
	Model
	// 评论数
	Comments int64 `gorm:"comment: 评论数"`
	// 留言
	Guestbooks []Guestbook `gorm:"foreignKey:AdminID;comment:留言"`
	// 头像
	IconURL string `gorm:"type: varchar(600);comment:头像"`
	// 背景
	ImageURL string `gorm:"type: varchar(600);comment: 名片背景图片"`
	// 点赞数目
	Like int64 `gorm:"comment: 点赞数"`
	// 签名
	Signature string `gorm:"type: varchar(150);comment:个性签名"`
	// 称号？
	Title string `gorm:"type: varchar(150);comment: 称号"`
}
