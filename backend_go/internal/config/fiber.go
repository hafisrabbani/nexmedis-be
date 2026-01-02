package config

import (
	"encoding/json"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/joho/godotenv"
)

func NewFiber() *fiber.App {
	_ = godotenv.Load()

	app := fiber.New(fiber.Config{
		StrictRouting: true,
		Prefork:       true,
		CaseSensitive: true,
		AppName:       os.Getenv("APP_NAME"),
		JSONEncoder:   json.Marshal,
		JSONDecoder:   json.Unmarshal,
		BodyLimit:     2 * 1024 * 1024, // 2 MB
	})

	app.Use(logger.New(logger.Config{
		Format:     "${pid} ${status} - ${method} ${path}\n",
		TimeFormat: "02-Jan-2006",
		TimeZone:   os.Getenv("APP_TIMEZONE"),
	}))

	app.Use(recover.New())

	return app
}
