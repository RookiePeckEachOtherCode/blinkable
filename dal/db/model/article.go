package model

type Article struct {
	Model
	AuthorID uint32  `gorm:"not null;comment: 作者id"` // 作者id
	Title    string  `gorm:"not null;comment: 标题"`   // 标题
	Label    *string `gorm:"type: json;comment: 标签"` // 标签(json)
	Category *string `gorm:"type: json;comment: 分类"` // 分类(json)
	Content  string  `gorm:"not null;comment: 内容"`   // 内容
}
