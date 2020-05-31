package app

import (
	"github.com/ampazdev/simple-go-project/svc/userservice/internal"
	dbRepo "github.com/ampazdev/simple-go-project/svc/userservice/internal/repo/database"
	"github.com/gomodule/redigo/redis"
)

type Repos struct {
	UserReaderRepo internal.UserReaderRepo
}

func newRepos(database *Database, cache *redis.Pool) *Repos {
	return &Repos{
		UserReaderRepo: dbRepo.NewUserReaderRepo(database.DB, cache),
	}
}
