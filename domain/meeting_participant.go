package domain

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Participant struct {
	UserID    primitive.ObjectID `bson:"userID"`
	FirstName string             `bson:"firstName"`
	JoinTime  time.Time          `bson:"joinTime"`
}

type JoinRequest struct {
	MeetingCode string `json:"meetingCode" validate:"required"`
}

type RemoveRequest struct {
	MeetingId string `json:"meetingId" validate:"required"`
	UserId    string `json:"userId"`
}

type ParticipantUsecase interface {
	CheckMeetingByCode(c context.Context, meetingCode string) (Meeting, error)
	Add(c context.Context, participant *Participant, meeting *Meeting) (Meeting, error)
	Delete(c context.Context, meeting *Meeting, userId string) (Meeting, error)
}
