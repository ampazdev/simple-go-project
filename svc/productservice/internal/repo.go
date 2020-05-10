package internal

import (
	"context"

	"github.com/ampazdev/simple-go-project/svc/productservice/internal/entity"
)

type ProductReaderRepo interface {
	GetByID(ctx context.Context, id int64) (*entity.Product, error)
}

type ProductWriterRepo interface{}
