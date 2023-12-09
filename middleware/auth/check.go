package auth

import (
	"os"

	jwtware "github.com/gofiber/contrib/jwt"
	"github.com/gofiber/fiber/v2"
)

func Check() fiber.Handler {
	secret := os.Getenv("JWT_SECRET")

	return jwtware.New(jwtware.Config{
		SuccessHandler: authSuccessHandler,
		SigningKey: jwtware.SigningKey{
			Key: []byte(secret),
		},
	})
}

func authSuccessHandler(c *fiber.Ctx) error {
	return c.Next()
}
