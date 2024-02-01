package consts

const (
	KlogFilePath       = "./tmp/klog/logs"
	HlogFIlePath       = "./tmp/hlog/logs"
	UserConfigPath     = "../../service/user/config.yaml"
	APIConfigPath      = "../../service/api/config.yaml"
	HomepageConfigPath = "../../service/homepage/config.yaml"
)

// SnowflakeNode    = 1
const (
	NacosSnowflakeNode = 1
	UserSnowflakeNode  = 2
	MinioSnowflakeNode = 3
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

	ErrUserInfoNotChange = "userinfo not change"
)

const (
	DSN = "%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local"
)

const (
	AvatarBucket     = 0
	BackgroundBucket = 1
)

var AdminIds = [3]int{1, 2, 3}
