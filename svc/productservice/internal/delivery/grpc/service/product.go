package service

import "github.com/ampazdev/simple-go-project/svc/productservice/internal"

type Product struct {
	UC internal.ProductUseCase
}

func NewProduct(UC internal.ProductUseCase) *Product {
	return &Product{
		UC: UC,
	}
}
