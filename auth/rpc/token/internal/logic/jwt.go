package logic

import (
	"rsapi/auth/rpc/token/internal"
	"time"

	"github.com/golang-jwt/jwt"
)

type Claims struct {
	Email  string
	Role   string
	Expire int64
	jwt.StandardClaims
}

func GenerateJwt(in *Claims, secret string) (string, error) {
	in.ExpiresAt = time.Now().Unix() + in.Expire
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, in)
	tokenString, err := token.SignedString([]byte(secret))
	if err != nil {
		return "", err
	}
	return tokenString, nil
}

func ParseJwt(tokenString string, secret string) (*Claims, error) {
	if token, err := jwt.ParseWithClaims(tokenString, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(secret), nil
	}); err != nil {
		return nil, err
	} else if claims, ok := token.Claims.(*Claims); ok && token.Valid {
		return claims, nil
	} else {
		return claims, internal.ErrJwtExpired
	}
}
