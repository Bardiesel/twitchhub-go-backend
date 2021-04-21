package models

import (
	"github.com/bardiesel/twitchhub-backend-go/utils"

	"gorm.io/gorm"
)

type Emote struct {
	gorm.Model
	Name        string
	Creator     string
	EmoteUrl    string
	Description string
}

func (e *Emote) GetAll() []Emote {
	var emotes []Emote
	db := utils.DatabaseConnection()
	db.Find(&emotes)
	return emotes
}

func (e *Emote) Create() (int64, error) {
	db := utils.DatabaseConnection()
	res := db.Create(&e)
	return res.RowsAffected, res.Error
}
