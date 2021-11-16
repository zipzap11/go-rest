package middleware

import (
	"restful-api/config"
	"restful-api/model"

	"github.com/labstack/echo/v4"
)

func BasicAuthDB(username, password string, e echo.Context) (bool, error) {
	var user model.User

	err := config.DB.Where("email = ? and password = ?", username, password).First(&user).Error
	if err != nil {
		return false, err
	}

	return true, nil
}
