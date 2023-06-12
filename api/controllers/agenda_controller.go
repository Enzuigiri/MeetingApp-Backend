package controllers

import (
	"backend/app"
	"backend/domain"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type AgendaController struct {
	AgendaUsecase  domain.AgendaUsecase
	MeetingUsecase domain.MeetingUsecase
	Env            *app.Env
	Validator      validator.Validate
}

func (ac *AgendaController) Add(c *fiber.Ctx) error {
	user := c.Locals("user").(domain.JWTUserData)

	var requests domain.AgendaRequests

	err := c.BodyParser(&requests)
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	meeting, _, err := ac.MeetingUsecase.FetchByID(c.Context(), user.ID, requests.MeetingId)
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	response, err := ac.AgendaUsecase.Add(c.Context(), &meeting, &requests.Agendas, user.ID, user.FirstName)
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}

	return c.Status(fiber.StatusOK).JSON(response)
}

// func (ac *AgendaController) Edit(c *fiber.Ctx) error {
// 	user := c.Locals("user").(domain.JWTUserData)

// 	var request domain.AgendaRequest
// 	var agenda *domain.Agenda

// 	err := c.BodyParser(&request)
// 	if err != nil {
// 		return fiber.NewError(fiber.StatusBadRequest, err.Error())
// 	}

// 	agendaIdHex, err := primitive.ObjectIDFromHex(request.AgendaId)
// 	if err != nil {
// 		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
// 	}

// 	meeting, _, err := ac.MeetingUsecase.FetchByID(c.Context(), user.ID, request.MeetingId)
// 	if err != nil {
// 		return fiber.NewError(fiber.StatusBadRequest, err.Error())
// 	}

// 	agenda.ID = agendaIdHex
// 	agenda.Title = request.Title
// 	agenda.Desription = request.Desription

// 	response, err := ac.AgendaUsecase.Edit(c.Context(), &meeting, agenda, user.ID)
// 	if err != nil {
// 		return fiber.NewError(fiber.StatusBadRequest, err.Error())
// 	}

// 	return c.Status(fiber.StatusOK).JSON(response)
// }

func (ac *AgendaController) Delete(c *fiber.Ctx) error {
	user := c.Locals("user").(domain.JWTUserData)

	var request domain.DeleteAgendaRequest

	err := c.BodyParser(&request)
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	meeting, _, err := ac.MeetingUsecase.FetchByID(c.Context(), user.ID, request.MeetingId)
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	response, err := ac.AgendaUsecase.Delete(c.Context(), &meeting, request.AgendaId, user.ID)
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	return c.Status(fiber.StatusOK).JSON(response)
}

func (ac *AgendaController) Vote(c *fiber.Ctx) error {
	user := c.Locals("user").(domain.JWTUserData)

	var request domain.AgendaVote

	err := c.BodyParser(&request)
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	err = ac.Validator.Struct(request)
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "Missing tag or value that required")
	}

	meeting, _, err := ac.MeetingUsecase.FetchByID(c.Context(), user.ID, request.MeetingId)
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	voterId, err := primitive.ObjectIDFromHex(user.ID)
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	var voter *domain.Voter

	voter.UserId = voterId
	voter.FirstName = user.FirstName

	response, err := ac.AgendaUsecase.Vote(c.Context(), &meeting, request.VoteValue, voter)
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}

	return c.Status(fiber.StatusOK).JSON(response)
}
