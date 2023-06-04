package usecase

import (
	"context"
	"fmt"
	"time"

	"backend/domain"
	"backend/internal/utils"

	"github.com/Timothylock/go-signin-with-apple/apple"
)

type signupUsecase struct {
	userRepository domain.UserRepository
	contextTimeout time.Duration
}

func NewsSignupUsecase(userRepository domain.UserRepository, timeout time.Duration) domain.SignupUsecase {
	return &signupUsecase{
		userRepository: userRepository,
		contextTimeout: timeout,
	}
}

func (su *signupUsecase) Create(c context.Context, user *domain.User) error {
	ctx, cancel := context.WithTimeout(c, su.contextTimeout)
	defer cancel()
	return su.userRepository.Create(ctx, user)
}

func (su *signupUsecase) CodeValidation(c context.Context, clientSecret string, code string) (domain.User, error) {
	ctx, cancel := context.WithTimeout(c, su.contextTimeout)
	defer cancel()

	client := apple.New()

	vReq := apple.AppValidationTokenRequest{
		ClientID:     "zyi.featureTesting",
		ClientSecret: clientSecret,
		Code:         code,
	}

	var resp apple.ValidationResponse
	var user domain.User

	err := client.VerifyAppToken(ctx, vReq, &resp)

	if err != nil {
		return user, fmt.Errorf("error verifying: " + err.Error())
	}

	if resp.Error != "" {
		return user, fmt.Errorf("apple returned an error: %v - %v\n", resp.Error, resp.ErrorDescription)
	}

	unique, err := apple.GetUniqueID(resp.IDToken)
	if err != nil {
		return user, fmt.Errorf("failed to get unique ID: " + err.Error())
	}

	claim, err := apple.GetClaims(resp.IDToken)
	if err != nil {
		return user, fmt.Errorf("failed to get claims: " + err.Error())
	}

	user.AppleID = unique
	user.Email = (*claim)["email"].(string)

	return user, nil

}

func (su *signupUsecase) GetByEmail(c context.Context, email string) (domain.User, error) {
	ctx, cancel := context.WithTimeout(c, su.contextTimeout)
	defer cancel()
	return su.userRepository.GetByEmail(ctx, email)
}

func (su *signupUsecase) GetByAppleID(c context.Context, id string) (domain.User, error) {
	ctx, cancel := context.WithTimeout(c, su.contextTimeout)
	defer cancel()
	return su.userRepository.GetByAppleID(ctx, id)
}

func (su *signupUsecase) CreateAccessToken(user *domain.User, secret string, expiry int) (accessToken string, err error) {
	return utils.CreateAccessToken(user, secret, expiry)
}

func (su *signupUsecase) CreateRefreshToken(user *domain.User, secret string, expiry int) (refreshToken string, err error) {
	return utils.CreateRefreshToken(user, secret, expiry)
}
