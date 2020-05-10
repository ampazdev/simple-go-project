package app

import (
	"github.com/ampazdev/simple-go-project/svc/productservice/internal/app/config"
)

type ProductService struct {
	Config   *config.Config
	Database *Database
	UseCases *UseCases
	Repos    *Repos
}

func NewProductService(filepath string) (*ProductService, error) {
	// read config from file
	cfg, err := config.NewConfig(filepath)
	if err != nil {
		return nil, err
	}

	db, err := newDatabase(&cfg.DB)
	if err != nil {
		return nil, err
	}

	repos := newRepos(db)
	usecases := newUseCases(repos)

	return &ProductService{
		Config:   cfg,
		Database: db,
		UseCases: usecases,
		Repos:    repos,
	}, nil
}

func (p *ProductService) Close() []error {
	var errs []error

	errs = append(errs, p.Database.Close())

	return errs
}
