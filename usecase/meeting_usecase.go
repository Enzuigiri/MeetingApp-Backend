package usecase

import (
	"backend/domain"
	"backend/internal/utils"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"golang.org/x/net/context"
)

type meetingUsecase struct {
	meetingRepository domain.MeetingRepository
	contextTimeout    time.Duration
}

func NewMeetingUsecase(meetingRepository domain.MeetingRepository, contextTimeout time.Duration) domain.MeetingUsecase {
	return &meetingUsecase{
		meetingRepository: meetingRepository,
		contextTimeout:    contextTimeout,
	}
}

func (mu *meetingUsecase) Create(c context.Context, meeting *domain.Meeting) error {
	ctx, cancel := context.WithTimeout(c, mu.contextTimeout)
	defer cancel()

	meeting.Code = utils.GenerateRandomCode()

	return mu.meetingRepository.Create(ctx, meeting)
}

func (mu *meetingUsecase) FetchByID(c context.Context, userId string, meetId string) (domain.Meeting, error) {
	ctx, cancel := context.WithTimeout(c, mu.contextTimeout)
	defer cancel()

	meeting, err := mu.meetingRepository.FetchByID(ctx, meetId)
	if err != nil {
		return meeting, err
	}

	userIdHex, err := primitive.ObjectIDFromHex(userId)
	if err != nil {
		return meeting, err
	}

	if meeting.PICID.UserID == userIdHex {
		return meeting, err
	}

	for _, user := range meeting.Participants {
		if user.UserID == userIdHex {
			return meeting, err
		}
	}

	return meeting, fmt.Errorf("Not authorized to view this meeting")
}

func (mu *meetingUsecase) FetchByUserID(c context.Context, id string) ([]domain.Meeting, error) {
	ctx, cancel := context.WithTimeout(c, mu.contextTimeout)
	defer cancel()
	return mu.meetingRepository.FetchByUserID(ctx, id)
}

func (mu *meetingUsecase) Update(c context.Context, meeting *domain.Meeting) (domain.Meeting, error) {
	ctx, cancel := context.WithTimeout(c, mu.contextTimeout)
	defer cancel()

	return mu.meetingRepository.Update(ctx, meeting)
}

func (mu *meetingUsecase) Delete(c context.Context, userId string, meetId string) error {
	ctx, cancel := context.WithTimeout(c, mu.contextTimeout)
	defer cancel()
	return mu.meetingRepository.Delete(ctx, userId, meetId)
}
