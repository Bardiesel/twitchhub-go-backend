package models

import (
	"github.com/bardiesel/twitchhub-backend-go/utils"

	"gorm.io/gorm"
)

type Extension struct {
	gorm.Model
	Title string
}

func (ex *Extension) GetAll() []Extension {
	var extensions []Extension
	db := utils.DatabaseConnection()
	db.Find(&extensions)
	return extensions
}

func (ex *Extension) Create() (int64, error) {
	db := utils.DatabaseConnection()
	res := db.Create(&ex)
	return res.RowsAffected, res.Error
}
