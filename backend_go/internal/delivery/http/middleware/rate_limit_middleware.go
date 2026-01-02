package middleware

import (
	"fmt"
	"strconv"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/hafisrabbani/technical-test-nexmedis/internal/model/response"
	"github.com/redis/go-redis/v9"

	"github.com/hafisrabbani/technical-test-nexmedis/internal/repository"
	"github.com/hafisrabbani/technical-test-nexmedis/internal/shared"
)

func RateLimit(rdb *redis.Client) fiber.Handler {
	limit, _ := strconv.Atoi(shared.GetEnv("RATE_LIMIT_PER_HOUR", "1000"))

	return func(c *fiber.Ctx) error {
		client := c.Locals("client").(*repository.Client)

		key := fmt.Sprintf(
			"ratelimit:%s:%s",
			client.ClientID,
			time.Now().Format("2006-01-02-15"),
		)

		count, err := rdb.Incr(c.Context(), key).Result()
		if err != nil {
			return c.Next()
		}

		if count == 1 {
			_ = rdb.Expire(c.Context(), key, time.Hour).Err()
		}

		if count > int64(limit) {
			return c.Status(fiber.StatusTooManyRequests).JSON(response.Error("rate limit exceeded, try again later"))
		}

		return c.Next()
	}
}
