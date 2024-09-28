package jwt

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"github.com/sail-host/cloud/internal/constants"
	"github.com/sail-host/cloud/internal/global"
)

type JWT struct {
	SigningKey []byte
}

type JwtRequest struct {
	BaseClaims
	BufferTime int64
	jwt.RegisteredClaims
}

type CustomClaims struct {
	BaseClaims
	BufferTime int64
	jwt.RegisteredClaims
}

type BaseClaims struct {
	ID   uint
	Name string
}

func NewJWT() *JWT {
	return &JWT{
		[]byte(global.CONF.System.EncryptKey),
	}
}

func (j *JWT) CreateClaims(baseClaims BaseClaims) CustomClaims {
	claims := CustomClaims{
		BaseClaims: baseClaims,
		BufferTime: constants.JWTBufferTime,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Second * time.Duration(constants.JWTBufferTime))),
			Issuer:    constants.JWTIssuer,
		},
	}
	return claims
}

func (j *JWT) CreateToken(request CustomClaims) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &request)
	return token.SignedString(j.SigningKey)
}

func (j *JWT) ParseToken(tokenStr string) (*JwtRequest, error) {
	token, err := jwt.ParseWithClaims(tokenStr, &JwtRequest{}, func(token *jwt.Token) (interface{}, error) {
		return j.SigningKey, nil
	})
	if err != nil || token == nil {
		return nil, errors.New(constants.ErrTokenParse)
	}
	if claims, ok := token.Claims.(*JwtRequest); ok && token.Valid {
		return claims, nil
	}
	return nil, errors.New(constants.ErrTokenParse)
}
