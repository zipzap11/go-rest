package main

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type article struct {
	ID      int    `json:"id"`
	Title   string `json:"title"`
	Content string `json:"content"`
}

var data = []article{
	{1, "lorem", "lorem"},
	{2, "ipsum", "ipsum"},
}

func main() {
	e := echo.New()
	e.GET("/", HelloController)
	e.GET("/articles", GetArticlesController)
	e.GET("/articles/:id", GetArticlesByIdController)
	e.POST("/articles", AddArticle)
	e.Start(":8000")
}

func GetArticlesByIdController(e echo.Context) error {
	id, _ := strconv.Atoi(e.Param("id"))
	for i := 0; i < len(data); i++ {

		if id == data[i].ID {
			return e.JSON(http.StatusOK, data[i])
		}
	}
	return e.JSON(http.StatusBadRequest, "Not Found")
}
func GetArticlesController(e echo.Context) error {
	return e.JSON(http.StatusOK, data)
}

func HelloController(e echo.Context) error {
	return e.String(http.StatusOK, "Hello World")
}

func AddArticle(e echo.Context) error {
	newArticle := article{}
	e.Bind(&newArticle)
	data = append(data, newArticle)
	return e.JSON(http.StatusOK, "Sucess add article")
}
