package model

import "time"

type User struct {
	UserID      int32     `gorm:"column:user_id;primaryKey;autoIncrement:true;comment:用户id" json:"user_id"`              // 用户id
	UserName    string    `gorm:"column:user_name;not null;comment:用户名" json:"user_name"`                                // 用户名
	Avatar      string    `gorm:"column:avatar;not null;comment:头像地址" json:"avatar"`                                     // 头像地址
	ArticlesNum int32     `gorm:"column:articles_num;not null;comment:文章数量" json:"articles_num"`                         // 文章数量
	Level       int32     `gorm:"column:level;not null;default:1;comment:等级" json:"level"`                               // 等级
	Signature   string    `gorm:"column:signature;not null;comment:签名" json:"signature"`                                 // 签名
	Experience  int32     `gorm:"column:experience;not null;comment:经验值" json:"experience"`                              // 经验值
	CreateTime  time.Time `gorm:"column:create_time;not null;default:CURRENT_TIMESTAMP;comment:创建时间" json:"create_time"` // 创建时间
}
