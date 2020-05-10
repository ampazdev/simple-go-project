package service

import (
	"context"
	"net/http"
	"strconv"

	"github.com/ampazdev/simple-go-project/svc/productservice/internal"
	"github.com/ampazdev/simple-go-project/svc/productservice/internal/entity"
	"github.com/labstack/echo/v4"
)

type Product struct {
	UC internal.ProductUseCase
}

func NewProduct(uc internal.ProductUseCase) *Product {
	return &Product{
		UC: uc,
	}
}

func (p *Product) GetProducts(c echo.Context) error {
	return c.JSON(http.StatusOK, entity.Product{
		Name:        "Test",
		Description: "test desc",
	})
}

func (p *Product) GetProduct(c echo.Context) error {
	paramID := c.Param("id")

	id, err := strconv.Atoi(paramID)
	if err != nil {
		return err
	}

	res, err := p.UC.GetByID(context.TODO(), int64(id))
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, res)
}
