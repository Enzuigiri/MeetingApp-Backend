package domain

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

const (
	CollectionUsers = "users"
)

type User struct {
	ID           primitive.ObjectID `bson:"_id,omitempty"`
	AppleID      string             `bson:"appleID"`
	FirstName    string             `bson:"firstName"`
	LastName     string             `bson:"lastName"`
	Email        string             `bson:"email"`
	RefreshToken string             `bson:"refreshToken"`
	LastLogin    time.Time          `bson:"lastLogin"`
	CreatedAt    time.Time          `bson:"createdAt"`
}

type UserRepository interface {
	Create(c context.Context, user *User) error
	// Fetch(c context.Context) error
	GetById(c context.Context, id string) (User, error)
	GetByEmail(c context.Context, email string) (User, error)
	GetByAppleID(c context.Context, id string) (User, error)
	Update(c context.Context, user *User) (User, error)
}
