package init

import (
	"blinkable/server/service/api/config"
	"context"
	"errors"

	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
)

func InitMinio() {
	minioClient, err := minio.New(config.GlobalServerConfig.MinioInfo.EndPoint, &minio.Options{
		Creds:  credentials.NewStaticV4(config.GlobalServerConfig.MinioInfo.AccessKeyID, config.GlobalServerConfig.MinioInfo.SecretAccessKey, ""),
		Secure: false,
	})
	if err != nil {
		hlog.Fatalf("minio init failed: %s", err)
	}

	config.GlobalMinioClient = *minioClient
}
func InitCreateALLBucket() {
	buckets := config.GlobalServerConfig.MinioInfo.Bucket

	for _, name := range buckets {
		err := CreateBucket(name)
		if err != nil {
			hlog.Fatalf("Create bucket %s is failed: %s", name, err)
		}
	}
}

func CreateBucket(bucketName string) error {

	if len(bucketName) == 0 {
		err := errors.New("BucketName is empty")
		return err
	}

	err := config.GlobalMinioClient.MakeBucket(context.Background(), bucketName, minio.MakeBucketOptions{})
	if err != nil {
		return err
	}
	return nil
}
