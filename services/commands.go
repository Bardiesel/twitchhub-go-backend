package services

import (
	"github.com/bardiesel/twitchhub-backend-go/models"

	validation "github.com/go-ozzo/ozzo-validation"
)

type CommandResponse struct {
	Message string         `json:"message"`
	Data    models.Command `json:"data"`
}

type CommandPayload struct {
	Title       string
	Description string
	User        *string
	Moderator   *string
	Bot         *string
}

func (c CommandPayload) Validate() error {
	return validation.ValidateStruct(&c,
		validation.Field(&c.Title, validation.Required, validation.Length(2, 20)),
		validation.Field(&c.Description, validation.Required, validation.Length(2, 50)),
	)
}
