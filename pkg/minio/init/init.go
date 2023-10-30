package main

import (
	"blinkable/pkg/minio"
	"blinkable/pkg/viper"
	zlog "blinkable/pkg/zap"

	"go.uber.org/zap"
)

var (
	cfg                  = viper.Load("oss").Viper
	avatarBucketName     = cfg.GetString("bucket.avatar.name")
	backgroundBucketName = cfg.GetString("bucket.background.name")
)

func main() {
	zlog.Init()
	errs := []error{}

	errs = append(errs, minio.CreateBucket(avatarBucketName))
	errs = append(errs, minio.CreateBucket(backgroundBucketName))
	for _, err := range errs {
		if err != nil {
			zap.S().Errorf("%v ===> %v", "创建bucket失败", err)
		}
	}
}
