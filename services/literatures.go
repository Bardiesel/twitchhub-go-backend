package services

import (
	"github.com/bardiesel/twitchhub-backend-go/models"

	validation "github.com/go-ozzo/ozzo-validation"
)

type LiteratureResponse struct {
	Message string            `json:"message"`
	Data    models.Literature `json:"data"`
}

type LiteraturePayload struct {
	Title       string
	Description string
}

func (l LiteraturePayload) Validate() error {
	return validation.ValidateStruct(&l,
		validation.Field(&l.Title, validation.Required, validation.Length(5, 20)),
		validation.Field(&l.Description, validation.Required, validation.Length(5, 50)),
	)
}
