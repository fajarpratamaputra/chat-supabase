package chat

import (
	"chitchat/domain/chats"
	"chitchat/domain/chats/models"
	"chitchat/helpers"
	"github.com/gofiber/fiber/v2"
	"math/rand"
	"strconv"
	"time"
)

type HandlerUser struct {
	service chats.Service
}

func NewChatHandler(svc chats.Service) *HandlerUser {
	return &HandlerUser{
		service: svc,
	}
}

func (h *HandlerUser) UserCreate(c *fiber.Ctx) error {

	var request models.Chat
	ctx := c.Context()
	request.ID = randomInt(10, 20)
	request.UserID, _ = strconv.Atoi(c.FormValue("user_id"))
	request.UserIdTarget, _ = strconv.Atoi(c.FormValue("user_id_target"))
	request.Message = c.FormValue("message")
	request.CreatedAt = time.Now()

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
