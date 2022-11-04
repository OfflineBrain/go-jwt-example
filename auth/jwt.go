package auth

import (
	"errors"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
)

var jwtKey = []byte("super-secret-key")

type JWTClaim struct {
	jwt.StandardClaims
	Username string `json:"username"`
	Email    string `json:"email"`
}

func GenerateJWT(email string, username string) (tokenString string, err error) {
	expTime := time.Now().Add(15 * time.Minute)

	claims := &JWTClaim{
		Email:    email,
		Username: username,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expTime.Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, err = token.SignedString(jwtKey)
	return
}

func ValidateToken(tokenString string) (err error) {
	token, err := jwt.ParseWithClaims(tokenString, &JWTClaim{},
		func(t *jwt.Token) (interface{}, error) {
			return []byte(jwtKey), nil
		})
	if err != nil {
		return
	}

	claims, ok := token.Claims.(*JWTClaim)

	if !ok {
		err = errors.New("couldn't parse claims")
		return
	}

	if claims.ExpiresAt < time.Now().Local().Unix() {
		err = errors.New("token expired")
		return
	}
	return
}
