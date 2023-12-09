package domain

import "github.com/gofiber/fiber/v2"

type AuthHandler interface {
	Login(c *fiber.Ctx) error
}
