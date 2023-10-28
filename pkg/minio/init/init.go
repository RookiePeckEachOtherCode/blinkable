package main

import (
	"blinkable/pkg/minio"
	zlog "blinkable/pkg/zap"

	"go.uber.org/zap"
)

func main() {
	zlog.Init()
	errs := []error{}

	errs = append(errs, minio.CreateBucket("avatar"))
	errs = append(errs, minio.CreateBucket("background"))
	for _, err := range errs {
		if err != nil {
			zap.S().Errorf("%v ===> %v", "创建bucket失败", err)
		}
	}
}
