package app

import (
	"github.com/ampazdev/simple-go-project/svc/userservice/internal"
	dbRepo "github.com/ampazdev/simple-go-project/svc/userservice/internal/repo/database"
)

type Repos struct {
	UserReaderRepo internal.UserReaderRepo
}

func newRepos(database *Database) *Repos {
	return &Repos{
		UserReaderRepo: dbRepo.NewUserReaderRepo(database.DB),
	}
}
