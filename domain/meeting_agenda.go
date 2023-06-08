package domain

import "go.mongodb.org/mongo-driver/bson/primitive"

type Agenda struct {
	ID         primitive.ObjectID `bson:"id" json:"id"`
	ProposerID primitive.ObjectID `bson:"proposerID" json:"proposerID"`
	FirstName  string             `bson:"firstName" json:"firstName"`
	Title      string             `bson:"title" json:"title" validate:"required"`
	Desription string             `bson:"description" json:"description" validate:"required"`
	Result     float64            `bson:"result" json:"result"`
}
