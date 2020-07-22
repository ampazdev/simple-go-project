package usecase

import (
	"context"

	"github.com/ampazdev/simple-go-project/svc/productservice/internal"
	"github.com/ampazdev/simple-go-project/svc/productservice/internal/entity"
)

type ProductUC struct {
	productReaderRepo internal.ProductReaderRepo
}

func NewProductUC(productReaderRepo internal.ProductReaderRepo) internal.ProductUseCase {
	return &ProductUC{productReaderRepo: productReaderRepo}
}

func (p *ProductUC) GetByID(ctx context.Context, id int64) (*entity.Product, error) {
	res, err := p.productReaderRepo.GetByID(ctx, id)
	if err != nil {
		return nil, err
	}

	return res, nil
}
