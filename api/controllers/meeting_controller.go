package controllers

import (
	"backend/app"
	"backend/domain"
	"fmt"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type MeetingController struct {
	MeetingUsecase domain.MeetingUsecase
	Env            *app.Env
	Validator      validator.Validate
}

func (mc *MeetingController) Create(c *fiber.Ctx) error {
	user := c.Locals("user").(domain.JWTUserData)

	var request domain.Meeting

	err := c.BodyParser(&request)
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	userIdHex, err := primitive.ObjectIDFromHex(user.ID)
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}

	request.PICID.UserID = userIdHex
	request.PICID.FirstName = user.FirstName
	request.Participants = []domain.Participant{}

	for i := range request.Agenda {
		request.Agenda[i].ID = primitive.NewObjectID()
		request.Agenda[i].ProposerID = userIdHex
		request.Agenda[i].FirstName = user.FirstName
		request.Agenda[i].Voters = []domain.Voter{}
		request.Agenda[i].CreatedAt = time.Now()
	}

	err = mc.Validator.Struct(request)
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "Missing tag or value that required")
	}

	request.ID = primitive.NewObjectID()
	request.CreatedAt = time.Now()

	err = mc.MeetingUsecase.Create(c.Context(), &request)
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}

	return c.Status(fiber.StatusOK).JSON(request)
}

func (mc *MeetingController) FetchByID(c *fiber.Ctx) error {
	user := c.Locals("user").(domain.JWTUserData)

	var request domain.FetchMeetingRequest

	err := c.BodyParser(&request)
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	err = mc.Validator.Struct(request)
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "Missing tag or value that required")
	}

	meeting, _, err := mc.MeetingUsecase.FetchByID(c.Context(), user.ID, request.ID)
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}

	return c.Status(fiber.StatusOK).JSON(meeting)
}

func (mc *MeetingController) FetchByUserID(c *fiber.Ctx) error {
	user := c.Locals("user").(domain.JWTUserData)

	meetings, err := mc.MeetingUsecase.FetchByUserID(c.Context(), user.ID)
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}
	var response []domain.FetchAllUserMeetingResponse

	for _, meeting := range meetings {
		response = append(response, domain.FetchAllUserMeetingResponse{
			ID:           meeting.ID.Hex(),
			Title:        meeting.Title,
			Location:     meeting.Location,
			PICID:        meeting.PICID,
			Schedule:     meeting.Schedule,
			VoteTime:     meeting.VoteTime,
			Participants: meeting.Participants,
		})
	}
	return c.Status(fiber.StatusOK).JSON(response)
}

func (mc *MeetingController) Update(c *fiber.Ctx) error {
	user := c.Locals("user").(domain.JWTUserData)

	var request domain.UpdateMeetingRequest
	err := c.BodyParser(&request)
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	err = mc.Validator.Struct(request)
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	meeting, isPIC, err := mc.MeetingUsecase.FetchByID(c.Context(), user.ID, request.ID)
	if err != nil || !isPIC {
		if !isPIC {
			err = fmt.Errorf("Not Authorized")
		}
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	meeting.Title = request.Title
	meeting.Desription = request.Description
	meeting.Location = request.Location
	meeting.Schedule = request.Schedule
	meeting.VoteTime = request.VoteTime

	response, err := mc.MeetingUsecase.Update(c.Context(), &meeting)
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}

	return c.Status(fiber.StatusOK).JSON(response)
}

func (mc *MeetingController) Delete(c *fiber.Ctx) error {
	user := c.Locals("user").(domain.JWTUserData)

	var request domain.FetchMeetingRequest
	err := c.BodyParser(&request)
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	err = mc.Validator.Struct(request)
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	err = mc.MeetingUsecase.Delete(c.Context(), user.ID, request.ID)
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	return c.SendStatus(fiber.StatusOK)
}
