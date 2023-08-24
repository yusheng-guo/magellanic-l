package service

import (
	"errors"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/yushengguo557/magellanic-l/global"
	"time"
)

type jwtService struct {
}

var JwtService = new(jwtService)

type JWTClaim struct {
	jwt.StandardClaims
}

type JwtUser interface {
	GetID() string
}

// GenerateToken 生成 token
func (s *jwtService) GenerateToken(user JwtUser) (tokenString string, err error) {
	expirationTime := time.Now().Add(global.App.Config.Jwt.TTL)
	claim := JWTClaim{
		StandardClaims: jwt.StandardClaims{
			Id:        user.GetID(),
			ExpiresAt: expirationTime.Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claim)
	tokenString, err = token.SignedString(global.App.Config.Jwt.Key)
	if err != nil {
		return "", fmt.Errorf("sign token, err: %w", err)
	}
	return tokenString, nil
}

// ValidateToken 验证 token
func (s *jwtService) ValidateToken(signedToken string) (err error) {
	token, err := jwt.ParseWithClaims(
		signedToken,
		&JWTClaim{},
		func(token *jwt.Token) (interface{}, error) {
			return []byte(global.App.Config.Jwt.Key), nil
		},
	)
	if err != nil {
		return fmt.Errorf("parse signed token, err: %w", err)
	}
	claims, ok := token.Claims.(*JWTClaim)
	if !ok {
		return errors.New("couldn't parse claims")
	}
	if claims.ExpiresAt < time.Now().Local().Unix() {
		return errors.New("token expired")
	}
	return nil
}
