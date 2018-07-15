package auth

import (
	"encoding/json"
	"errors"
	"github.com/astaxie/beego/orm"
	"github.com/dgrijalva/jwt-go"
	"os"
	"time"
)

var secret = []byte(os.Getenv("JWT_SECRET"))

func makeTokenFor(user User) (signedString string, err error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"exp":    time.Now().Unix() + 60*60,
		"iat":    time.Now().Unix(),
		"userId": user.Id,
	})
	signedString, err = token.SignedString(secret)
	return signedString, err
}

func validateToken(jwtString string) (user *User, err error) {
	token, err := jwt.Parse(jwtString, func(token *jwt.Token) (interface{}, error) {
		return secret, nil
	})
	if err != nil || jwt.SigningMethodHS256.Alg() != token.Header["alg"] || !token.Valid {
		return nil, errors.New("jwt token is invalid")
	}
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return nil, errors.New("jwt token is invalid")
	}
	ormObject := orm.NewOrm()
	user = new(User)
	user.Id = claims["userId"].(int64)
	err = ormObject.Read(user)
	return user, err
}

type TokenJson struct {
	Token string `json:"token"`
}

func makeTokenJson(user User) (token []byte, err error) {
	signedString, err := makeTokenFor(user)
	if err != nil {
		return json.Marshal(TokenJson{Token: ""})
	}
	return json.Marshal(TokenJson{Token: signedString})
}
