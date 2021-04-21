package controllers

import (
	"github.com/bardiesel/twitchhub-backend-go/models"
	"github.com/bardiesel/twitchhub-backend-go/services"

	"github.com/gofiber/fiber/v2"
)

type ExtensionController struct {
	model models.Extension
}

func (ex ExtensionController) Index(ctx *fiber.Ctx) error {
	extensions := ex.model.GetAll()
	return ctx.JSON(extensions)
}

func (ex ExtensionController) Create(ctx *fiber.Ctx) error {
	b := new(services.ExtensionPayload)

	if err := ctx.BodyParser(b); err != nil {
		return err
	}

	validationErrors := b.Validate()

	if validationErrors != nil {
		return ctx.JSON(validationErrors)
	}

	ex.model = models.Extension{
		Title: b.Title,
	}

	_, modelError := ex.model.Create()

	if modelError != nil {
		fiber.NewError(fiber.StatusConflict, modelError.Error())
	}

	return ctx.JSON(services.ExtensionResponse{
		Message: "success",
		Data:    ex.model,
	})

}
