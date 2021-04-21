package controllers

import (
	"github.com/bardiesel/twitchhub-backend-go/models"
	"github.com/bardiesel/twitchhub-backend-go/services"

	"github.com/gofiber/fiber/v2"
)

type LiteratureController struct {
	model models.Literature
}

func (l LiteratureController) Index(ctx *fiber.Ctx) error {
	literatures := l.model.GetAll()
	return ctx.JSON(literatures)
}

func (l LiteratureController) Create(ctx *fiber.Ctx) error {
	b := new(services.LiteraturePayload)

	if err := ctx.BodyParser(b); err != nil {
		return err
	}

	validationErrors := b.Validate()

	if validationErrors != nil {
		return ctx.JSON(validationErrors)
	}

	l.model = models.Literature{
		Title:       b.Title,
		Description: b.Description,
	}

	_, modelError := l.model.Create()

	if modelError != nil {
		fiber.NewError(fiber.StatusConflict, modelError.Error())
	}

	return ctx.JSON(services.LiteratureResponse{
		Message: "success",
		Data:    l.model,
	})

}
