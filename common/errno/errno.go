package errno

import "errors"

// JWT
var (
	ErrTokenMalformed   = errors.New("token is malformed")     // 令牌格式不正确
	ErrTokenExpired     = errors.New("token is expired")       // 令牌已过期
	ErrTokenNotValidYet = errors.New("token is not valid yet") // 令牌尚未生效
	ErrTokenInvalidId   = errors.New("token has invalid id")   // 令牌的 ID 无效
)
