package authhandler

import (
	"github.com/gofiber/fiber/v2"
)

func (a *authHandler) Login(c *fiber.Ctx) error {

	return c.JSON("ok")
}
