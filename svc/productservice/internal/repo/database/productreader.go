package database

import (
	"context"

	"github.com/ampazdev/simple-go-project/gopkg/bridge"
	"github.com/ampazdev/simple-go-project/svc/productservice/internal"
	"github.com/ampazdev/simple-go-project/svc/productservice/internal/entity"
)

type ProductReaderRepo struct {
	DB bridge.Database
}

func NewProductReaderRepo(db bridge.Database) internal.ProductReaderRepo {
	return &ProductReaderRepo{
		DB: db,
	}
}

func (p *ProductReaderRepo) GetByID(ctx context.Context, id int64) (*entity.Product, error) {
	q := `
			select 
				id,
				description
			from products
			where id = $1;
	`

	product := &entity.Product{}

	if err := p.DB.GetContext(ctx, product, q, &id); err != nil {
		return nil, err
	}

	return product, nil
}
