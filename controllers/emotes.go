package controllers

import (
	"github.com/bardiesel/twitchhub-backend-go/models"
	"github.com/bardiesel/twitchhub-backend-go/services"

	"github.com/gofiber/fiber/v2"
)

type EmoteController struct {
	model models.Emote
}

func (e EmoteController) Index(ctx *fiber.Ctx) error {
	emotes := e.model.GetAll()
	return ctx.JSON(emotes)
}

func (e EmoteController) Create(ctx *fiber.Ctx) error {
	b := new(services.EmotePayload)

	if err := ctx.BodyParser(b); err != nil {
		return err
	}

	validationErrors := b.Validate()

	if validationErrors != nil {
		return ctx.JSON(validationErrors)
	}

	e.model = models.Emote{
		Name:        b.Name,
		Creator:     b.Creator,
		EmoteUrl:    b.EmoteUrl,
		Description: b.Description,
	}

	_, modelError := e.model.Create()

	if modelError != nil {
		fiber.NewError(fiber.StatusConflict, modelError.Error())
	}

	return ctx.JSON(services.EmoteResponse{
		Message: "success",
		Data:    e.model,
	})

}
