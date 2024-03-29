package domain

import "context"

type Profile struct {
	ID        string `json:"id"`
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Email     string `json:"email"`
}

type ProfileUsecase interface {
	GetProfileByID(c context.Context, id string) (Profile, error)
}
