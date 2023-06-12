package domain

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Agenda struct {
	ID         primitive.ObjectID `bson:"id" json:"id"`
	ProposerID primitive.ObjectID `bson:"proposerID" json:"proposerID"`
	FirstName  string             `bson:"firstName" json:"firstName"`
	Title      string             `bson:"title" json:"title" validate:"required"`
	Desription string             `bson:"description" json:"description" validate:"required"`
	Result     float64            `bson:"result" json:"result"`
	Voters     []Voter            `bson:"voters" json:"voters"`
	CreatedAt  time.Time          `bson:"createdAt" json:"createdAt"`
}

type Voter struct {
	UserId    primitive.ObjectID `bson:"userId" json:"userId"`
	FirstName string             `bson:"firstName" json:"firstName"`
	Value     int                `bson:"value" json:"value"`
}

type AgendaRequest struct {
	Id         string `json:"id,omiempty"`
	Title      string `json:"title" validate:"required"`
	Desription string `json:"description" validate:"required"`
}

type AgendaRequests struct {
	MeetingId string          `json:"meetId" validate:"required"`
	Agendas   []AgendaRequest `json:"agendas" validate:"required"`
}

// type EditAgendaRequest struct {
// 	MeetingId  string `json:"meetId" validate:"required"`
// 	AgendaId   string `json:"agendaId" validate:"required"`
// 	Title      string `json:"title" validate:"required"`
// 	Desription string `json:"description" validate:"required"`
// }

type DeleteAgendaRequest struct {
	MeetingId string `json:"meetId" validate:"required"`
	AgendaId  string `json:"agendaId" validate:"required"`
}

type AgendaVote struct {
	MeetingId string `json:"meetId" validate:"required"`
	VoteValue []int  `json:"voteValue" validate:"required"`
}

type ResultAgendaChangesRequest struct {
	MeetingId string   `json:"meetId" validate:"required"`
	AgendasId []string `json:"agendasId" validate:"required"`
}

type AgendaUsecase interface {
	Add(c context.Context, meeting *Meeting, agenda *[]AgendaRequest, proposerId string, firstName string) (Meeting, error)
	Delete(c context.Context, meeting *Meeting, agendaId string, propeserId string) (Meeting, error)
	Edit(c context.Context, meeting *Meeting, agenda *Agenda, proposerId string) (Meeting, error)
	Vote(c context.Context, meeting *Meeting, votes []int, voter *Voter) (Meeting, error)
	ResultChange(c context.Context, meeting *Meeting, agendasId []string) (Meeting, error)
}
