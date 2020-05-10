package service

import (
	"net/http"

	"github.com/ampazdev/simple-go-project/svc/productservice/internal/entity"
	"github.com/labstack/echo"
)

type Product struct{}

func NewProduct() *Product {
	return &Product{}
}

func (p *Product) GetProducts(c echo.Context) error {
	return c.JSON(http.StatusOK, entity.Product{
		Name:        "Test",
		Description: "test desc",
	})
}
