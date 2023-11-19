package main

import (
	"blinkable/common/errno"
	"blinkable/dal/db"
	"blinkable/dal/db/model"
	"blinkable/dal/db/query"
	"blinkable/kitex_gen/Mainview"
	"context"
	"go.uber.org/zap"
)

func init() {
	query.SetDefault(db.DB)
}

// MainviewServiceImpl implements the last service interface defined in the IDL.
type MainviewServiceImpl struct{}

// LikeAction implements the MainviewServiceImpl interface.
func (s *MainviewServiceImpl) LikeAction(ctx context.Context, req *mainview.LikeRequest) (resp *mainview.LikeResponse, err error) {
	ud := query.Q.Admin
	_, err = ud.WithContext(ctx).Where(ud.ID.Eq(req.AdminId)).Update(ud.Like, ud.Like.Add(1))
	if err != nil {
		resp = &mainview.LikeResponse{
			StatusCode: -1,
			StatusMsg:  errno.ErrLike.Error(),
			Succed:     false,
		}
		zap.S().Errorf("%v ===> %v", errno.ErrLike, resp.StatusMsg)
		return resp, nil
	} else {
		resp = &mainview.LikeResponse{
			StatusCode: 0,
			StatusMsg:  "点赞成功",
			Succed:     true,
		}
	}
	return
}
func convertToAdminType(dbAdmins []*model.Admin) []*mainview.Admin {
	var result []*mainview.Admin

	for _, dbAdmin := range dbAdmins {
		admin := &mainview.Admin{
			AdminId:   dbAdmin.ID,
			Signature: dbAdmin.Signature,
			Like:      int32(dbAdmin.Like),
			Title:     dbAdmin.Title,
			Comments:  int32(dbAdmin.Comments),
			IconUrl:   dbAdmin.IconURL,
			ImageUrl:  dbAdmin.ImageURL,
		}

		// Convert adminHasManyGuestbooks to []*Guestbook
		for _, guestbook := range dbAdmin.Guestbooks {
			admin.Guestbooks = append(admin.Guestbooks, &mainview.Guestbook{
				BookId:     guestbook.ID,
				UserId:     int32(guestbook.UserID),
				Context:    guestbook.Context,
				AdminId:    int32(guestbook.AdminID),
				CreateTime: guestbook.CreateAt.Format("2006-01-02 15:04:05"),
			})
		}

		result = append(result, admin)
	}

	return result
}

// GetMainview implements the MainviewServiceImpl interface.
func (s *MainviewServiceImpl) GetMainview(ctx context.Context, req *mainview.GetMainvewRequest) (resp *mainview.GetMainviewResponse, err error) {
	ud := query.Q.Admin
	admins, err := ud.WithContext(ctx).Where(ud.ID.In(1, 2, 3)).Find()
	if err != nil {
		resp = &mainview.GetMainviewResponse{
			Admins:     nil,
			StatusCode: -1,
			StatusMsg:  errno.ErrGetAdmins.Error(),
			Succed:     false,
		}
		zap.S().Errorf("%v ===> %v", errno.ErrGetAdmins, resp.StatusMsg)
	} else {
		resp = &mainview.GetMainviewResponse{
			Admins:     convertToAdminType(admins),
			StatusCode: 0,
			StatusMsg:  "获取管理信息列表成功",
			Succed:     true,
		}
	}
	return
}

// AddGuestbook implements the MainviewServiceImpl interface.
func (s *MainviewServiceImpl) AddGuestbook(ctx context.Context, req *mainview.AddGuestbookRequest) (resp *mainview.AddGuestbookResponse, err error) {
	ud := query.Q.Guestbook
	err = ud.WithContext(ctx).Select(ud.Context, ud.UserID, ud.AdminID).Create(&model.Guestbook{
		AdminID: uint32(req.AdminId),
		Context: req.Context,
		UserID:  uint32(req.UserId),
	})
	if err != nil {
		resp = &mainview.AddGuestbookResponse{
			StatusCode: -1,
			StatusMsg:  errno.ErrAddGuestbook.Error(),
			Succed:     false,
		}
		zap.S().Errorf("%v ===> %v", errno.ErrAddGuestbook, resp.StatusMsg)
	} else {
		resp = &mainview.AddGuestbookResponse{
			StatusCode: 0,
			StatusMsg:  "留言成功",
			Succed:     true,
		}
	}
	return
}
