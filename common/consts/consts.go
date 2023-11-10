package consts

import "blinkable/pkg/viper"

// API
var (
	API_PORT = "8080"
)

// ETCD
var (
	ETCD_ADDR = "127.0.0.1:2379"
)

// UserService
var (
	UserServiceName    = "Blinkable-UserService"
	UserServiceAPIPort = "10001"
	UserServiceRPCPort = "11001"
)

// jwt
var (
	JwtKey = "114514"
)

// minio
var (
	cfg                  = viper.Load("oss").Viper
	AvatarBucketName     = cfg.GetString("bucket.avatar.name")
	BackgroundBucketName = cfg.GetString("bucket.background.name")
)
