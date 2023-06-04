package domain

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

const (
	CollectionMeetings = "meetings"
)

type MeetingTime struct {
	Date      time.Time `bson:"date" json:"date" validate:"required,datetime"`
	StartTime time.Time `bson:"startTime" json:"startTime" validate:"required,datetime"`
	EndTime   time.Time `bson:"endTime" json:"endTime" validate:"required,datetime"`
}

type PICID struct {
	UserID    primitive.ObjectID `bson:"userID" json:"userID" `
	FirstName string             `bson:"firstName" json:"firstName"`
}

type Meeting struct {
	ID           primitive.ObjectID `bson:"_id,omitempty" json:"meetId"`
	Title        string             `bson:"title" json:"title" validate:"required"`
	Desription   string             `bson:"description" json:"description" validate:"required"`
	PICID        PICID              `bson:"picid" json:"picid" validate:"required"`
	Code         string             `bson:"code" json:"code"`
	Location     string             `bson:"location" json:"location" validate:"required"`
	Schedule     MeetingTime        `bson:"schedule" json:"schedule" validate:"required"`
	VoteTime     MeetingTime        `bson:"voteTime" json:"voteTime" validate:"required"`
	Participants []Participant      `bson:"participants" json:"participants"`
	Agenda       []Agenda           `bson:"agenda" json:"agenda" validate:"required"`
	CreatedAt    time.Time          `bson:"createdAt"`
}

type FetchMeetingRequest struct {
	ID string `json:"meetID" validate:"required"`
}

type FetchAllUserMeetingResponse struct {
	ID           string        `json:"meetId"`
	Title        string        `json:"title"`
	Location     string        `json:"location"`
	Schedule     MeetingTime   `json:"schedule"`
	VoteTime     MeetingTime   `json:"voteTime"`
	Participants []Participant `json:"participants"`
}

type UpdateMeetingRequest struct {
	ID          string      `json:"meetId" validate:"required"`
	Title       string      `json:"title" validate:"required"`
	Description string      `json:"description" validate:"required"`
	Location    string      `json:"location" validate:"required"`
	Schedule    MeetingTime `json:"schedule" validate:"required"`
	VoteTime    MeetingTime `json:"voteTime" validate:"required"`
}

type MeetingRepository interface {
	Create(c context.Context, meeting *Meeting) error
	FetchByID(c context.Context, id string) (Meeting, error)
	FetchByUserID(c context.Context, id string) ([]Meeting, error)
	Update(c context.Context, meeting *Meeting) (Meeting, error)
	Delete(c context.Context, userId string, meetId string) error
}

type MeetingUsecase interface {
	Create(c context.Context, meeting *Meeting) error
	FetchByID(c context.Context, userId string, meetId string) (Meeting, error)
	FetchByUserID(c context.Context, id string) ([]Meeting, error)
	Update(c context.Context, meeting *Meeting) (Meeting, error)
	Delete(c context.Context, userId string, meetId string) error
}
