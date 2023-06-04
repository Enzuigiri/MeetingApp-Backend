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

type ParticipantUsecase interface {
	CheckUserId(c context.Context, userId string, meetingId string) (Meeting, error)
	Add(c context.Context, userId string, meetingId string) error
	Delete(c context.Context, userId string, meetingId string) error
}
