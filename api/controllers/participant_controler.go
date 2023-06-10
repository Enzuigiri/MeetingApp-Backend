package controllers

import (
	"backend/app"
	"backend/domain"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type ParticipantController struct {
	MeetingUsecase    domain.MeetingUsecase
	ParticpantUsecase domain.ParticipantUsecase
	Env               *app.Env
	Validattor        validator.Validate
}

func (pc *ParticipantController) Join(c *fiber.Ctx) error {
	user := c.Locals("user").(domain.JWTUserData)

	var request domain.JoinRequest

	err := c.BodyParser(&request)
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	userIdHex, err := primitive.ObjectIDFromHex(user.ID)
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}

	meeting, err := pc.ParticpantUsecase.CheckMeetingByCode(c.Context(), request.MeetingCode)
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	participant := domain.Participant{
		UserID:    userIdHex,
		FirstName: user.FirstName,
		JoinTime:  time.Now(),
	}

	response, err := pc.ParticpantUsecase.Add(c.Context(), &participant, &meeting)
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}

	return c.Status(fiber.StatusOK).JSON(response)
}

func (pc *ParticipantController) Remove(c *fiber.Ctx) error {
	user := c.Locals("user").(domain.JWTUserData)

	var request domain.RemoveRequest

	err := c.BodyParser(&request)
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	meeting, isPIC, err := pc.MeetingUsecase.FetchByID(c.Context(), user.ID, request.MeetingId)
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	if request.UserId != "" && !isPIC {
		return fiber.NewError(fiber.StatusBadRequest, "Not Authorized")
	}

	if request.UserId != "" && isPIC {
		response, err := pc.ParticpantUsecase.Delete(c.Context(), &meeting, request.UserId)
		if err != nil {
			return fiber.NewError(fiber.StatusBadRequest, err.Error())
		}
		return c.Status(fiber.StatusOK).JSON(response)
	}

	response, err := pc.ParticpantUsecase.Delete(c.Context(), &meeting, user.ID)
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}
	return c.Status(fiber.StatusOK).JSON(response)
}
