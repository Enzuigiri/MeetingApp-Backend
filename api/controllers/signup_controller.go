package controllers

import (
	"backend/app"
	"backend/domain"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type SignupController struct {
	SignupUsecase domain.SignupUsecase
	Env           *app.Env
	Validator     validator.Validate
}

func (sc *SignupController) Signup(c *fiber.Ctx) error {
	var request domain.SignupRequest

	err := c.BodyParser(&request)
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	err = sc.Validator.Struct(request)
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "Missing tag or value that required")
	}

	user, err := sc.SignupUsecase.CodeValidation(c.Context(), sc.Env.ClientID, sc.Env.ClientSecret, request.ValidateCode)
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	user.ID = primitive.NewObjectID()
	user.FirstName = request.FirstName
	user.LastName = request.LastName
	user.LastLogin = time.Now()

	userDataByEmail, errCheck := sc.SignupUsecase.GetByAppleID(c.Context(), user.AppleID)

	// Fix This No LOgic Shouldn't Perform in This Layer
	if userDataByEmail.Email != "" {
		user = userDataByEmail
	}

	refreshToken, err := sc.SignupUsecase.CreateRefreshToken(&user, sc.Env.RefreshTokenSecret, sc.Env.RefreshTokenExpiryHour)
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}

	if errCheck != nil {
		err = nil

		user.RefreshToken = refreshToken
		user.CreatedAt = time.Now()

		sc.SignupUsecase.Create(c.Context(), &user)
	}

	accessToken, err := sc.SignupUsecase.CreateAccessToken(&user, sc.Env.AccessTokenSecret, sc.Env.AccessTokenExpiryHour)
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}

	signupResponse := domain.SignupResponse{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}

	return c.Status(fiber.StatusOK).JSON(signupResponse)
}
