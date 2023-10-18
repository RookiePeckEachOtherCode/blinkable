package model

import "time"

type Model struct {
	ID       uint32    `gorm:"primaryKey;comment: id"` //ID
	CreateAt time.Time `gorm:"comment: 创建时间"`          //创建时间
	UpdateAt time.Time `gorm:"comment: 更新时间"`          //更新时间
}
