package app

import (
	"github.com/go-playground/validator/v10"
	"go.mongodb.org/mongo-driver/mongo"
)

type Application struct {
	Env       *Env
	Mongo     mongo.Client
	Validator validator.Validate
}

func NewApp() Application {
	app := &Application{}
	app.Env = NewEnv()
	app.Mongo = NewMongoDatabase(app.Env)
	app.Validator = *validator.New()
	return *app
}

func (app *Application) CloseDBConnection() {
	CloseMongoDBConnection(&app.Mongo)
}
