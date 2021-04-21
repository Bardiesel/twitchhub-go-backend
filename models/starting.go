package models

import (
	"github.com/bardiesel/twitchhub-backend-go/utils"

	"gorm.io/gorm"
)

type Starting struct {
	gorm.Model
	Description string
}

func (s *Starting) GetAll() []Starting {
	var startings []Starting
	db := utils.DatabaseConnection()
	db.Find(&startings)
	return startings
}

func (s *Starting) Create() (int64, error) {
	db := utils.DatabaseConnection()
	res := db.Create(&s)
	return res.RowsAffected, res.Error
}
