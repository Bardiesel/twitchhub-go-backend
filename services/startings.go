package services

import (
	"github.com/bardiesel/twitchhub-backend-go/models"

	validation "github.com/go-ozzo/ozzo-validation"
)

type StartingResponse struct {
	Message string          `json:"message"`
	Data    models.Starting `json:"data"`
}

type StartingPayload struct {
	Description string
}

func (s StartingPayload) Validate() error {
	return validation.ValidateStruct(&s,
		validation.Field(&s.Description, validation.Required, validation.Length(5, 50)),
	)
}
