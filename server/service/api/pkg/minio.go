package pkg

import (
	"blinkable/server/common/consts"
	"blinkable/server/service/api/config"
	"context"

	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/minio/minio-go/v7"
)

func MinioAvatarUpload(tmpFilePath string, fileName string) error {
	_, err := config.GlobalMinioClient.FPutObject(context.Background(), config.GlobalServerConfig.MinioInfo.Bucket[consts.AvatarBucket], fileName, tmpFilePath, minio.PutObjectOptions{
		ContentType: "image/png",
	})
	if err != nil {
		hlog.Error("minio avatar upload failed,", err)
		return err
	}
	return nil
}
func MinioBackgroundUpload(tmpFilePath string, fileName string) error {
	_, err := config.GlobalMinioClient.FPutObject(context.Background(), config.GlobalServerConfig.MinioInfo.Bucket[consts.BackgroundBucket], fileName, tmpFilePath, minio.PutObjectOptions{
		ContentType: "image/png",
	})
	if err != nil {
		hlog.Error("minio avatar upload failed,", err)
		return err
	}
	return nil
}
