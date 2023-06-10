package usecase

import (
	"backend/domain"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"golang.org/x/net/context"
)

type agendaUsecase struct {
	meetingRepository domain.MeetingRepository
	contextTimeout    time.Duration
}

func NewAgendaUsecase(meetingRepository domain.MeetingRepository,
	contextTimeout time.Duration) domain.AgendaUsecase {
	return &agendaUsecase{
		meetingRepository: meetingRepository,
		contextTimeout:    contextTimeout,
	}
}

func (au *agendaUsecase) Add(c context.Context, meeting *domain.Meeting,
	agendasParam *[]domain.AgendaRequest, proposerId string, firstName string) (domain.Meeting, error) {
	ctx, cancel := context.WithTimeout(c, au.contextTimeout)
	defer cancel()

	// Fix
	// now := time.Now()

	// if meeting.VoteTime.StartTime < now && meeting.VoteTime.EndTime > now {

	// }

	proposerIdHex, err := primitive.ObjectIDFromHex(proposerId)
	if err != nil {
		return domain.Meeting{}, err
	}

	var agendas []domain.Agenda

	for _, agenda := range *agendasParam {
		if agenda.Id == "" {
			agendas = append(agendas, domain.Agenda{
				ID:         primitive.NewObjectID(),
				ProposerID: proposerIdHex,
				FirstName:  firstName,
				Title:      agenda.Title,
				Desription: agenda.Desription,
				CreatedAt:  time.Now(),
				Voters:     []domain.Voter{},
			})
		}
		if agenda.Id != "" {
			agendaIdHex, err := primitive.ObjectIDFromHex(agenda.Id)
			if err != nil {
				return domain.Meeting{}, err
			}
			for i := 0; i < len(meeting.Agenda); i++ {
				if meeting.Agenda[i].ID == agendaIdHex && meeting.Agenda[i].ProposerID == proposerIdHex {
					meeting.Agenda[i].Title = agenda.Title
					meeting.Agenda[i].Desription = agenda.Desription
					break
				}
			}
		}
	}

	meeting.Agenda = append(meeting.Agenda, agendas...)

	return au.meetingRepository.Update(ctx, meeting)
}

func (au *agendaUsecase) Edit(c context.Context,
	meeting *domain.Meeting, agenda *domain.Agenda, proposerId string) (domain.Meeting, error) {
	ctx, cancel := context.WithTimeout(c, au.contextTimeout)
	defer cancel()

	isValid := false
	proposerIdHex, err := primitive.ObjectIDFromHex(proposerId)
	if err != nil {
		return domain.Meeting{}, err
	}

	for i := 0; i < len(meeting.Agenda); i++ {
		if meeting.Agenda[i].ID == agenda.ID && meeting.Agenda[i].ProposerID == proposerIdHex {
			meeting.Agenda[i].Title = agenda.Title
			meeting.Agenda[i].Desription = agenda.Desription
			isValid = true
			break
		}
	}

	if isValid {
		return au.meetingRepository.Update(ctx, meeting)
	}

	return domain.Meeting{}, fmt.Errorf("Not authorized or agenda not exist")
}

func (au *agendaUsecase) Delete(c context.Context,
	meeting *domain.Meeting, agendaId string, propeserId string) (domain.Meeting, error) {
	ctx, cancel := context.WithTimeout(c, au.contextTimeout)
	defer cancel()

	isValid := false
	proposerIdHex, err := primitive.ObjectIDFromHex(propeserId)
	if err != nil {
		return domain.Meeting{}, err
	}

	agendaIdHex, err := primitive.ObjectIDFromHex(agendaId)
	if err != nil {
		return domain.Meeting{}, err
	}

	for i := 0; i < len(meeting.Agenda); i++ {
		if meeting.Agenda[i].ID == agendaIdHex &&
			meeting.Agenda[i].ProposerID == proposerIdHex ||
			meeting.PICID.UserID == proposerIdHex {
			meeting.Participants = append(meeting.Participants[:i], meeting.Participants[i+1:]...)
			isValid = true
			break
		}
	}

	if isValid {
		return au.meetingRepository.Update(ctx, meeting)
	}

	return domain.Meeting{}, fmt.Errorf("Not authorized or agenda not exist")
}

func (au *agendaUsecase) Vote(c context.Context,
	meeting *domain.Meeting, votes []int, voter *domain.Voter) (domain.Meeting, error) {
	ctx, cancel := context.WithTimeout(c, au.contextTimeout)
	defer cancel()

	tempVoter := voter

	for i := 0; i < len(meeting.Agenda); i++ {
		tempVoter.Value = votes[i]
		meeting.Agenda[i].Voters = append(meeting.Agenda[i].Voters, *tempVoter)
	}

	return au.meetingRepository.Update(ctx, meeting)
}