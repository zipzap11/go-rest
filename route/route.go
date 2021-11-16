package route

import (
	"restful-api/controller"
	md "restful-api/middleware"

	"github.com/labstack/echo/v4"
	mid "github.com/labstack/echo/v4/middleware"
)

func New() *echo.Echo {
	e := echo.New()
	// md.HTTPSMiddleware(e)
	md.LogMiddleware(e)
	e.POST("/users", controller.CreateUserController)
	e.GET("/users", controller.GetUsersController)
	e.GET("/companies", controller.GetCompaniesController)
	e.POST("/companies", controller.InsertCompanyController)
	e.GET("/users/:userid", controller.GetSpecificUserController)
	e.GET("/companies/:company_id/users", controller.GetEmployeeFromCompanyController)
	eAuthBasic := e.Group("/auth")
	eAuthBasic.Use(mid.BasicAuth(md.BasicAuthDB))
	eAuthBasic.GET("/users", controller.GetUsersController)
	return e
}
