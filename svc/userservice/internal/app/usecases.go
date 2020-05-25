package app

import (
	"github.com/ampazdev/simple-go-project/svc/userservice/internal"
	"github.com/ampazdev/simple-go-project/svc/userservice/internal/usecase"
)

type UseCases struct {
	UserUseCase internal.UserUseCase
}

func newUseCases(repos *Repos) *UseCases {
	return &UseCases{
		UserUseCase: usecase.NewUserUC(repos.UserReaderRepo),
	}
}

func (*UseCases) Close() []error {
	var errs []error
	return errs
}
