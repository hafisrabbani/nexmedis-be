package handler

import (
	"bufio"
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/hafisrabbani/technical-test-nexmedis/internal/repository"

	"github.com/hafisrabbani/technical-test-nexmedis/internal/model/response"
	"github.com/hafisrabbani/technical-test-nexmedis/internal/service"
)

type UsageHandler struct {
	service *service.UsageService
}

func NewUsageHandler(service *service.UsageService) *UsageHandler {
	return &UsageHandler{service: service}
}

func (h *UsageHandler) Log(c *fiber.Ctx) error {
	client := c.Locals("client").(*repository.Client)

	h.service.Log(c.Context(), client.ClientID)

	return c.JSON(response.Ok())
}

func (h *UsageHandler) Daily(c *fiber.Ctx) error {
	client := c.Locals("client").(*repository.Client)

	data, err := h.service.Daily(c.Context(), client.ClientID)
	if err != nil {
		return c.Status(500).JSON(response.Error("failed to fetch usage"))
	}

	return c.JSON(response.Success(data))
}

func (h *UsageHandler) RealtimeDailyUsage(c *fiber.Ctx) error {
	client := c.Locals("client").(*repository.Client)

	c.Set("Content-Type", "text/event-stream")
	c.Set("Cache-Control", "no-cache")
	c.Set("Connection", "keep-alive")

	c.Context().SetBodyStreamWriter(func(w *bufio.Writer) {
		for {
			data, err := h.service.Daily(context.Background(), client.ClientID)
			if err != nil {
				_, err = fmt.Fprintf(w, "data: %s\n\n", `{"error":"failed to fetch usage"}`)
				if err != nil {
					return
				}
				if err = w.Flush(); err != nil {
					return
				}
				time.Sleep(2 * time.Second)
				continue
			}

			payload, _ := json.Marshal(response.Success(data))

			_, err = fmt.Fprintf(w, "data: %s\n\n", payload)
			if err != nil {
				return
			}

			if err = w.Flush(); err != nil {
				return
			}

			time.Sleep(1 * time.Second)
		}
	})

	return nil
}

func (h *UsageHandler) Top(c *fiber.Ctx) error {
	data, err := h.service.Top(c.Context())
	if err != nil {
		return c.Status(500).JSON(response.Error("failed to fetch top usage"))
	}

	return c.JSON(response.Success(data))
}
