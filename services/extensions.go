package services

import (
	"github.com/bardiesel/twitchhub-backend-go/models"

	validation "github.com/go-ozzo/ozzo-validation"
)

type ExtensionResponse struct {
	Message string           `json:"message"`
	Data    models.Extension `json:"data"`
}

type ExtensionPayload struct {
	Title string
}

func (ex ExtensionPayload) Validate() error {
	return validation.ValidateStruct(&ex,
		validation.Field(&ex.Title, validation.Required),
	)
}
