package errno

import (
	"errors"
)

// JWT
var (
	ErrTokenMalformed   = errors.New("token is malformed")     // 令牌格式不正确
	ErrTokenExpired     = errors.New("token is expired")       // 令牌已过期
	ErrTokenNotValidYet = errors.New("token is not valid yet") // 令牌尚未生效
	ErrTokenInvalidId   = errors.New("token has invalid id")   // 令牌的 ID 无效
	ErrTokenRelease     = errors.New("token release failed")   // 令牌颁发失败
)

// User
var (
	ErrUserIsExist               = errors.New("用户已存在")
	ErrUserNotExist              = errors.New("用户不存在")
	ErrUserNameOrPasswordIsWrong = errors.New("用户名或密码错误")
	ErrUserNameOrPasswordIsEmpty = errors.New("用户名或密码不能为空")
	ErrUserNameOrPasswordIsLong  = errors.New("用户名或密码过长")
	ErrHashPasswordIsWrong       = errors.New("密码加密失败")
	ErrCreateUserIsWrong         = errors.New("创建用户失败")
	ErrFindUserIsWrong           = errors.New("查找用户失败")
	ErrLogin                     = errors.New("登录失败")
	ErrRegister                  = errors.New("注册失败")
)

// Server
var (
	ErrInternalServerError = errors.New("服务内部错误")
)
