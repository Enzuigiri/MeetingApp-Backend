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

func NewAgendaRoute(env *app.Env, timeout time.Duration,
	db mongo.Database, group fiber.Router, validator validator.Validate) {
	ur := repository.NewMeetingRepository(db, domain.CollectionMeetings)
	ac := controllers.AgendaController{
		MeetingUsecase: usecase.NewMeetingUsecase(ur, timeout),
		AgendaUsecase:  usecase.NewAgendaUsecase(ur, timeout),
		Env:            env,
		Validator:      validator,
	}

	group.Post("/meeting/agenda", ac.Add)
	group.Post("/meeting/agenda/result", ac.SaveResult)
	group.Delete("/meeting/agenda", ac.Delete)
	group.Post("/meeting/vote", ac.Vote)
}
