package repository

import (
	"backend/domain"
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type userRepository struct {
	database   mongo.Database
	collection string
}

func NewUserRepository(db mongo.Database, collection string) domain.UserRepository {
	return &userRepository{
		database:   db,
		collection: collection,
	}
}

func (ur *userRepository) Create(c context.Context, user *domain.User) error {
	collection := ur.database.Collection(ur.collection)

	_, err := collection.InsertOne(c, user)

	return err
}

func (ur *userRepository) GetByEmail(c context.Context, email string) (domain.User, error) {
	collection := ur.database.Collection(ur.collection)

	var user domain.User

	err := collection.FindOne(c, bson.M{"email": email}).Decode(&user)

	return user, err
}

func (ur *userRepository) GetByAppleID(c context.Context, id string) (domain.User, error) {
	collection := ur.database.Collection(ur.collection)

	var user domain.User

	err := collection.FindOne(c, bson.M{"appleID": id}).Decode(&user)

	return user, err
}

func (ur *userRepository) GetById(c context.Context, id string) (domain.User, error) {
	collection := ur.database.Collection(ur.collection)

	var user domain.User

	idHex, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return user, err
	}

	err = collection.FindOne(c, bson.M{"_id": idHex}).Decode(&user)

	return user, err
}

func (ur *userRepository) Update(c context.Context, user *domain.User) (domain.User, error) {
	collection := ur.database.Collection(ur.collection)

	filter := bson.M{"_id": user.ID}

	update := bson.M{"$set": user}

	var result domain.User

	err := collection.FindOneAndUpdate(c, filter, update).Decode(result)

	return result, err
}
