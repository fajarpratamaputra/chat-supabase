package checker

import (
	"chitchat/domain/checker"
	"github.com/gofiber/fiber/v2"
)

type HandlerChecker struct {
	service checker.Service
}

func NewCheckerHandler(svc checker.Service) *HandlerChecker {
	return &HandlerChecker{
		service: svc,
	}
}

func (h *HandlerChecker) HealthCheck(c *fiber.Ctx) error {
	ctx := c.Context()
	res, err := h.service.HealthCheck(ctx)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(res)
}
