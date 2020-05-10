package internal

import (
	"context"
	"github.com/ampazdev/simple-go-project/services/userservices/internal/entity"
)

//go:generate mockgen -destination=repo/mock_repo/user.go -package=mock_repo github.com/ampazdev/simple-go-project/services/userservices/internal UserRepo
type UserRepo interface {
	GetByID(ctx context.Context, id int64) (*entity.User, error)
	GetAll(ctx context.Context) ([]*entity.User, error)
}

