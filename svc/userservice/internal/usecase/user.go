package usecase

import (
	"context"
	"fmt"
	"github.com/ampazdev/simple-go-project/svc/userservice/internal/entity"
	"golang.org/x/crypto/bcrypt"
)

func (p *UserUC) GetByUserInfoByEmail(ctx context.Context, user entity.User) (*entity.User, error) {
	// first get from cache
	res, err := p.userReaderRepo.GetUserDetailByEmailCache(ctx, user)
	if err != nil {
		return nil, err
	}
	// immediate return when cache not nil
	if nil != res && res.Email != "" {
		return res, nil
	}
	// if not nil then get from db
	res, err = p.userReaderRepo.GetByUserInfoByEmail(ctx, user)
	if err != nil {
		return nil, err
	}
	// set cache so detail get from cache not from db
	if nil != res && res.Email != "" {
		err = p.userReaderRepo.SetUserDetailByEmailCache(ctx, *res)
		if err != nil {
			return nil, err
		}
	}
	return res, nil
}

func (p *UserUC) InsertUserInfo(ctx context.Context, user entity.User) (*entity.User, error) {
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
	hash, err := bcrypt.GenerateFromPassword([]byte(user.Password), 10)
	if nil != err {
		fmt.Println(err)
	}
	user.Password = string(hash)
	res, err := p.userReaderRepo.InsertUserCred(ctx, user)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (p *UserUC) SignupUser(ctx context.Context, user entity.User) (*entity.User, error) {
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

	if user.FullName == "" {
		//respond error Bad request
		errorObject.Message = "Full Name is missing"
		//utils.RespondWithError(w, http.StatusBadRequest, error)
		return &responseUser, err
	}

	if user.PhoneNumber == "" {
		//respond error Bad request
		errorObject.Message = "Phone Number is missing"
		//utils.RespondWithError(w, http.StatusBadRequest, error)
		return &responseUser, err
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(user.Password), 10)
	if nil != err {
		fmt.Println(err)
	}
	user.Password = string(hash)
	// insert user cred first
	res, err := p.userReaderRepo.InsertUserCred(ctx, user)
	if err != nil {
		return nil, err
	}
	user.ID = res.ID

	// insert user info
	res, err = p.userReaderRepo.InsertUserInfo(ctx, user)
	if err != nil {
		return nil, err
	}
	res.ID = user.ID
	res.FullName = user.FullName
	res.Email = user.Email
	res.PhoneNumber = user.PhoneNumber

	return res, nil
}

func (p *UserUC) SetUserDetailByEmailCache(ctx context.Context, user entity.User) error {
	fmt.Println("set user cache usecase")
	return nil
}

func (p *UserUC) GetUserDetailByEmailCache(ctx context.Context, user entity.User) (*entity.User, error) {
	fmt.Println("get user cache usecase")
	return &entity.User{}, nil
}
