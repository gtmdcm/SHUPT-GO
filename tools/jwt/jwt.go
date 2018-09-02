package jwt

import (
	"github.com/dgrijalva/jwt-go"
	"os"
	"strconv"
	"time"
)

var secret = []byte(os.Getenv("JWT_SECRET"))

func MakeToken(userId uint64) string {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"exp":    time.Now().Unix() + 60*60,
		"iat":    time.Now().Unix(),
		"userId": userId,
	})
	tokenString, err := token.SignedString(secret)
	if err != nil {
		panic("Failed to make jwt token with id " + strconv.FormatUint(userId, 10))
	}
	return tokenString
}
