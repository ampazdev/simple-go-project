package database

import (
	"context"
	"database/sql"

	"github.com/ampazdev/simple-go-project/svc/productservice/internal"
	"github.com/ampazdev/simple-go-project/svc/productservice/internal/entity"
)

type ProductReaderRepo struct {
	DB *sql.DB // TODO: create bridges
}

func NewProductReaderRepo(db *sql.DB) internal.ProductReaderRepo {
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

	stmt, err := p.DB.PrepareContext(ctx, q)
	if err != nil {
		return nil, err
	}

	row := stmt.QueryRowContext(ctx, &id)

	res := entity.Product{}
	err = row.Scan(
		&res.Name,
		&res.Description,
	)

	if err != nil {
		return nil, err
	}

	return &res, nil
}
