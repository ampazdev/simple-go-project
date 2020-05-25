package service

import (
	"context"
	"fmt"
	"github.com/ampazdev/simple-go-project/svc/userservice/utils"
	"github.com/dgrijalva/jwt-go"
	"net/http"
	"os"
	"strings"

	"github.com/ampazdev/simple-go-project/svc/userservice/internal"
	"github.com/ampazdev/simple-go-project/svc/userservice/internal/entity"
	"github.com/labstack/echo/v4"
)

type User struct {
	UC internal.UserUseCase
}

func NewUser(uc internal.UserUseCase) *User {
	return &User{
		UC: uc,
	}
}

func (l *User) SignupUser(c echo.Context) error {
	var err error
	user := new(entity.User)
	if err = c.Bind(user); err != nil {
		fmt.Println("error bind")
	}
	res, err := l.UC.SignupUser(context.TODO(), *user)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	return c.JSON(http.StatusOK, res)
}

func (l *User) Login(c echo.Context) error {
	var err error
	user := new(entity.User)
	if err = c.Bind(user); err != nil {
		fmt.Println("error bind")
	}
	res, err := l.UC.Login(context.TODO(), *user)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	return c.JSON(http.StatusOK, res)
}

//Auth middleware
func (l *User) Auth(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		authHeader := c.Request().Header.Get("Authorization")
		bearerToken := strings.Split(authHeader, " ")
		var errorObject entity.Error
		if len(bearerToken) == 2 {
			authToken := bearerToken[1]
			token, err := jwt.Parse(authToken, func(token *jwt.Token) (i interface{}, err error) {
				if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
					return nil, fmt.Errorf("There was an error")
				}
				return []byte(os.Getenv("SECRET")), nil
			})
			if nil != err {
				errorObject.Message = err.Error()
				return c.JSON(http.StatusBadRequest, errorObject)
			}
			if token.Valid {
				return next(c)

			} else {
				errorObject.Message = "Invalid Token 1"
				return c.JSON(http.StatusBadRequest, errorObject)

			}
		} else {
			errorObject.Message = "Invalid Token 2"
			return c.JSON(http.StatusBadRequest, errorObject)
		}

	}
}

func (l *User) GetUser(c echo.Context) error {
	var err error
	var errObject entity.Error
	user := new(entity.User)
	if err = c.Bind(user); err != nil {
		fmt.Println("error bind")
	}
	authHeader := c.Request().Header.Get("Authorization")
	bearerToken := strings.Split(authHeader, " ")
	if len(bearerToken) == 2 {
		authToken := bearerToken[1]
		mapClaims, valid := utils.ExtractClaims(authToken)
		if !valid {
			errObject.Message = "Token not valid!"
			return c.JSON(http.StatusBadRequest, errObject)
		}
		params := utils.GetParams(mapClaims)
		user.Email = params.Email
	}
	res, err := l.UC.GetByUserInfoByEmail(context.TODO(), *user)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	return c.JSON(http.StatusOK, res)
}
