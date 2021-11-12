package route

import (
	"restful-api/controller"

	"github.com/labstack/echo/v4"
)

func New() *echo.Echo {
	e := echo.New()
	e.POST("/users", controller.CreateUserController)
	e.GET("/users", controller.GetUsersController)
	return e
}
