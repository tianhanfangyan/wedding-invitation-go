package utils

import (
	"errors"
	"github.com/dgrijalva/jwt-go"
	"time"
)

var (
	jwtSecret        = []byte(GetConf().String("app::jwt_secret"))
	TokenExpired     = errors.New("令牌授权过期")
	TokenNotValidYet = errors.New("令牌没有激活")
	TokenMalformed   = errors.New("令牌未知格式")
	TokenInvalid     = errors.New("令牌无法解析")
)

// Custom claims structure
type CustomClaims struct {
	Username   string
	Password   string
	BufferTime int64
	jwt.StandardClaims
}

// 生成token
func GenerateToken(userName, passWord string) (string, error) {
	claims := CustomClaims{
		Username: EncodeMD5(userName),
		Password: EncodeMD5(passWord),
		StandardClaims: jwt.StandardClaims{
			NotBefore: time.Now().Unix() - 1000,              // 签名生效时间
			ExpiresAt: time.Now().Add(time.Hour * 12).Unix(), // 过期时间
			Issuer:    "weeding-invitation",                  // 签名的发行者
		},
	}

	tokenClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	token, err := tokenClaims.SignedString(jwtSecret)

	return token, err
}

// 解析token
func ParseToken(tokenString string) (*CustomClaims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &CustomClaims{}, func(token *jwt.Token) (i interface{}, e error) {
		return jwtSecret, nil
	})

	if err != nil {
		if ve, ok := err.(*jwt.ValidationError); ok {
			if ve.Errors&jwt.ValidationErrorMalformed != 0 {
				return nil, TokenMalformed
			} else if ve.Errors&jwt.ValidationErrorExpired != 0 {
				// Token is expired
				return nil, TokenExpired
			} else if ve.Errors&jwt.ValidationErrorNotValidYet != 0 {
				return nil, TokenNotValidYet
			} else {
				return nil, TokenInvalid
			}
		}
	}

	if token != nil {
		if claims, ok := token.Claims.(*CustomClaims); ok && token.Valid {
			return claims, nil
		}
		return nil, TokenInvalid
	} else {
		return nil, TokenInvalid

	}
}
