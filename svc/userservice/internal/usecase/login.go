package usecase

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/ampazdev/simple-go-project/svc/userservice/utils"

	"github.com/ampazdev/simple-go-project/svc/userservice/internal"
	"github.com/ampazdev/simple-go-project/svc/userservice/internal/entity"
	"golang.org/x/crypto/bcrypt"
)

type UserUC struct {
	userReaderRepo internal.UserReaderRepo
}

func NewUserUC(loginReaderRepo internal.UserReaderRepo) internal.UserReaderRepo {
	return &UserUC{userReaderRepo: loginReaderRepo}
}

func (p *UserUC) GetUserCredByEmail(ctx context.Context, user entity.User) (*entity.User, error) {
	res, err := p.userReaderRepo.GetByUserInfoByEmail(ctx, user)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (p *UserUC) InsertUserCred(ctx context.Context, user entity.User) (*entity.User, error) {
	res, err := p.userReaderRepo.GetByUserInfoByEmail(ctx, user)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (p *UserUC) Login(ctx context.Context, user entity.User) (*entity.User, error) {
	var errorObject entity.Error
	var responseUser entity.User
	var err error
	if user.Email == "" {
		//respond error Bad request
		errorObject.Message = "Email is missing"
		//utils.RespondWithError(w, http.StatusBadRequest, error)
		return &responseUser, err
	}
	if user.Password == "" {
		//respond error Bad Request
		errorObject.Message = "Password is missing"
		//utils.RespondWithError(w, http.StatusUnauthorized, error)
		return &responseUser, err
	}
	password := user.Password
	//get user password from db
	userCred, err := p.userReaderRepo.GetUserCredByEmail(ctx, user)
	if err != nil {
		if err == sql.ErrNoRows {
			fmt.Println("The User Doesnt Exist")
		}
		return nil, err
	}
	if nil == userCred {
		fmt.Println("nil user")
		return nil, err
	}
	hashedPassword := userCred.Password
	err = bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
	if nil != err {
		fmt.Println("Password doesnt match")
		return nil, err
	}
	token, err := utils.GenerateToken(*userCred)
	if nil != err {
		fmt.Println("Error Generate Token")
		return nil, err
	}
	//get user  from db
	userDB, err := p.userReaderRepo.GetByUserInfoByEmail(ctx, user)
	if err != nil {
		if err == sql.ErrNoRows {
			fmt.Println("The User Doesnt Exist")
		}
		return nil, err
	}
	jwt := entity.JWT{Token: token}
	res := userDB
	res.JWT = jwt
	return res, nil
}
