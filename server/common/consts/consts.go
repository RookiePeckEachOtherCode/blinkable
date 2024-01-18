package consts

// import "blinkable/pkg/viper"

// API
// const (
//
//	API_PORT = "8080"
//
// )
//
// // ETCD
// const (
//
//	ETCD_ADDR = "127.0.0.1:2379"
//
// )
//
// // UserService
// const (
//
//	UserServiceName    = "Blinkable-UserService"
//	UserServiceAPIPort = "10001"
//	UserServiceRPCPort = "11001"
//
// )
//
// // jwt
// const (
//
//	JwtKey = "114514"
//
// )
//
// // minio
// const (
//
//	cfg                  = viper.Load("oss".Viper
//	AvatarBucketName     = cfg.GetString("bucket.avatar.name"
//	BackgroundBucketName = cfg.GetString("bucket.background.name"
//
// )
// PATH
const (
	KlogFilePath   = "./tmp/klog/logs"
	UserConfigPath = "../../service/user/config.yaml"
)

// SnowflakeNode    = 1
const (
	NacosSnowflakeNode = 1
	UserSnowflakeNode  = 2
)

// Redis
const (
	RedisUserDB = 1
)

// Err
const (
	ErrUserIsExist               = "用户已存在"
	ErrUserNotExist              = "用户不存在"
	ErrUserNameOrPasswordIsWrong = "用户名或密码错误"
	ErrUserNameOrPasswordIsEmpty = "用户名或密码不能为空"
	ErrUserNameOrPasswordIsLong  = "用户名或密码过长"
	ErrHashPasswordIsWrong       = "密码加密失败"
	ErrCreateUserIsWrong         = "创建用户失败"
	ErrFindUserIsWrong           = "查找用户失败"
	ErrLogin                     = "登录失败"
	ErrRegister                  = "注册失败"
	ErrUserNotFound              = "用户没有找到"
	ErrUserInfo                  = "获取用户信息失败"
	ErrUserUpdateInfo            = "更新用户信息失败"
	ErrAddGuestbook              = "无法添加留言"
	ErrGetAdmins                 = "未能正确获取管理员信息"
	ErrChangeCard                = "未能变更卡片"
	ErrMinioInit                 = "minio对象初始化失败"
	ErrMinioBuket                = "查询桶时出现错误"
	Errupload                    = "未能正确上传到桶"
	Errsave                      = "桶未能储存数据"
	ErrLike                      = "喜欢失败"

	ErrTokenMalformed   = "token is malformed"     // 令牌格式不正确
	ErrTokenExpired     = "token is expired"       // 令牌已过期
	ErrTokenNotValidYet = "token is not valid yet" // 令牌尚未生效
	ErrTokenInvalidId   = "token has invalid id"   // 令牌的 ID 无效
	ErrTokenRelease     = "token release failed"   // 令牌颁发失败
	ErrTokenParse       = "token parse failed"     // 令牌解析失败

	ErrInternalServerError  = "服务内部错误"
	ErrRequestParamsIsWrong = "请求参数错误"

	ErrNewBucketIsWrong              = "创建 bucket 失败"
	ErrBucketNameIsEmpty             = "bucket name 不能为空"
	ErrBucketNameOrObjectNameIsEmpty = "bucket name 或 object name 不能为空"
	ErrGetTemporarilyUrlIsWrong      = "获取临时 url 失败"
	ErrUploadFile                    = "上传文件失败"
	ErrUploadImgTypeIsWrong          = "上传图片类型错误"
)
