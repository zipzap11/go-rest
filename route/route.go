package route

import (
	"restful-api/controller"
	md "restful-api/middleware"

	"github.com/labstack/echo/v4"
)

func New() *echo.Echo {
	e := echo.New()
	md.HTTPSMiddleware(e)
	md.LogMiddleware(e)
	e.POST("/users", controller.CreateUserController)
	e.GET("/users", controller.GetUsersController)
	return e
}
