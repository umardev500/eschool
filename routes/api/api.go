package api

import (
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

type Api struct {
	validate *validator.Validate
}

func NewApi(validate *validator.Validate) *Api {
	return &Api{
		validate: validate,
	}
}

// Load loads the API routes into the provided router.
//
// router: The router to load the routes into.
func (a *Api) Load(router fiber.Router) {
	router.Route("auth", a.Auth) // Authentication router
}
