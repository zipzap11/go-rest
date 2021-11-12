package main

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type User struct {
	ID    int    `json: "id" form: "id"`
	Name  string `json: "name" form: "name"`
	Email string `json: "email" form: "email"`
}

var Users = []User{
	User{1, "Abadi", "abadi@gmail.com"},
	User{2, "Suryo", "suryo@gmail.com"},
}

var DB *gorm.DB

func initDB() {
	dsn := "root:@tcp(127.0.0.1:3306)/training?charset=utf8mb4&parseTime=True&loc=Local"
	var err error
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err.Error())
	}
	initMigrate()
}

func initMigrate() {
	DB.AutoMigrate(&User{})
}

func main() {
	initDB()
	e := echo.New()
	e.POST("/users", CreateUserController)
	e.GET("/users", GetUsersController)
	e.Start("127.0.0.1:8000")
}

func CreateUserController(e echo.Context) error {
	user := User{}
	e.Bind(&user)
	Users = append(Users, user)
	return e.JSON(http.StatusOK, map[string]interface{}{
		"message": "success create user",
		"user":    user,
	})
}

func GetUsersController(e echo.Context) error {
	var users []User

	err := DB.Find(&users).Error
	if err != nil {
		return e.JSON(http.StatusInternalServerError, map[string]interface{}{
			"messages": err.Error(),
		})
	}
	return e.JSON(http.StatusOK, map[string]interface{}{
		"messages": "Success",
		"users":    users,
	})
}
