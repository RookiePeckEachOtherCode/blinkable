package dao

import (
	"blinkable/server/common/consts"
	"blinkable/server/service/homepage/dao/query"
	model2 "blinkable/server/service/homepage/model"
	"context"

	"github.com/cloudwego/kitex/pkg/klog"
	"gorm.io/gorm"
)

type User struct {
	q *query.Query
}

func NewUser(db *gorm.DB) *User {
	m := db.Migrator()
	if !m.HasTable(&model2.User{}) {
		if err := m.CreateTable(&model2.User{}); err != nil {
			klog.Errorf("create user table failed: %s", err)
		}
	}
	if !m.HasTable(&model2.Guestbook{}) {
		if err := m.CreateTable(&model2.Guestbook{}); err != nil {
			klog.Errorf("create guestbook table failed: %s", err)
		}
	}
	qr := query.Use(db)
	return &User{q: qr}
}

func (u User) GetUsersByIds(ctx context.Context, ids []int64) ([]*model2.User, error) {
	users, err := u.q.User.WithContext(ctx).Where(u.q.User.ID.In(ids...)).Find()
	if err != nil {
		return nil, err
	}
	return users, nil
}

func (u User) GetGuestbooksById(ctx context.Context, id int64) ([]*model2.Guestbook, error) {
	guestbooks, err := u.q.Guestbook.WithContext(ctx).Where(u.q.Guestbook.UserID.Eq(id)).Find()
	if err != nil {
		return nil, err
	}
	return guestbooks, nil
}

func (u User) AddGuestbook(ctx context.Context, guestbook *model2.Guestbook) error {
	err := u.q.Guestbook.WithContext(ctx).Create(guestbook)
	if err != nil {
		klog.Error(consts.ErrAddGuestbook, err)
		return err
	}
	return nil
}
func (u User) AddLikeNumberById(ctx context.Context, id int64) error {
	_, err := u.q.User.WithContext(ctx).Where(u.q.User.ID.Eq(id)).Update(u.q.User.Like, u.q.User.Like.Add(1))
	return err
}
