package usecase

import (
	"context"
	"github.com/ampazdev/simple-go-project/services/userservices/internal"
	"github.com/ampazdev/simple-go-project/services/userservices/internal/entity"
	"time"
)

type userUseCase struct {
	userRepo       internal.UserRepo
	contextTimeout time.Duration
}

func (u userUseCase) GetAll(c context.Context) ([]*entity.User, error) {
	ctx, cancel := context.WithTimeout(c, u.contextTimeout)
	defer cancel()

	res, err := u.userRepo.GetAll(ctx)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (u userUseCase) GetByID(c context.Context, id int64) (*entity.User, error) {
	ctx, cancel := context.WithTimeout(c, u.contextTimeout)
	defer cancel()

	us, err := u.userRepo.GetByID(ctx, id)
	if err != nil {
		return nil, err
	}
	return us, nil
}

func NewUserUseCase(userRepo internal.UserRepo, contextTimeout time.Duration) *userUseCase {
	return &userUseCase{
		userRepo:       userRepo,
		contextTimeout: contextTimeout,
	}
}
