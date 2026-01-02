package middleware

import (
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/hafisrabbani/technical-test-nexmedis/internal/service"
)

func JWTAuth(jwtService *service.JWTService) fiber.Handler {
	return func(c *fiber.Ctx) error {
		auth := c.Get("Authorization")
		if auth == "" {
			return fiber.ErrUnauthorized
		}

		parts := strings.Split(auth, " ")
		if len(parts) != 2 || parts[0] != "Bearer" {
			return fiber.ErrUnauthorized
		}

		claims, err := jwtService.Validate(parts[1])
		if err != nil {
			return fiber.ErrUnauthorized
		}

		c.Locals("jwt_claims", claims)
		return c.Next()
	}
}
