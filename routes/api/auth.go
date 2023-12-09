package api

import (
	"github.com/gofiber/fiber/v2"
	authhandler "github.com/umardev500/eschool/application/handlers/auth"
)

func (a *Api) Auth(router fiber.Router) {
	handler := authhandler.NewAuthHandler()

	router.Post("/", handler.Login)
}
