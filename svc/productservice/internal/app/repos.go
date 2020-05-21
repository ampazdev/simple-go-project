package app

import (
	"github.com/ampazdev/simple-go-project/svc/productservice/internal"
	dbRepo "github.com/ampazdev/simple-go-project/svc/productservice/internal/repo/database"
)

type Repos struct {
	ProductReaderRepo internal.ProductReaderRepo
}

func newRepos(bridges *Bridges) *Repos {
	return &Repos{
		ProductReaderRepo: dbRepo.NewProductReaderRepo(bridges.DB),
	}
}
