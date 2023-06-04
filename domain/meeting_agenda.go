package domain

import "go.mongodb.org/mongo-driver/bson/primitive"

type Agenda struct {
	ID         primitive.ObjectID `bson:"id"`
	ProposerID primitive.ObjectID `bson:"proposerID"`
	FirstName  string             `bson:"firstName"`
	Title      string             `bson:"title"`
	Desription string             `bson:"description"`
	Result     float64            `bson:"result"`
}
