package routes

import (
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/umardev500/eschool/routes/api"
)

type Router struct {
	app *fiber.App
}

func NewRouter(app *fiber.App) *Router {
	return &Router{
		app: app,
	}
}

func (r *Router) Load() {
	v := validator.New()

	// Middlewares
	r.app.Use(cors.New())

	api := api.NewApi(v)
	r.app.Route("api", api.Load)
}
