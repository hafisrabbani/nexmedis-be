package handler

import (
	"github.com/gofiber/fiber/v2"
	"github.com/hafisrabbani/technical-test-nexmedis/internal/repository"

	"github.com/hafisrabbani/technical-test-nexmedis/internal/model/response"
	"github.com/hafisrabbani/technical-test-nexmedis/internal/service"
)

type TokenHandler struct {
	jwt *service.JWTService
}

func NewTokenHandler(jwt *service.JWTService) *TokenHandler {
	return &TokenHandler{jwt: jwt}
}

func (h *TokenHandler) Issue(c *fiber.Ctx) error {
	client := c.Locals("client").(*repository.Client)

	token, err := h.jwt.Generate(client.ClientID)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(response.Error("failed to generate token"))
	}

	return c.JSON(response.Success(fiber.Map{
		"token": token,
	}))
}
