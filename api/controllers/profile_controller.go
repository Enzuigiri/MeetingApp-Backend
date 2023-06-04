package controllers

import (
	"backend/app"
	"backend/domain"

	"github.com/gofiber/fiber/v2"
)

type ProfileController struct {
	ProfileUsecase domain.ProfileUsecase
	Env            *app.Env
}

func (pc *ProfileController) Profile(c *fiber.Ctx) error {
	userId := c.Locals("user").(domain.JWTUserData)

	if userId.ID == "" {
		return fiber.NewError(fiber.StatusInternalServerError, "Missing temporary variable")
	}

	profile, err := pc.ProfileUsecase.GetProfileByID(c.Context(), userId.ID)
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}

	return c.Status(fiber.StatusOK).JSON(profile)
}
