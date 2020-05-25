package utils

import (
	"encoding/json"
	"fmt"
	"github.com/ampazdev/simple-go-project/svc/userservice/constant"
	"github.com/ampazdev/simple-go-project/svc/userservice/internal/entity"
	"github.com/dgrijalva/jwt-go"
	"log"
	"net/http"
	"os"
)

type ParamsToken struct {
	Email string
	Lang  string
}

func RespondWithError(w http.ResponseWriter, status int, error entity.Error) {
	w.WriteHeader(status)
	w.Header().Set("Content-type", "application/json")
	json.NewEncoder(w).Encode(error)
}

func ResponseJSON(w http.ResponseWriter, data interface{}) {
	w.Header().Set("Content-type", "application/json")
	json.NewEncoder(w).Encode(data)
}

func GenerateToken(user entity.User) (string, error) {
	var err error
	secret := os.Getenv("SECRET")

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"email": user.Email,
		"iss":   "course",
	})
	tokenString, err := token.SignedString([]byte(secret))
	if nil != err {
		log.Println(err)
	}
	return tokenString, nil
}

func ExtractClaims(tokenStr string) (jwt.MapClaims, bool) {
	hmacSecretString := constant.SecretJWT
	hmacSecret := []byte(hmacSecretString)
	token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
		// check token signing method etc
		return hmacSecret, nil
	})

	if err != nil {
		return nil, false
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		mapClaims := token.Claims.(jwt.MapClaims)
		fmt.Printf("%+v \n", mapClaims["email"].(string))
		return claims, true
	} else {
		fmt.Printf("Invalid JWT Token")
		return nil, false
	}
}

func GetParams(jwtClaims jwt.MapClaims) ParamsToken {
	var params ParamsToken
	if value, ok := jwtClaims["email"]; ok {
		params.Email = value.(string)
	}
	if value, ok := jwtClaims["lang"]; ok {
		params.Lang = value.(string)
	}
	return params
}
