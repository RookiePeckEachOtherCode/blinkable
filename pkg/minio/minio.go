package minio

import (
	"blinkable/common/errno"
	"blinkable/pkg/viper"
	zlog "blinkable/pkg/zap"
	"bytes"
	"context"
	"fmt"
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

// 上传图片，返回图片的url和错误
func UpLoadImg(bucketName string, data []byte, typ string) (string, error) {
	reader := bytes.NewReader(data)
	var contentType string

	zap.S().Debugf("typ: %v", typ)
	zap.S().Debugf("bucketName: %v", bucketName)

	switch typ {
	case ".jpeg":
		contentType = "image/jpeg"
	case ".jpg":
		contentType = "image/jpg"
	case ".png":
		contentType = "image/png"
	default:
		return "", fmt.Errorf("%v: [%v]", errno.ErrUploadImgTypeIsWrong, typ)
	}

	fileName := time.Now().Format("2006-01-02-15-04-05") + typ // eg. [2006-01-02-15-04-05.jpg]

	_, err := UpLoadFile(bucketName, fileName, reader, int64(len(data)), contentType)
	if err != nil {
		zap.S().Errorf("%v ===> %v", errno.ErrUploadFile, err)
		return "", err
	}

	url, err := GetFileTemporaryURL(bucketName, fileName)

	if err != nil {
		zap.S().Errorf("%v ===> %v", errno.ErrGetTemporarilyUrlIsWrong, err)
		return "", err
	}
	return url, nil
}
