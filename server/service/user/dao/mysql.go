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

type MysqlCen struct {
	q *query.Query
}

func NewUser(db *gorm.DB) *MysqlCen {
	qr := query.Use(db)
	return &MysqlCen{q: qr}
}

func (u MysqlCen) CreateUser(ctx context.Context, user *model.User) error {
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
func (u MysqlCen) GetUserByUserName(ctx context.Context, username string) (*model.User, error) {
	user, err := u.q.User.WithContext(ctx).Where(u.q.User.Username.Eq(username)).First()
	if err != nil {
		return nil, err
	}
	return user, nil
}
func (u MysqlCen) GetUserById(ctx context.Context, id int64) (*model.User, error) {
	user, err := u.q.User.WithContext(ctx).Where(u.q.User.ID.Eq(id)).First()
	if err != nil {
		return nil, err
	}
	user.Guestbooks, err = u.GetGuestBookListByUserId(ctx, id)
	return user, nil
}

func (u MysqlCen) UpdateUserInfo(ctx context.Context, user *model.User) error {
	_, err := u.q.User.WithContext(ctx).Where(u.q.User.ID.Eq(user.ID)).
		Updates(map[string]interface{}{
			"username":   user.Username,
			"signature":  user.Signature,
			"title":      user.Title,
			"github_url": user.GithubUrl,
		})
	return err
}

func (u MysqlCen) GetGuestBookListByUserId(ctx context.Context, userId int64) ([]*model.Guestbook, error) {
	query := u.q.Guestbook
	return query.WithContext(ctx).Where(query.UserID.Eq(userId)).Find()
}
