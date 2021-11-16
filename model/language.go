package model

import "gorm.io/gorm"

type Language struct {
	gorm.Model
	ID       int    `json: "id" form: "id"`
	Language string `json: "language" form: "language"`
}
