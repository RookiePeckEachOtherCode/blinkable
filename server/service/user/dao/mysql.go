package dao

import (
	"blinkable/server/common/consts"
	"blinkable/server/service/user/model"
	"context"
	"errors"

	"github.com/cloudwego/kitex/pkg/klog"
	"gorm.io/gorm"
)

type User struct {
	db *gorm.DB
}

func NewUser(db *gorm.DB) *User {
	m := db.Migrator()
	if !m.HasTable(&model.User{}) {
		if err := m.CreateTable(&model.User{}); err != nil {
			klog.Errorf("create user table failed: %s", err)
		}
	}
	return &User{db: db}
}
func (u User) CreateUser(ctx context.Context, user *model.User) error {
	var _user model.User
	err := u.db.Where("username = ?", user.Username).First(&_user).Error
	if err != gorm.ErrRecordNotFound && err != nil {
		klog.Errorf(consts.ErrFindUserIsWrong, err)
		return err
	}
	if _user.Username != "" {
		return errors.New(consts.ErrUserIsExist)
	}
	err = u.db.Create(&user).Error
	if err != nil {
		klog.Error(consts.ErrCreateUserIsWrong, err)
		return err
	}

	return nil
}
func (u User) GetUserByUserName(ctx context.Context, username string) (*model.User, error) {
	var user model.User
	if err := u.db.Where("username = ?", username).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}
func (u User) GetUserById(ctx context.Context, id int64) (*model.User, error) {
	var user model.User
	if err := u.db.Where("id = ?", id).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}
