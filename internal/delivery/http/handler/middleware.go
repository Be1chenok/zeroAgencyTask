package handler

import (
	"context"
	"errors"
	"strings"

	"github.com/Be1chenok/zeroAgencyTask/internal/domain"
	"github.com/gofiber/fiber/v2"
)

const authorizationHeader = "Authorization"

func (h Handler) UserAccessIdentity(c *fiber.Ctx) error {
	ctx, cancel := context.WithTimeout(c.Context(), h.conf.Server.RequestTimeout)
	defer cancel()
	header := c.Get(authorizationHeader)
	if header == "" {
		return c.Status(fiber.StatusUnauthorized).JSON(ErrEmptyHeader.Error())
	}
	headerParts := strings.Split(header, " ")
	if len(headerParts) != 2 || headerParts[0] != "Bearer" {
		return c.Status(fiber.StatusUnauthorized).JSON(ErrInvalidAuthHeader.Error())
	}
	if len(headerParts[1]) == 0 {
		return c.Status(fiber.StatusUnauthorized).JSON(ErrEmptyToken.Error())
	}

	userId, err := h.service.ParseToken(ctx, headerParts[1])
	if err != nil {
		if errors.Is(err, domain.ErrNothingFound) {
			return c.Status(fiber.StatusUnauthorized).JSON(ErrEmptyToken.Error())
		}
		return c.Status(fiber.StatusInternalServerError).JSON(ErrSomethingWentWrong)
	}

	c.Locals("userId", userId)

	err = c.Next()

	return err
}
