package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/hafisrabbani/technical-test-nexmedis/internal/config"
	"github.com/hafisrabbani/technical-test-nexmedis/internal/delivery/http/middleware"
	clientModule "github.com/hafisrabbani/technical-test-nexmedis/internal/module/client"
	"github.com/hafisrabbani/technical-test-nexmedis/internal/module/ipwhitelist"
	usageModule "github.com/hafisrabbani/technical-test-nexmedis/internal/module/usage"
	"github.com/hafisrabbani/technical-test-nexmedis/internal/shared"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	app := config.NewFiber()
	db, err := config.NewPostgres()
	if err != nil {
		log.Fatalf("db init failed: %v", err)
	}

	sqlDB, _ := db.DB()
	defer sqlDB.Close()

	rdb, err := config.NewRedis(ctx)
	if err != nil {
		log.Fatalf("redis init failed: %v", err)
	}
	defer rdb.Close()

	clientModule := clientModule.Register(app, db)

	jwtMiddleware := middleware.JWTAuth(clientModule.JWTService)

	ipWhiteListModule := ipwhitelist.Register(
		app,
		db,
		rdb,
		clientModule.APIKeyMiddleware,
		jwtMiddleware,
	)

	usageModule.Register(
		app,
		db,
		rdb,
		clientModule.APIKeyMiddleware,
		jwtMiddleware,
		ipWhiteListModule.IPWhitelistMiddleware,
	)

	// Health Check
	app.Get("/health", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{"status": "ok"})
	})

	go startServer(app)

	waitForShutdown(ctx, app)
}

func startServer(app *fiber.App) {
	port := shared.GetEnv("APP_PORT", "8080")
	log.Printf("server running on :%s", port)

	if err := app.Listen(":" + port); err != nil {
		log.Printf("fiber stopped: %v", err)
	}
}

func waitForShutdown(ctx context.Context, app *fiber.App) {
	ch := make(chan os.Signal, 1)
	signal.Notify(ch, syscall.SIGINT, syscall.SIGTERM)

	<-ch
	log.Println("shutting down...")

	ctx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	if err := app.ShutdownWithContext(ctx); err != nil {
		log.Printf("shutdown error: %v", err)
	}
}
