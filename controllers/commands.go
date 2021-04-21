package controllers

import (
	"github.com/bardiesel/twitchhub-backend-go/models"
	"github.com/bardiesel/twitchhub-backend-go/services"

	"github.com/gofiber/fiber/v2"
)

type CommandController struct {
	model models.Command
}

func (c CommandController) Index(ctx *fiber.Ctx) error {
	commands := c.model.GetAll()
	return ctx.JSON(commands)
}

func (c CommandController) Create(ctx *fiber.Ctx) error {
	b := new(services.CommandPayload)

	if err := ctx.BodyParser(b); err != nil {
		return err
	}

	validationErrors := b.Validate()

	if validationErrors != nil {
		return ctx.JSON(validationErrors)
	}

	c.model = models.Command{
		Title:       b.Title,
		Description: b.Description,
		User:        b.User,
		Moderator:   b.Moderator,
		Bot:         b.Bot,
	}

	_, modelError := c.model.Create()

	if modelError != nil {
		fiber.NewError(fiber.StatusConflict, modelError.Error())
	}

	return ctx.JSON(services.CommandResponse{
		Message: "success",
		Data:    c.model,
	})

}
