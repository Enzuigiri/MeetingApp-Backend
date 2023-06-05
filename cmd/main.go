package main

import (
	route "backend/api/routes"
	"backend/app"
	"fmt"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func helloWorld(c *fiber.Ctx) error {
	return c.SendString("Hello, World!")
}

func main() {
	fmt.Println("Hello world")
	app := app.NewApp()

	env := app.Env

	validator := app.Validator

	db := app.Mongo.Database(env.DBName)
	defer app.CloseDBConnection()

	timeout := time.Duration(env.ContextTimeout) * time.Second

	fiber := fiber.New()

	fiber.Use(logger.New())

	fiber.Get("/test", helloWorld)

	route.Setup(env, timeout, *db, fiber, validator)

	err := fiber.Listen(":8000")

	if err != nil {
		panic(err)
	}

	// app.Post("/auth/validation", func(c *fiber.Ctx) error {

	// 	type CodeValidate struct {
	// 		Token string `json:"token"`
	// 		FirstName string `json:"first_name"`
	// 		LastName string `json:"last_name"`
	// 	}

	// 	var body CodeValidate
	// 	err := c.BodyParser(&body)

	// 	if err != nil {
	// 		fmt.Println(err)
	// 		return c.SendStatus(fiber.StatusInternalServerError)
	// 	}

	// 	println(body.Token)

	// 	// teamID := "BS6K64RQ23"
	// 	// clientID := "zyi.featureTesting"
	// 	// keyID := "3QQV4Z6GA9"
	// 	// secret := `-----BEGIN PRIVATE KEY-----
	// 	// MIGTAgEAMBMGByqGSM49AgEGCCqGSM49AwEHBHkwdwIBAQQg+64JDWKrQOckjb+x
	// 	// HqMbwySJ9e3qJaAJ90idw4J3oe2gCgYIKoZIzj0DAQehRANCAAR4ekcM4E8+uwmN
	// 	// a/Mq7C4WA8F0gYWdUKws0ZX11NLfpucqYQbEE9yC9trSNvRBfWLlmZgM37r3jN6+
	// 	// EqA/Ujqg
	// 	// -----END PRIVATE KEY-----`

	// 	// secret, err = apple.GenerateClientSecret(secret, teamID, clientID, keyID)
	// 	// if err != nil {
	// 	// 	fmt.Println("error generating secret: " + err.Error())
	// 	// 	return c.SendStatus(fiber.StatusInternalServerError)
	// 	// }

	// 	client := apple.New()

	// 	vReq := apple.AppValidationTokenRequest{
	// 		ClientID:     "zyi.featureTesting",
	// 		ClientSecret: "",
	// 		Code:         body.Token,
	// 	}

	// 	var resp apple.ValidationResponse

	// 	err = client.VerifyAppToken(c.Context(), vReq, &resp)

	// 	if err != nil {
	// 		fmt.Println("error verifying: " + err.Error())
	// 		return c.SendStatus(fiber.StatusInternalServerError)
	// 	}

	// 	if resp.Error != "" {
	// 		fmt.Printf("apple returned an error: %s - %s\n", resp.Error, resp.ErrorDescription)
	// 		return c.SendStatus(fiber.StatusInternalServerError)
	// 	}

	// 	unique, err := apple.GetUniqueID(resp.IDToken)
	// 	if err != nil {
	// 		fmt.Println("failed to get unique ID: " + err.Error())
	// 		return c.SendStatus(fiber.StatusInternalServerError)
	// 	}

	// 	claim, err := apple.GetClaims(resp.IDToken)
	// 	if err != nil {
	// 		fmt.Println("failed to get claims: " + err.Error())
	// 		return c.SendStatus(fiber.StatusInternalServerError)
	// 	}

	// 	email := (*claim)["email"]
	// 	emailVerified := (*claim)["email_verified"]
	// 	isPrivateEmail := (*claim)["is_private_email"]

	// 	fmt.Println(unique)
	// 	fmt.Println(email)
	// 	fmt.Println(emailVerified)
	// 	fmt.Println(isPrivateEmail)

	// 	return c.SendStatus(fiber.StatusOK)

	// })
}
