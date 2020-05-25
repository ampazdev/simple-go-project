package app

import (
	"github.com/ampazdev/simple-go-project/svc/userservice/internal/app/config"
)

type UserService struct {
	Config   *config.Config
	Database *Database
	UseCases *UseCases
	Repos    *Repos
}

func NewUserservice(filepath string) (*UserService, error) {
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

	return &UserService{
		Config:   cfg,
		Database: db,
		UseCases: usecases,
		Repos:    repos,
	}, nil
}

func (p *UserService) Close() []error {
	var errs []error

	errs = append(errs, p.Database.Close())

	return errs
}
