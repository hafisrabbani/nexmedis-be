package middleware

import (
	"strings"

	"github.com/gofiber/fiber/v2"

	"github.com/hafisrabbani/technical-test-nexmedis/internal/service"
)

func APIKeyAuth(clientService *service.ClientService) fiber.Handler {
	return func(c *fiber.Ctx) error {
		apiKey := c.Get("X-API-Key")
		if apiKey == "" {
			return fiber.ErrUnauthorized
		}

		apiKey = strings.TrimSpace(apiKey)

		client, err := clientService.ValidateAPIKey(
			c.Context(),
			apiKey,
		)
		if err != nil {
			return fiber.ErrUnauthorized
		}

		c.Locals("client", client)

		return c.Next()
	}
}
