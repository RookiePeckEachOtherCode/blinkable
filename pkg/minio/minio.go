package minio

import (
	"blinkable/common/errno"
	"blinkable/pkg/viper"
	zlog "blinkable/pkg/zap"
	"context"
	"io"
	"time"

	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
	"go.uber.org/zap"
)

var (
	cfg             = viper.Load("oss").Viper
	minioUrl        = cfg.GetString("endpoint")
	accessID        = cfg.GetString("accessKeyID")
	secretAccessKey = cfg.GetString("secretAccessKey")
	expireTime      = cfg.GetInt("expireTime")
	useSSL          = cfg.GetBool("useSSL")
	minioClient     *minio.Client
)

func init() {
	zlog.Init()

	_minioClient, err := minio.New(minioUrl, &minio.Options{
		Creds:  credentials.NewStaticV4(accessID, secretAccessKey, ""),
		Secure: useSSL,
	})

	if err != nil {
		zap.S().Panicf("%v ===> %v", "minio init error", err)
	}

	minioClient = _minioClient
}

func CreateBucket(bucketName string) error {
	if len(bucketName) == 0 {
		err := errno.ErrBucketNameIsEmpty
		zap.S().Errorf("%v ===> %v", errno.ErrNewBucketIsWrong, err)
		return err
	}

	err := minioClient.MakeBucket(context.Background(), bucketName, minio.MakeBucketOptions{})
	if err != nil {
		zap.S().Errorf("%v ===> %v", errno.ErrNewBucketIsWrong, err)
		return err
	}
	return nil
}

func GetFileTemporaryURL(bucketName, objectName string) (string, error) {
	if len(bucketName) <= 0 || len(objectName) <= 0 {
		return "", errno.ErrBucketNameOrObjectNameIsEmpty
	}
	expiry := time.Second * time.Duration(expireTime)
	presignedURL, err := minioClient.PresignedGetObject(context.Background(), bucketName, objectName, expiry, nil)

	if err != nil {
		zap.S().Errorf("%v ===> %v", errno.ErrGetTemporarilyUrlIsWrong, err)
		return "", err
	}
	return presignedURL.String(), nil
}

func UpLoadFile(bucketName string, fileName string, reader io.Reader, size int64, contentType string) (int64, error) {
	if len(bucketName) <= 0 || len(fileName) <= 0 {
		return -1, errno.ErrBucketNameOrObjectNameIsEmpty
	}

	uploadInfo, err := minioClient.PutObject(context.Background(), bucketName, fileName, reader, size, minio.PutObjectOptions{
		ContentType: contentType,
	})

	if err != nil {
		zap.S().Errorf("%v ===> %v", errno.ErrUploadFile, err)
		return -1, err
	}

	return uploadInfo.Size, nil
}
