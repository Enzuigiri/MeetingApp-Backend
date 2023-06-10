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

func NewMeetingRoute(env *app.Env, timeout time.Duration, db mongo.Database, group fiber.Router, validator validator.Validate) {
	ur := repository.NewMeetingRepository(db, domain.CollectionMeetings)
	mc := controllers.MeetingController{
		MeetingUsecase: usecase.NewMeetingUsecase(ur, timeout),
		Env:            env,
		Validator:      validator,
	}

	group.Post("/meeting", mc.Create)
	group.Post("/meeting/id", mc.FetchByID)
	group.Get("/user/meetings", mc.FetchByUserID)
	group.Put("/meeting", mc.Update)
	group.Delete("/meeting", mc.Delete)
}
