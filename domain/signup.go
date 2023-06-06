package domain

import (
	"context"
)

type SignupRequest struct {
	ValidateCode string `json:"validateCode" validate:"required,min=10"`
	FirstName    string `json:"firstName" validate:"required"`
	LastName     string `json:"lastName,omitempty"`
}

type SignupResponse struct {
	AccessToken  string `json:"accessToken"`
	RefreshToken string `json:"refreshToken"`
}

type SignupUsecase interface {
	Create(c context.Context, user *User) error
	GetByAppleID(c context.Context, id string) (User, error)
	GetByEmail(c context.Context, email string) (User, error)
	CodeValidation(c context.Context, clientId string, clientSecret string, code string) (User, error)
	CreateAccessToken(user *User, secret string, expiry int) (accessToken string, err error)
	CreateRefreshToken(user *User, secret string, expiry int) (refreshToken string, err error)
}
