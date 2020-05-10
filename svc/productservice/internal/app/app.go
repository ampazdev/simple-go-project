package app

import (
	"github.com/isfanazha/simple-go-project/svc/productservice/internal/app/config"
)

type ProductService struct {
	Config config.Config
	//UseCases *UseCase
	//Repos    *repository
}

func NewProductService(filepath string) (*ProductService, error) {
	// read config from file
	cfg, err := config.NewConfig(filepath)
	if err != nil {
		return nil, err
	}

	app := &ProductService{
		Config: cfg,
		//UseCases: nil,
		//Repos:    nil,
	}

	return app, nil
}

func (a *ProductService) Close() []error {
	var errs []error

	return errs
}
