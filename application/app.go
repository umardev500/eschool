package application

import (
	"context"
	"fmt"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/rs/zerolog/log"
	"github.com/umardev500/eschool/routes"
)

type Application struct{}

func NewApp() *Application {
	return &Application{}
}

func (a *Application) Start(ctx context.Context) error {
	// Create a new Fiber app
	app := fiber.New(fiber.Config{
		DisableStartupMessage: true,
	})

	// Load the routes
	router := routes.NewRouter(app)
	router.Load()

	ch := make(chan error, 1)
	go func() {
		port := os.Getenv("APP_PORT")
		msg := fmt.Sprintf("🚀 Server is running on port %s", port)
		log.Info().Msg(msg)
		if err := app.Listen(port); err != nil {
			ch <- err
		}
		close(ch)
	}()

	select {
	case err := <-ch:
		return err
	case <-ctx.Done():
		app.Shutdown() // Shutdown the application
		fmt.Println()
		msg := "🚀 Application is shutting down gracefully... 😴"
		log.Info().Msg(msg)
	}

	return nil
}
