package utils

import (
	"NcuhomeBlog/conf"
	"time"

	"github.com/dgrijalva/jwt-go"
)

var jwtSecret = []byte(conf.GetConfig().JWTSecret)

type Claims struct {
	ID       int
	Username string
	jwt.StandardClaims
}

func GenerateToken(id int, username string) (token string, err error) {
	nowTime := time.Now()
	expireTime := nowTime.Add(3 * time.Hour)

	claims := Claims{
		id,
		username,
		jwt.StandardClaims{
			ExpiresAt: expireTime.Unix(),
			Issuer:    "NcuhomeBlog",
		},
	}
	tokenClaims := jwt.NewWithClaims(jwt.SigningMethodES256, claims)
	token, err = tokenClaims.SignedString(jwtSecret)
	return token, err
}

func ParseToken(token string) (*Claims, error) {
	tokenClaims, err := jwt.ParseWithClaims(token, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return jwtSecret, nil
	})


	if tokenClaims!=nil {
		if claims, ok :=tokenClaims.Claims.(*Claims); ok&&tokenClaims.Valid {
			return claims, nil
		}
	} 
		return nil,err
	
}
