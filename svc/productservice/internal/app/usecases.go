package app

import (
	"github.com/ampazdev/simple-go-project/svc/productservice/internal"
	"github.com/ampazdev/simple-go-project/svc/productservice/internal/usecase"
)

type UseCases struct {
	ProductUC internal.ProductUseCase
}

func newUseCases(repos *Repos) *UseCases {
	return &UseCases{
		ProductUC: usecase.NewProductUC(repos.ProductReaderRepo),
	}
}

func (*UseCases) Close() []error {
	var errs []error
	return errs
}
