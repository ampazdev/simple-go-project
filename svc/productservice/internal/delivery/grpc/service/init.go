package service

import (
	"context"

	"github.com/ampazdev/simple-go-project/svc/productservice/internal/app"
	productproto "github.com/ampazdev/simple-go-project/myproto/productservice"
)

// List all GRPC services
type Services struct {
	*Product
}

// TODO: Connect to uc layer
func (s Services) GetProducts(ctx context.Context, req *productproto.GetProductsRequest) (*productproto.GetProductsResponse, error) {
	return &productproto.GetProductsResponse{
		Products:             []*productproto.Product{
			{
				Id:                   1,
				Name:                 "Prod A",
				Description:          "Desc prod a",
			},
		},
	}, nil
}

func GetServices(app *app.ProductService) *Services {
	return &Services{
		Product: NewProduct(app.UseCases.ProductUC),
	}
}
