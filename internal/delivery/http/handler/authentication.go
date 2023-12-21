package handler

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"strings"

	"github.com/Be1chenok/zeroAgencyTask/internal/domain"
	"github.com/gofiber/fiber/v2"
)

func (h Handler) Register(c *fiber.Ctx) error {
	ctx, cancel := context.WithTimeout(c.Context(), h.conf.Server.RequestTimeout)
	defer cancel()

	body := c.Body()

	var input domain.User
	if err := json.Unmarshal(body, &input); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(ErrInvalidInput.Error())
	}

	userId, err := h.service.SignUp(ctx, input)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(ErrSomethingWentWrong.Error())
	}

	return c.Status(fiber.StatusOK).JSON(fmt.Sprintf("user id: %d", userId))
}

func (h Handler) Login(c *fiber.Ctx) error {
	ctx, cancel := context.WithTimeout(c.Context(), h.conf.Server.RequestTimeout)
	defer cancel()

	body := c.Body()
	var input domain.SignInInput
	if err := json.Unmarshal(body, &input); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(ErrInvalidInput.Error())
	}

	tokens, err := h.service.Authentication.SignIn(ctx, input)
	if err != nil {
		if errors.Is(err, domain.ErrNothingFound) {
			return c.Status(fiber.StatusBadRequest).JSON(ErrInvalidInput.Error())
		}
		return c.Status(fiber.StatusInternalServerError).JSON(ErrSomethingWentWrong.Error())
	}

	return c.Status(fiber.StatusOK).JSON(tokens)
}

func (h Handler) LogOut(c *fiber.Ctx) error {
	ctx, cancel := context.WithTimeout(c.Context(), h.conf.Server.RequestTimeout)
	defer cancel()

	header := c.Get(authorizationHeader)
	headerParts := strings.Split(header, " ")

	accessToken := headerParts[1]

	if err := h.service.Authentication.SignOut(ctx, accessToken); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(ErrSomethingWentWrong.Error())
	}

	return c.Status(fiber.StatusOK).JSON(nil)
}

func (h Handler) FullLogOut(c *fiber.Ctx) error {
	ctx, cancel := context.WithTimeout(c.Context(), h.conf.Server.RequestTimeout)
	defer cancel()

	header := c.Get(authorizationHeader)
	headerParts := strings.Split(header, " ")
	accessToken := headerParts[1]

	if err := h.service.FullSignOut(ctx, accessToken); err != nil {
		if errors.Is(err, domain.ErrNothingFound) {
			return c.Status(fiber.StatusBadRequest).JSON(ErrInvalidInput.Error())
		}
		return c.Status(fiber.StatusInternalServerError).JSON(ErrSomethingWentWrong.Error())
	}

	return c.Status(fiber.StatusOK).JSON(nil)
}

func (h Handler) Refresh(c *fiber.Ctx) error {
	ctx, cancel := context.WithTimeout(c.Context(), h.conf.Server.RequestTimeout)
	defer cancel()

	header := c.Get(authorizationHeader)
	headerParts := strings.Split(header, " ")
	refreshToken := headerParts[1]

	tokens, err := h.service.RefreshTokens(ctx, refreshToken)
	if err != nil {
		if errors.Is(err, domain.ErrNothingFound) {
			return c.Status(fiber.StatusBadRequest).JSON(ErrInvalidInput.Error())
		}
		return c.Status(fiber.StatusInternalServerError).JSON(ErrSomethingWentWrong.Error())
	}

	return c.Status(fiber.StatusOK).JSON(tokens)
}
