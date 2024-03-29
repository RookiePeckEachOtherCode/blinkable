package config

import (
	article "blinkable/server/kitex_gen/Article/articleservice"
	"blinkable/server/kitex_gen/Homepage/homepageservice"
	"blinkable/server/kitex_gen/user/userservice"

	"github.com/minio/minio-go/v7"
)

var (
	GlobalServerConfig ServerConfig
	GlobalNacosConfig  NacosConfig

	GlobalUserClient     userservice.Client
	GlobalMinioClient    minio.Client
	GlobalHomepageClient homepageservice.Client
	GlobalArticleClient  article.Client
)
