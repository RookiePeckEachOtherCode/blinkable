package dao

import (
	"blinkable/server/common/consts"
	"blinkable/server/service/user/dao/query"
	"blinkable/server/service/user/model"
	"context"
	"errors"

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

func (u User) CreateUser(ctx context.Context, user *model.User) error {
	_user, err := u.q.User.WithContext(ctx).Where(u.q.User.Username.Eq(user.Username)).First()
	if err != gorm.ErrRecordNotFound && err != nil {
		klog.Errorf(consts.ErrFindUserIsWrong, err)
		return err
	}
	if _user != nil {
		return errors.New(consts.ErrUserIsExist)
	}
	err = u.q.User.WithContext(ctx).Create(user)
	if err != nil {
		klog.Error(consts.ErrCreateUserIsWrong, err)
		return err
	}

	return nil
}
func (u User) GetUserByUserName(ctx context.Context, username string) (*model.User, error) {
	user, err := u.q.User.WithContext(ctx).Where(u.q.User.Username.Eq(username)).First()
	if err != nil {
		return nil, err
	}
	return user, nil
}
func (u User) GetUserById(ctx context.Context, id int64) (*model.User, error) {
	user, err := u.q.User.WithContext(ctx).Where(u.q.User.ID.Eq(id)).First()
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (u User) UpdateUserInfo(ctx context.Context, user *model.User) error {
	_, err := u.q.User.WithContext(ctx).Where(u.q.User.ID.Eq(user.ID)).Select(u.q.User.Username, u.q.User.Signature).Updates(map[string]interface{}{
		"username":  user.Username,
		"signature": user.Signature,
		"title":     user.Title,
	})
	return err
}
