package Jwt

import (
	"blinkable/common/consts"
	"blinkable/common/errno"

	"github.com/dgrijalva/jwt-go"
)

type Jwt struct {
	SigningKey []byte
}

type Claims struct {
	Id          int64
	AuthorityId int64
	jwt.StandardClaims
}

func NewJwt() *Jwt {
	return &Jwt{SigningKey: []byte(consts.JwtKey)}
}

func (j *Jwt) ReleaseToken(claim Claims) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claim)
	return token.SignedString(j.SigningKey)
}
func (j *Jwt) ParseToken(tokenStr string) (*Claims, error) {
	token, err := jwt.ParseWithClaims(tokenStr, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return j.SigningKey, nil
	})

	if err != nil {
		if ve, ok := err.(*jwt.ValidationError); ok {
			if ve.Errors&jwt.ValidationErrorMalformed != 0 { // token格式错误
				return nil, errno.ErrTokenMalformed
			} else if ve.Errors&jwt.ValidationErrorExpired != 0 { // token过期
				return nil, errno.ErrTokenExpired
			} else if ve.Errors&jwt.ValidationErrorNotValidYet != 0 { // token尚未生效
				return nil, errno.ErrTokenNotValidYet
			} else { // token无效
				return nil, errno.ErrTokenInvalidId
			}
		}
	}
	if claims, ok := token.Claims.(*Claims); ok && token.Valid { // token有效
		return claims, nil
	}
	return nil, errno.ErrTokenInvalidId
}
