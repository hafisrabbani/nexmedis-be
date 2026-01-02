package client

import (
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"

	httpDelivery "github.com/hafisrabbani/technical-test-nexmedis/internal/delivery/http"
	"github.com/hafisrabbani/technical-test-nexmedis/internal/delivery/http/middleware"
	"github.com/hafisrabbani/technical-test-nexmedis/internal/repository"
	"github.com/hafisrabbani/technical-test-nexmedis/internal/service"
)

type Module struct {
	APIKeyMiddleware fiber.Handler
	JWTService       *service.JWTService
}

func Register(app *fiber.App, db *gorm.DB) *Module {
	repo := repository.NewClientRepository(db)
	svc := service.NewClientService(repo)

	jwtService := service.NewJWTService()

	httpDelivery.RegisterRoutes(app, svc)
	httpDelivery.RegisterAuthRoutes(
		app,
		jwtService,
		middleware.APIKeyAuth(svc),
	)

	return &Module{
		APIKeyMiddleware: middleware.APIKeyAuth(svc),
		JWTService:       jwtService,
	}
}
