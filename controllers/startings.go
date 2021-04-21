package controllers

import (
	"github.com/bardiesel/twitchhub-backend-go/models"
	"github.com/bardiesel/twitchhub-backend-go/services"

	"github.com/gofiber/fiber/v2"
)

type StartingController struct {
	model models.Starting
}

func (s StartingController) Index(ctx *fiber.Ctx) error {
	startings := s.model.GetAll()
	return ctx.JSON(startings)
}

func (s StartingController) Create(ctx *fiber.Ctx) error {
	b := new(services.StartingPayload)

	if err := ctx.BodyParser(b); err != nil {
		return err
	}

	validationErrors := b.Validate()

	if validationErrors != nil {
		return ctx.JSON(validationErrors)
	}

	s.model = models.Starting{
		Description: b.Description,
	}

	_, modelError := s.model.Create()

	if modelError != nil {
		fiber.NewError(fiber.StatusConflict, modelError.Error())
	}

	return ctx.JSON(services.StartingResponse{
		Message: "success",
		Data:    s.model,
	})

}
