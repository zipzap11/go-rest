package controller

import (
	"fmt"
	"net/http"
	"restful-api/config"
	"restful-api/model"
	"strconv"

	"github.com/labstack/echo/v4"
)

func CreateUserController(e echo.Context) error {
	user := model.User{}
	e.Bind(&user)
	// Users = append(Users, user)
	fmt.Println(user)
	err := config.DB.Save(&user).Error

	if err != nil {
		return e.JSON(http.StatusInternalServerError, map[string]interface{}{
			"messages": err.Error(),
		})
	}

	return e.JSON(http.StatusOK, map[string]interface{}{
		"message": "success create user",
		"user":    user,
	})
}

func GetUsersController(e echo.Context) error {
	var users []model.User

	err := config.DB.Find(&users).Error
	if err != nil {
		return e.JSON(http.StatusInternalServerError, map[string]interface{}{
			"messages": err.Error(),
		})
	}
	return e.JSON(http.StatusOK, map[string]interface{}{
		"messages": "Success read users",
		"users":    users,
	})
}

func GetSpecificUserController(e echo.Context) error {
	id, _ := strconv.Atoi(e.Param("userid"))
	fmt.Println("ID =======", id)
	var user model.User

	err := config.DB.First(&user, id).Error
	if err != nil {
		return e.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": err.Error(),
		})
	}
	return e.JSON(http.StatusOK, map[string]interface{}{
		"message": "success get specific user",
		"user":    user,
	})
}
