package models

import (
	"github.com/bardiesel/twitchhub-backend-go/utils"

	"gorm.io/gorm"
)

type Command struct {
	gorm.Model
	Title       string
	Description string
	User        *string
	Moderator   *string
	Bot         *string
}

func (c *Command) GetAll() []Command {
	var commands []Command
	db := utils.DatabaseConnection()
	db.Find(&commands)
	return commands
}

func (c *Command) Create() (int64, error) {
	db := utils.DatabaseConnection()
	res := db.Create(&c)
	return res.RowsAffected, res.Error
}
