package echo

import (
	"net/http"

	"github.com/isfanazha/simple-go-project/svc/productservice/internal/app"
	"github.com/isfanazha/simple-go-project/svc/productservice/internal/delivery/rest/service"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func Start(app *app.ProductService) {
	const op = "rest.Start"

	e := echo.New()
	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "time=${time_rfc3339}, method=${method}, uri=${uri}, status=${status} latency=${latency_human}\n",
	}))

	v1 := e.Group("/v1")

	prod := service.NewProduct()
	v1.Add(http.MethodGet, "/products", prod.GetProducts)

	e.Logger.Fatal(e.Start(":5000"))
}
