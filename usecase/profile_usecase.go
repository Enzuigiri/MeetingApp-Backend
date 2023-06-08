package usecase

import (
	"backend/domain"
	"time"

	"golang.org/x/net/context"
)

type profileUsecase struct {
	userRepository domain.UserRepository
	contextTimeout time.Duration
}

func NewProfileUsecase(userRepository domain.UserRepository, contextTimeout time.Duration) domain.ProfileUsecase {
	return &profileUsecase{
		userRepository: userRepository,
		contextTimeout: contextTimeout,
	}
}

func (pu *profileUsecase) GetProfileByID(c context.Context, id string) (domain.Profile, error) {
	ctx, cancel := context.WithTimeout(c, pu.contextTimeout)
	defer cancel()

	var userProfile domain.Profile

	userData, err := pu.userRepository.GetById(ctx, id)
	if err != nil {
		return userProfile, err
	}

	userProfile.ID = id
	userProfile.Email = userData.Email
	userProfile.FirstName = userData.FirstName
	userProfile.LastName = userData.LastName

	return userProfile, nil
}
