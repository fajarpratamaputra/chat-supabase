package user

import (
	"chitchat/domain/users"
	"chitchat/domain/users/models"
	"chitchat/helpers"
	"github.com/gofiber/fiber/v2"
	"math/rand"
	"time"
)

type HandlerUser struct {
	service users.Service
}

func NewUserHandler(svc users.Service) *HandlerUser {
	return &HandlerUser{
		service: svc,
	}
}

func (h *HandlerUser) UserCreate(c *fiber.Ctx) error {

	var request models.User
	ctx := c.Context()
	request.ID = randomInt(10, 20)
	request.Name = c.FormValue("name")
	request.Username = c.FormValue("username")
	request.Password = c.FormValue("password")

	err := h.service.CreateUser(ctx, request)
	if err != nil {
		return c.Status(fiber.StatusOK).JSON(helpers.ResponseSuccess("2", err.Error(), nil))
	}

	return c.Status(fiber.StatusOK).JSON(helpers.ResponseSuccess("0", "Success", nil))
}

func randomInt(min, max int) int {
	rand.Seed(time.Now().UnixNano()) // Seed the random number generator with current time
	return rand.Intn(max-min+1) + min
}
