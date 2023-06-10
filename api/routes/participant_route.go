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

func NewParticipantRoute(env *app.Env, timeout time.Duration,
	db mongo.Database, group fiber.Router, validator validator.Validate) {
	ur := repository.NewMeetingRepository(db, domain.CollectionMeetings)
	pc := controllers.ParticipantController{
		MeetingUsecase:    usecase.NewMeetingUsecase(ur, timeout),
		ParticpantUsecase: usecase.NewParticipantUsecase(ur, timeout),
		Env:               env,
		Validator:         validator,
	}

	group.Post("/meeting/participant", pc.Join)
	group.Delete("/meeting/participant", pc.Remove)
}
