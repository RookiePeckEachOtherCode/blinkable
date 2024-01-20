package config

import (
	"blinkable/server/kitex_gen/user/userservice"

	"github.com/minio/minio-go/v7"
)

var (
	GlobalServerConfig ServerConfig
	GlobalNacosConfig  NacosConfig

	GlobalUserClient  userservice.Client
	GlobalMinioClient minio.Client
)
