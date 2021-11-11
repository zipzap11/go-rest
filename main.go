package main

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()
	e.GET("/", HelloController)
	e.Start(":8000")
}

func HelloController(e echo.Context) error {
	return e.String(http.StatusOK, "Hello World")
}