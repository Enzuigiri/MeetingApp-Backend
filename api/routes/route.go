package route

import (
	"backend/api/middleware"
	"backend/app"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/mongo"
)

func Setup(env *app.Env, timeout time.Duration, db mongo.Database, fiber *fiber.App, validator validator.Validate) {
	publicRouter := fiber.Group("")
	NewSignupRouter(env, timeout, db, publicRouter, validator)

	protectedRouter := fiber.Group("")
	protectedRouter.Use(middleware.JWTAuthMiddleware(env.AccessTokenSecret))

	NewProfileRouter(env, timeout, db, protectedRouter, validator)
	NewMeetingRoute(env, timeout, db, protectedRouter, validator)
}
