package utils

import (
	"math/rand"
	"time"
)

func GenerateRandomCode() string {
	// Define the character set that includes all alphabets (lowercase and uppercase) and numbers
	charSet := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

	// Initialize the random number generator with a seed based on the current time
	rand.Seed(time.Now().UnixNano())

	// Generate a random 4-digit code by selecting 4 characters from the character set
	code := ""
	for i := 0; i < 4; i++ {
		randomIndex := rand.Intn(len(charSet))
		code += string(charSet[randomIndex])
	}

	return code
}
