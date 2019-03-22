package tokentool

import (
	"errors"
	"fmt"
	"time"

	"github.com/dgrijalva/jwt-go"
)

type jwtPri struct {
	rsaPriKey string
}

type jwtPub struct {
	rsaPubKey string
}

// 私钥Key, 用来创建Token
func NewJwtPri(rsaPriKey string) *jwtPri {
	return &jwtPri{
		rsaPriKey: rsaPriKey,
	}
}

// 公钥Key, 用来验证Token
func NewJwtPub(rsaPubKey string) *jwtPub {
	return &jwtPub{
		rsaPubKey: rsaPubKey,
	}
}

// 创建Token
func (this *jwtPri) CreateToken(uuid string, expire time.Duration) (tokenString string, err error) {
	parsedKey, err := jwt.ParseRSAPrivateKeyFromPEM([]byte(this.rsaPriKey))
	if err != nil {
		return
	}

	now := time.Now()
	claims := jwt.StandardClaims{
		Id:        uuid,
		IssuedAt:  now.Unix(),
		ExpiresAt: now.Add(expire).Unix(),
	}

	return jwt.NewWithClaims(jwt.SigningMethodRS512, claims).SignedString(parsedKey)
}

// 验证Token
func (this *jwtPub) VerifyToken(tokenString string) (id string, err error) {
	token, err := jwt.ParseWithClaims(tokenString, &jwt.StandardClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodRSA); !ok {
			return nil, fmt.Errorf("Unexpected signing method %v", token.Header["alg"])
		}
		return jwt.ParseRSAPublicKeyFromPEM([]byte(this.rsaPubKey))
	})
	if err != nil {
		return
	}

	claims, ok := token.Claims.(*jwt.StandardClaims)
	if !ok || !token.Valid {
		return "", errors.New("tokenString not convert *MyClaims")
	}

	now := time.Now().Unix()
	if claims.IssuedAt > now || claims.ExpiresAt < now {
		return "", errors.New("tokenString time expired")
	}

	return claims.Id, nil
}
