package api

import (
	"context"
	"os"

	"github.com/go-playground/validator"
	"github.com/gofiber/fiber"
	"github.com/gofiber/session"
	"github.com/rs/zerolog"

	"github.com/hi019/fiber-boilerplate/ent"
	us "github.com/hi019/fiber-boilerplate/pkg/api/user"
	uw "github.com/hi019/fiber-boilerplate/pkg/api/user/web"

	_ "github.com/mattn/go-sqlite3"
)

// Config contains config values for the API
type Config struct {
	Port           string
	DriverName     string
	DataSourceName string
}

// Start starts the api
func Start(cfg *Config) (*fiber.App, *ent.Client, error) {
	// Configure db
	db, err := ent.Open(cfg.DriverName, cfg.DataSourceName)
	if err != nil {
		panic(err)
	}

	// Run the auto migration tool.
	if err := db.Schema.Create(context.Background()); err != nil {
		panic(err)
	}

	// Configure logger and validator
	logger := zerolog.New(zerolog.ConsoleWriter{Out: os.Stderr}).With().Timestamp().Logger()

	vd := validator.New()

	// Initialize session
	sessions := session.New()

	// Default error handler
	errorHandler := func(ctx *fiber.Ctx, err error) error {
		// Statuscode defaults to 500
		code := fiber.StatusInternalServerError
		message := "Internal Server Error"

		// Retrieve the custom statuscode if it's an fiber.*Error
		if e, ok := err.(*fiber.Error); ok {
			code = e.Code
			message = e.Message
		}

		// If its an Internal Server Error, we'll log it
		if code == fiber.StatusInternalServerError {
			logger.Error().Msg(err.Error())
		}

		return ctx.Status(code).JSON(fiber.Map{"status": "error", "message": message})
	}

	// Configure api
	app := fiber.New(fiber.Config{ErrorHandler: errorHandler})

	// Configure api routes and services
	uw.NewHTTP(us.Initialize(db, &logger), app, vd, sessions)

	return app, db, nil
}
