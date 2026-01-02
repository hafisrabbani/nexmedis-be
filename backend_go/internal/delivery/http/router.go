package http

import (
	"github.com/gofiber/fiber/v2"
	"github.com/hafisrabbani/technical-test-nexmedis/internal/delivery/http/handler"
	"github.com/hafisrabbani/technical-test-nexmedis/internal/service"
)

// CLIENT ROUTES (PUBLIC)
func RegisterRoutes(
	app *fiber.App,
	clientService *service.ClientService,
) {
	api := app.Group("/api")

	h := handler.NewClientHandler(clientService)

	// public
	api.Post("/register", h.Register)
}

// AUTH ROUTES (API KEY)
func RegisterAuthRoutes(
	app *fiber.App,
	jwtService *service.JWTService,
	apiKeyMiddleware fiber.Handler,
) {
	api := app.Group("/api", apiKeyMiddleware)

	h := handler.NewTokenHandler(jwtService)

	api.Post("/token", h.Issue)
}

// USAGE ROUTES
func RegisterUsageRoutes(
	app *fiber.App,
	usageService *service.UsageService,
	apiKeyMiddleware fiber.Handler,
	jwtMiddleware fiber.Handler,
	rateLimitMiddleware fiber.Handler,
	ipWhitelistMiddleware fiber.Handler,
) {
	h := handler.NewUsageHandler(usageService)

	logs := app.Group(
		"/api",
		apiKeyMiddleware,
		ipWhitelistMiddleware,
	)

	logs.Post("/logs", h.Log)

	usage := app.Group("/api",
		apiKeyMiddleware,
		jwtMiddleware,
		rateLimitMiddleware,
	)

	usage.Get("/usage/daily", h.Daily)
	usage.Get("/usage/top", h.Top)
	usage.Get("/usage/stream", h.RealtimeDailyUsage)
}

// IPWHITELIST
func RegisterIPWhitelistRoutes(
	app *fiber.App,
	service *service.IPWhitelistService,
	apiKey fiber.Handler,
	jwt fiber.Handler,
) {
	api := app.Group("/api/whitelist",
		apiKey,
		jwt,
	)

	h := handler.NewIPWhitelistHandler(service)
	api.Post("", h.Register)
}
