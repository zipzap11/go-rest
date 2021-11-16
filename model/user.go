package model

import "gorm.io/gorm"

type User struct {
	gorm.Model
	ID        int `json: "id" form: "id"`
	CompanyID int `json: "companyid" form: "companyid"`
	// Languages []Language `gorm:"many2many:user_languages;`
	Name     string `json: "name" form: "name"`
	Email    string `json: "email" form: "email"`
	Password string `json: "password" form: "password"`
}
