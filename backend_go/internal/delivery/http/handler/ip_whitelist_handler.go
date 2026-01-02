package handler

import (
	"github.com/gofiber/fiber/v2"
	"github.com/hafisrabbani/technical-test-nexmedis/internal/model/request"

	"github.com/hafisrabbani/technical-test-nexmedis/internal/model/response"
	"github.com/hafisrabbani/technical-test-nexmedis/internal/repository"
	"github.com/hafisrabbani/technical-test-nexmedis/internal/service"
)

type IPWhitelistHandler struct {
	service *service.IPWhitelistService
}

func NewIPWhitelistHandler(s *service.IPWhitelistService) *IPWhitelistHandler {
	return &IPWhitelistHandler{service: s}
}

func (h *IPWhitelistHandler) Register(c *fiber.Ctx) error {
	client := c.Locals("client").(*repository.Client)

	var req request.IPWhitelistRequest
	if err := c.BodyParser(&req); err != nil {
		return fiber.ErrBadRequest
	}

	if err := h.service.ReplaceAll(
		c.Context(),
		client.ID,
		req.IPs,
	); err != nil {
		return fiber.ErrInternalServerError
	}

	return c.JSON(response.Ok())
}
