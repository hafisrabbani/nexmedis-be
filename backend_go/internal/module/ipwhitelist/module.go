package ipwhitelist

import (
	"github.com/gofiber/fiber/v2"
	httpDelivery "github.com/hafisrabbani/technical-test-nexmedis/internal/delivery/http"
	"github.com/hafisrabbani/technical-test-nexmedis/internal/delivery/http/middleware"
	"github.com/hafisrabbani/technical-test-nexmedis/internal/repository"
	"github.com/hafisrabbani/technical-test-nexmedis/internal/service"
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

type Module struct {
	IPWhitelistMiddleware fiber.Handler
}

func Register(
	app *fiber.App,
	db *gorm.DB,
	rdb *redis.Client,
	apiKeyMiddleware fiber.Handler,
	jwtMiddleware fiber.Handler,
) *Module {
	repo := repository.NewIPWhitelistRepository(db)
	svc := service.NewIPWhitelistService(repo, rdb)

	httpDelivery.RegisterIPWhitelistRoutes(
		app,
		svc,
		apiKeyMiddleware,
		jwtMiddleware,
	)

	return &Module{
		IPWhitelistMiddleware: middleware.IPWhitelist(rdb, repo),
	}
}
