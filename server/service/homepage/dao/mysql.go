package dao

import (
	"blinkable/server/service/homepage/dao/query"
	"blinkable/server/service/user/model"

	"github.com/cloudwego/kitex/pkg/klog"
	"gorm.io/gorm"
)

type User struct {
	q *query.Query
}

func NewUser(db *gorm.DB) *User {
	m := db.Migrator()
	if !m.HasTable(&model.User{}) {
		if err := m.CreateTable(&model.User{}); err != nil {
			klog.Errorf("create user table failed: %s", err)
		}
	}
	qr := query.Use(db)
	return &User{q: qr}
}
