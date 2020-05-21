package app

import (
	"github.com/ampazdev/simple-go-project/svc/productservice/internal/app/config"
)

type ProductService struct {
	Config   *config.Config
	Bridges  *Bridges
	UseCases *UseCases
	Repos    *Repos
}

func NewProductService(filepath string) (*ProductService, error) {
	// read config from file
	cfg, err := config.NewConfig(filepath)
	if err != nil {
		return nil, err
	}

	bridges, err := newBridges(&cfg.DB)
	if err != nil {
		return nil, err
	}

	repos := newRepos(bridges)
	usecases := newUseCases(repos)

	return &ProductService{
		Config:   cfg,
		Bridges:  bridges,
		UseCases: usecases,
		Repos:    repos,
	}, nil
}

func (p *ProductService) Close() []error {
	var errs []error

	errs = append(errs, p.Bridges.Close()...)

	return errs
}
