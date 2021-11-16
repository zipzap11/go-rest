package controller

import (
	"fmt"
	"net/http"
	"restful-api/config"
	"restful-api/model"

	"github.com/labstack/echo/v4"
)

func GetCompaniesController(e echo.Context) error {
	var companies []model.Company

	err := config.DB.Find(&companies).Error
	if err != nil {
		return e.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": err.Error(),
		})
	}
	return e.JSON(http.StatusOK, map[string]interface{}{
		"message":   "Success",
		"companies": companies,
	})
}

func InsertCompanyController(e echo.Context) error {
	var company model.Company

	e.Bind(&company)

	err := config.DB.Save(&company).Error
	if err != nil {
		return e.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": err.Error(),
		})
	}
	return e.JSON(http.StatusOK, map[string]interface{}{
		"message": "success",
		"company": company,
	})
}

func GetEmployeeFromCompanyController(e echo.Context) error {
	var users []model.User
	company_id := e.Param("company_id")
	fmt.Println("COMPANY ID ======= ", company_id)

	err := config.DB.Where("company_id = ?", company_id).Find(&users).Error
	if err != nil {
		return e.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": err.Error,
		})
	}

	return e.JSON(http.StatusOK, map[string]interface{}{
		"message":   "success",
		"employees": users,
	})
}
