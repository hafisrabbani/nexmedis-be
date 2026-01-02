package usage

import (
	"context"

	"github.com/gofiber/fiber/v2"
	"github.com/hafisrabbani/technical-test-nexmedis/internal/delivery/http/middleware"
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"

	httpDelivery "github.com/hafisrabbani/technical-test-nexmedis/internal/delivery/http"
	"github.com/hafisrabbani/technical-test-nexmedis/internal/repository"
	"github.com/hafisrabbani/technical-test-nexmedis/internal/service"
)

func Register(
	app *fiber.App,
	db *gorm.DB,
	rdb *redis.Client,
	apiKeyMiddleware fiber.Handler,
	jwtMiddleware fiber.Handler,
	ipWhitelistMiddleware fiber.Handler,
) {
	repo := repository.NewUsageRepository(db, rdb)
	clientRepo := repository.NewClientRepository(db)
	svc := service.NewUsageService(repo)
	worker := service.NewUsageBatchWorker(repo, clientRepo, rdb)
	worker.Start(context.Background())

	rateLimit := middleware.RateLimit(rdb)
	httpDelivery.RegisterUsageRoutes(
		app,
		svc,
		apiKeyMiddleware,
		jwtMiddleware,
		rateLimit,
		ipWhitelistMiddleware,
	)

}
