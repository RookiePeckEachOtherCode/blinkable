package dao

import (
	"blinkable/server/common/consts"
	"blinkable/server/service/homepage/dao/query"
	"blinkable/server/service/homepage/model"
	"context"
	"time"

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

func (u MysqlCen) GetUserById(ctx context.Context, id int64) (*model.User, error) {
	user, err := u.q.User.WithContext(ctx).Where(u.q.User.ID.Eq(id)).First()
	if err != nil {
		return nil, err
	}
	user.Guestbooks, err = u.GetGuestBookListByUserId(ctx, id)
	return user, nil
}

func (u MysqlCen) GetGuestBookListByUserId(ctx context.Context, userId int64) ([]*model.Guestbook, error) {
	query := u.q.Guestbook
	return query.WithContext(ctx).Where(query.UserID.Eq(userId)).Find()
}

func (u MysqlCen) GetUsersByIds(ctx context.Context, ids []int64) ([]*model.User, error) {
	users, err := u.q.User.WithContext(ctx).Where(u.q.User.ID.In(ids...)).Find()
	if err != nil {
		return nil, err
	}
	return users, nil
}

func (u MysqlCen) GetGuestbooksById(ctx context.Context, id int64) ([]*model.Guestbook, error) {
	guestbooks, err := u.q.Guestbook.WithContext(ctx).Where(u.q.Guestbook.UserID.Eq(id)).Find()
	if err != nil {
		return nil, err
	}
	return guestbooks, nil
}

func (u MysqlCen) GetAllAdminUserInfo(ctx context.Context) ([]*model.User, error) {
	query := u.q.User
	res, err := query.WithContext(ctx).Where(query.IsAdmin.Is(true)).Find()
	return res, err
}

func (u MysqlCen) AddGuestbook(ctx context.Context, context string, userId int64, fromUserId int64) error {
	gbk := model.Guestbook{
		CreateTime: time.Now(),
		Context:    context,
		UserID:     userId,
		FromUserID: fromUserId,
	}
	err := u.q.Guestbook.WithContext(ctx).Create(&gbk)
	if err != nil {
		klog.Error(consts.ErrAddGuestbook, err)
		return err
	}
	return nil
}
func (u MysqlCen) AddLikeNumberById(ctx context.Context, id int64) error {
	_, err := u.q.User.WithContext(ctx).Where(u.q.User.ID.Eq(id)).Update(u.q.User.LikeNum, u.q.User.LikeNum.Add(1))
	return err
}
