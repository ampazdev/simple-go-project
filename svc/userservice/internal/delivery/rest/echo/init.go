package echo

import (
	"net/http"

	"github.com/ampazdev/simple-go-project/svc/userservice/internal/app"
	"github.com/ampazdev/simple-go-project/svc/userservice/internal/delivery/rest/service"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func Start(app *app.UserService) {
	const op = "rest.Start"

	e := echo.New()
	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "time=${time_rfc3339}, method=${method}, uri=${uri}, status=${status} latency=${latency_human}\n",
	}))

	v1 := e.Group("/v1")

	user := service.NewUser(app.UseCases.UserUseCase)
	v1.Add(http.MethodPost, "/signup", user.SignupUser)
	v1.Add(http.MethodPost, "/login", user.Login)
	v1User := e.Group("/v1/user")
	v1User.Use(user.Auth)
	v1User.Add(http.MethodGet, "/getProfile", user.GetUser)

	e.Logger.Fatal(e.Start(":5005"))
}
