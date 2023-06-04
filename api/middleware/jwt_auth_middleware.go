package middleware

import (
	"backend/internal/utils"
	"strings"

	"github.com/gofiber/fiber/v2"
)

func JWTAuthMiddleware(secret string) fiber.Handler {
	return func(c *fiber.Ctx) error {
		authHeader := c.Get("Authorization")

		if authHeader == "" {
			return fiber.NewError(fiber.StatusBadRequest, "Missing Authorization header")
		}

		t := strings.Split(authHeader, " ")

		if len(t) == 2 {
			authToken := t[1]
			authorized, err := utils.IsAuthorized(authToken, secret)
			if authorized {
				user, err := utils.ExtractUserFromToken(authToken, secret)
				if err != nil {
					return fiber.NewError(fiber.StatusUnauthorized, err.Error())
				}
				c.Locals("user", user)

				return c.Next()
			}

			return fiber.NewError(fiber.StatusUnauthorized, err.Error())
		}
		return fiber.NewError(fiber.StatusUnauthorized, "Not Authorized")
	}
}
