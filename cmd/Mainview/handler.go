package main

import (
	"blinkable/common/errno"
	"blinkable/dal/db"
	"blinkable/dal/db/model"
	"blinkable/dal/db/query"
	"blinkable/kitex_gen/Mainview"
	"bytes"
	"context"
	"fmt"
	"github.com/google/uuid"
	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
	"go.uber.org/zap"
	"log"
	"net/url"
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

// ChangeCard implements the MainviewServiceImpl interface.
func (s *MainviewServiceImpl) ChangeCard(ctx context.Context, req *mainview.ChangeCardRequest) (resp *mainview.ChangeCardResponse, err error) {

	endpoint := "localhost:9000"
	accessKey := "your-access-key"
	secretKey := "your-secret-key"
	useSSL := false
	ud := query.Q.Admin
	minioClient, err := minio.New(endpoint, &minio.Options{
		Creds:  credentials.NewStaticV4(accessKey, secretKey, ""),
		Secure: useSSL,
	})

	if err != nil {
		resp = &mainview.ChangeCardResponse{
			StatusCode: -1,
			StatusMsg:  "minio初始化失败",
			Succed:     false,
		}
		zap.S().Errorf("%v ===> %v", errno.ErrMinioInit, resp.StatusMsg)
	}

	err = minioClient.MakeBucket(context.Background(), "iocn", minio.MakeBucketOptions{})
	if err != nil {
		// 检查存储桶是否已经存在。
		exists, err := minioClient.BucketExists(context.Background(), "icon")
		if err == nil && exists {
			log.Printf("存储桶 %s 已经存在", "icon")
		} else {
			zap.S().Errorf("%v ===> %v", errno.ErrMinioBuket, err.Error())
		}
	}
	err = minioClient.MakeBucket(context.Background(), "image", minio.MakeBucketOptions{})
	if err != nil {
		// 检查存储桶是否已经存在。
		exists, err := minioClient.BucketExists(context.Background(), "image")
		if err == nil && exists {
			log.Printf("存储桶 %s 已经存在", "icon")
		} else {
			zap.S().Errorf("%v ===> %v", errno.ErrMinioBuket, err.Error())
		}
	}

	log.Printf("创建存储桶 %s 成功", "icon")
	imageData := req.Icon
	IconNmae := uuid.New().String() + ".png"
	imageBuffer := bytes.NewReader(imageData)
	contentType := "image/png"
	// 使用 PutObject 将图片上传到 "icon" 存储桶
	IconSize, err := minioClient.PutObject(context.Background(), "icon", IconNmae, imageBuffer, int64(imageBuffer.Len()), minio.PutObjectOptions{
		ContentType: contentType,
	})
	if err != nil {
		// 处理上传失败的情况
		resp = &mainview.ChangeCardResponse{
			StatusCode: -1,
			StatusMsg:  "图片上传失败",
			Succed:     false,
		}
		zap.S().Errorf("%v ===> %v", errno.Errupload, err.Error())
	}
	imageData = req.Image
	imageNmae := uuid.New().String() + ".png"
	imageBuffer = bytes.NewReader(imageData)
	IamgeSize, err := minioClient.PutObject(context.Background(), "image", imageNmae, imageBuffer, int64(imageBuffer.Len()), minio.PutObjectOptions{
		ContentType: contentType,
	})
	if err != nil {
		// 处理上传失败的情况
		resp = &mainview.ChangeCardResponse{
			StatusCode: -1,
			StatusMsg:  "图片上传失败",
			Succed:     false,
		}
		zap.S().Errorf("%v ===> %v", errno.Errupload, err.Error())
	}
	log.Printf("imagesize:%v\n iocnsize:%v", IamgeSize, IconSize)
	_, err = minioClient.StatObject(context.Background(), "icon", IconNmae, minio.StatObjectOptions{})
	if err != nil {
		resp = &mainview.ChangeCardResponse{
			StatusCode: -1,
			StatusMsg:  "保存头像失败",
			Succed:     false,
		}
		zap.S().Errorf("%v ===> %v", errno.Errsave, err.Error())
	}
	IconURL, err := url.Parse(fmt.Sprintf("https://%s/%s/%s", endpoint, "icon", IconNmae))
	if err != nil {
		resp = &mainview.ChangeCardResponse{
			StatusCode: -1,
			StatusMsg:  "获取头像url失败",
			Succed:     false,
		}
		zap.S().Errorf("%v ===> %v", errno.Errsave, err.Error())
	}
	_, err = minioClient.StatObject(context.Background(), "image", imageNmae, minio.StatObjectOptions{})
	if err != nil {
		resp = &mainview.ChangeCardResponse{
			StatusCode: -1,
			StatusMsg:  "保存背景失败",
			Succed:     false,
		}
		zap.S().Errorf("%v ===> %v", errno.Errsave, err.Error())
	}
	ImageURL, err := url.Parse(fmt.Sprintf("https://%s/%s/%s", endpoint, "iamge", imageNmae))
	if err != nil {
		resp = &mainview.ChangeCardResponse{
			StatusCode: -1,
			StatusMsg:  "获取背景url失败",
			Succed:     false,
		}
		zap.S().Errorf("%v ===> %v", errno.Errsave, err.Error())
	}
	_, err = ud.WithContext(ctx).Where(ud.ID.Eq(req.AdminId)).Updates(model.Admin{
		IconURL:  IconURL.String(),
		ImageURL: ImageURL.String(),
	})
	if err != nil {
		resp = &mainview.ChangeCardResponse{
			StatusCode: -1,
			StatusMsg:  "更新数据库失败",
			Succed:     false,
		}
		zap.S().Errorf("%v ===> %v", errno.Errsave, err.Error())
	}
	resp = &mainview.ChangeCardResponse{
		StatusCode: 0,
		StatusMsg:  "更新成功",
		Succed:     true,
	}
	return
}
