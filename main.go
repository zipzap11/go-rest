package main

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

type article struct {
	ID int
	Title string
	Content string
}

var data = []article{
	{1, "lorem", "lorem"},
	{2, "ipsum", "ipsum"},
}

func main() {
	e := echo.New()
	e.GET("/", HelloController)
	e.GET("/articles", GetArticlesController)
	e.Start(":8000")
}

func GetArticlesController(e echo.Context) error {
	return e.JSON(http.StatusOK, data)
}

func HelloController(e echo.Context) error {
	return e.String(http.StatusOK, "Hello World")
}