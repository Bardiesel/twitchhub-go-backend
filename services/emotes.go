package services

import (
	"github.com/bardiesel/twitchhub-backend-go/models"

	validation "github.com/go-ozzo/ozzo-validation"
)

type EmoteResponse struct {
	Message string       `json:"message"`
	Data    models.Emote `json:"data"`
}

type EmotePayload struct {
	Name        string
	Creator     string
	EmoteUrl    string
	Description string
}

func (e EmotePayload) Validate() error {
	return validation.ValidateStruct(&e,
		validation.Field(&e.Name, validation.Required, validation.Length(5, 20)),
		validation.Field(&e.Description, validation.Required, validation.Length(5, 50)),
		validation.Field(&e.Creator, validation.Required, validation.Length(5, 20)),
		validation.Field(&e.EmoteUrl, validation.Required, validation.Length(5, 20)),
	)
}
