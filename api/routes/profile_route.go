package route

import (
	"backend/api/controllers"
	"backend/app"
	"backend/domain"
	"backend/repository"
	"backend/usecase"

	"time"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/mongo"
)

func NewProfileRouter(env *app.Env, timeout time.Duration, db mongo.Database, group fiber.Router, validator validator.Validate) {
	ur := repository.NewUserRepository(db, domain.CollectionUsers)
	pc := controllers.ProfileController{
		ProfileUsecase: usecase.NewProfileUsecase(ur, timeout),
		Env:            env,
	}
	group.Get("/user", pc.Profile)
}
