package internal

import (
	"context"

	"github.com/ampazdev/simple-go-project/svc/userservice/internal/entity"
)

type UserReaderRepo interface {
	GetByUserInfoByEmail(ctx context.Context, user entity.User) (*entity.User, error)
	GetUserCredByEmail(ctx context.Context, user entity.User) (*entity.User, error)
	InsertUserCred(ctx context.Context, user entity.User) (*entity.User, error)
	InsertUserInfo(ctx context.Context, user entity.User) (*entity.User, error)
	SignupUser(ctx context.Context, user entity.User) (*entity.User, error)
	Login(ctx context.Context, user entity.User) (*entity.User, error)

	SetUserDetailByEmailCache(ctx context.Context, user entity.User) error
	GetUserDetailByEmailCache(ctx context.Context, user entity.User) (*entity.User, error)
}

type UserWriterRepo interface{}
