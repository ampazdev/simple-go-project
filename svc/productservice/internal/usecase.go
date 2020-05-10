package internal

import (
	"context"
	"github.com/isfanazha/simple-go-project/services/userservices/internal/entity"
)

//go:generate mockgen -destination=usecase/mock_usecase/user.go -package=mock_usecase github.com/isfanazha/simple-go-project/services/userservices/internal UserUseCase
type UserUseCase interface {
	GetByID(ctx context.Context, id int64) (*entity.User, error)
	GetAll(ctx context.Context) ([]*entity.User, error)
}
