package models

import (
	"enterprise-api/app/config"
	"github.com/golang-jwt/jwt"
)

type MyCustomClaims struct {
	Id       int    `json:"id"`
	Username string `json:"username"`
	jwt.StandardClaims
}

func CreateToken(claims MyCustomClaims) (string, error) {
	//使用HS256加密方式
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signToken, err := token.SignedString([]byte(config.GetConfig().JWTKey))
	if err != nil {
		return "", err
	}
	return signToken, nil

}

func VerifyToken(signToken string) (*MyCustomClaims, error) {
	var claims MyCustomClaims
	token, err := jwt.ParseWithClaims(signToken, &claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(config.GetConfig().JWTKey), nil
	})
	if token.Valid {
		return &claims, nil
	} else {
		return nil, err
	}
}
