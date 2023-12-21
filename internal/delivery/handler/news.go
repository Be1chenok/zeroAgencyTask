package handler

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"

	"github.com/Be1chenok/zeroAgencyTask/internal/domain"
	"github.com/gofiber/fiber/v2"
)

const (
	defaultPage = 1
	defaultSize = 10
)

func (h Handler) GetNewsList(c *fiber.Ctx) error {
	ctx, cancel := context.WithTimeout(c.Context(), h.conf.Server.RequestTimeout)
	defer cancel()

	page := c.QueryInt("page", defaultPage)
	size := c.QueryInt("size", defaultSize)

	searchParams := domain.NewsSearchParams{
		Offset: (page - 1) * size,
		Limit:  size,
	}

	news, err := h.service.FindNews(ctx, searchParams)
	if err != nil {
		if errors.Is(err, domain.ErrNothingFound) {
			return c.Status(fiber.StatusBadRequest).JSON(domain.ErrNothingFound.Error())
		}
		return c.Status(fiber.StatusInternalServerError).JSON(ErrSomethingWentWrong.Error())
	}

	return c.Status(fiber.StatusOK).JSON(&news)
}

func (h Handler) EditNewsById(c *fiber.Ctx) error {
	ctx, cancel := context.WithTimeout(c.Context(), h.conf.Server.RequestTimeout)
	defer cancel()

	id, err := c.ParamsInt("id", 0)
	if err != nil || id < 1 {
		return c.Status(fiber.StatusBadRequest).JSON(ErrInvalidInput.Error())
	}

	body := c.Body()

	var input domain.News
	if err := json.Unmarshal(body, &input); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(ErrInvalidInput.Error())
	}

	input.Id = id

	if err := h.service.UpdateNewsById(ctx, &input); err != nil {
		if errors.Is(err, domain.ErrNothingUpdated) {
			return c.Status(fiber.StatusBadRequest).JSON(domain.ErrNothingUpdated.Error())
		}
		return c.Status(fiber.StatusInternalServerError).JSON(ErrSomethingWentWrong.Error())

	}

	return c.Status(fiber.StatusNoContent).JSON(fmt.Sprintf("updated news by id: %d", id))
}
