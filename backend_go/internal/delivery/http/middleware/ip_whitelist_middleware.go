package middleware

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/hafisrabbani/technical-test-nexmedis/internal/repository"
	"github.com/redis/go-redis/v9"
)

func IPWhitelist(
	rdb *redis.Client,
	repo *repository.IPWhitelistRepository,
) fiber.Handler {
	return func(c *fiber.Ctx) error {
		client := c.Locals("client").(*repository.Client)
		clientUUID := client.ID
		ip := c.IP()

		if ip == "" {
			return fiber.ErrForbidden
		}

		key := fmt.Sprintf("ip_whitelist:%s", clientUUID)

		exists, err := rdb.Exists(c.Context(), key).Result()
		if err == nil && exists > 0 {
			ok, _ := rdb.SIsMember(c.Context(), key, ip).Result()
			if !ok {
				return c.Status(fiber.StatusForbidden).
					JSON(fiber.Map{"message": "ip not allowed"})
			}
			return c.Next()
		}

		ips, err := repo.FindByClientID(c.Context(), clientUUID)
		if err != nil || len(ips) == 0 {
			// empty = wildcard
			return c.Next()
		}

		_ = rdb.SAdd(c.Context(), key, ips).Err()

		for _, allowedIP := range ips {
			if allowedIP == ip {
				return c.Next()
			}
		}

		return c.Status(fiber.StatusForbidden).
			JSON(fiber.Map{"message": "ip not allowed"})
	}
}
