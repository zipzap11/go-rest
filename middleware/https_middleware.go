package middleware

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func HTTPSMiddleware(e *echo.Echo) {
	e.Use(middleware.HTTPSRedirect())
}
