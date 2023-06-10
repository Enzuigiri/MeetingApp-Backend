package repository

import (
	"backend/domain"
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type meetingRepository struct {
	database   mongo.Database
	collection string
}

func NewMeetingRepository(db mongo.Database, collection string) domain.MeetingRepository {
	return &meetingRepository{
		database:   db,
		collection: collection,
	}
}

func (mr *meetingRepository) Create(c context.Context, meeting *domain.Meeting) error {
	collection := mr.database.Collection(mr.collection)

	_, err := collection.InsertOne(c, meeting)

	return err
}

func (mr *meetingRepository) FetchByID(c context.Context, id string) (domain.Meeting, error) {
	collection := mr.database.Collection(mr.collection)

	var meeting domain.Meeting

	err := collection.FindOne(c, bson.M{"_id": id}).Decode(&meeting)

	return meeting, err
}

func (mr *meetingRepository) FetchByCode(c context.Context, code string) (domain.Meeting, error) {
	collection := mr.database.Collection(mr.collection)

	var meeting domain.Meeting

	err := collection.FindOne(c, bson.M{"code": code}).Decode(&meeting)

	return meeting, err
}

func (mr *meetingRepository) FetchByUserID(c context.Context, id string) ([]domain.Meeting, error) {
	collection := mr.database.Collection(mr.collection)

	var meetings []domain.Meeting

	idHex, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return meetings, err
	}

	filter := bson.M{
		"$or": bson.A{
			bson.M{"PICID": bson.M{"userID": idHex}},
			bson.M{"participants": bson.M{"$elemMatch": bson.M{"userID": idHex}}},
		},
	}

	cursor, err := collection.Find(c, filter)
	if err != nil {
		return meetings, err
	}

	err = cursor.All(c, &meetings)
	if err != nil {
		return meetings, err
	}

	return meetings, err
}

func (mr *meetingRepository) Update(c context.Context, meeting *domain.Meeting) (domain.Meeting, error) {
	collection := mr.database.Collection(mr.collection)

	filter := bson.M{"_id": meeting.ID}

	update := bson.M{"$set": meeting}

	var result domain.Meeting

	err := collection.FindOneAndUpdate(c, filter, update).Decode(result)

	return result, err
}

func (mr *meetingRepository) Delete(c context.Context, userId string, meetId string) error {
	collection := mr.database.Collection(mr.collection)

	userIdHex, err := primitive.ObjectIDFromHex(userId)
	if err != nil {
		return err
	}

	meetIdHex, err := primitive.ObjectIDFromHex(meetId)
	if err != nil {
		return err
	}

	filter := bson.M{
		"_id":   meetIdHex,
		"picid": bson.M{"userID": userIdHex},
	}

	err = collection.FindOneAndDelete(c, filter).Err()

	return err
}
