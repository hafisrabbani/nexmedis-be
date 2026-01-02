package handler

import (
	"errors"

	"github.com/gofiber/fiber/v2"

	domainErr "github.com/hafisrabbani/technical-test-nexmedis/internal/model/error"
	"github.com/hafisrabbani/technical-test-nexmedis/internal/model/request"
	"github.com/hafisrabbani/technical-test-nexmedis/internal/model/response"
	"github.com/hafisrabbani/technical-test-nexmedis/internal/service"
)

type ClientHandler struct {
	service *service.ClientService
}

func NewClientHandler(service *service.ClientService) *ClientHandler {
	return &ClientHandler{service: service}
}

func (h *ClientHandler) Register(c *fiber.Ctx) error {
	var req request.RegisterRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).
			JSON(response.Error("invalid request body"))
	}

	if req.ClientID == "" || req.Name == "" || req.Email == "" {
		return c.Status(fiber.StatusBadRequest).
			JSON(response.Error("missing required fields"))
	}

	apiKey, err := h.service.Register(
		c.Context(),
		req.ClientID,
		req.Name,
		req.Email,
	)
	if err != nil {
		if errors.Is(err, domainErr.ErrClientAlreadyExists) {
			return c.Status(fiber.StatusBadRequest).
				JSON(response.Error("client already exists"))
		}

		return c.Status(fiber.StatusInternalServerError).
			JSON(response.Error("failed to register client"))
	}

	return c.JSON(
		response.Success(response.RegisterResponse{
			APIKey: apiKey,
		}),
	)
}
