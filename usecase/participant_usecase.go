package usecase

import (
	"backend/domain"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"golang.org/x/net/context"
)

type participantUsecase struct {
	meetingRepository domain.MeetingRepository
	contextTimeout    time.Duration
}

func NewParticipantUsecase(meetingRepository domain.MeetingRepository,
	contextTimeout time.Duration) domain.ParticipantUsecase {
	return &participantUsecase{
		meetingRepository: meetingRepository,
		contextTimeout:    contextTimeout,
	}
}

func (pu *participantUsecase) CheckMeetingByCode(c context.Context,
	meetingCode string) (domain.Meeting, error) {
	ctx, cancel := context.WithTimeout(c, pu.contextTimeout)
	defer cancel()
	return pu.meetingRepository.FetchByCode(ctx, meetingCode)
}

func (pu *participantUsecase) Add(c context.Context,
	participant *domain.Participant, meeting *domain.Meeting) (domain.Meeting, error) {
	ctx, cancel := context.WithTimeout(c, pu.contextTimeout)
	defer cancel()

	for _, meetParticipant := range meeting.Participants {
		if meetParticipant.UserID == participant.UserID {
			return *meeting, nil
		}
	}

	meeting.Participants = append(meeting.Participants, *participant)
	return pu.meetingRepository.Update(ctx, meeting)
}

func (pu *participantUsecase) Delete(c context.Context,
	meeting *domain.Meeting, userId string) (domain.Meeting, error) {
	ctx, cancel := context.WithTimeout(c, pu.contextTimeout)
	defer cancel()

	userIdHex, err := primitive.ObjectIDFromHex(userId)
	if err != nil {
		return domain.Meeting{}, err
	}

	for i := 0; i < len(meeting.Participants); i++ {
		if meeting.Participants[i].UserID == userIdHex {
			meeting.Participants = append(meeting.Participants[:i], meeting.Participants[i+1:]...)
			break
		}
	}

	for i := 0; i < len(meeting.Participants); i++ {
		if meeting.Agenda[i].ProposerID == userIdHex {
			meeting.Agenda = append(meeting.Agenda[:i], meeting.Agenda[i+1:]...)
		}
	}

	return pu.meetingRepository.Update(ctx, meeting)
}
