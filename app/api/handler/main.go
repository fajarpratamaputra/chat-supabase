package handler

import (
	"chitchat/app"
	"chitchat/app/api/handler/chat"
	"chitchat/app/api/handler/checker"
	"chitchat/app/api/handler/user"
	chatRepo "chitchat/repositories/chats"
	userRepo "chitchat/repositories/users"
	chatSvc "chitchat/services/chats"
	checkerSvc "chitchat/services/checker"
	userSvc "chitchat/services/users"
	"github.com/gofiber/fiber/v2"
	jwtFiber "github.com/gofiber/jwt/v3"
	"github.com/golang-jwt/jwt/v4"
	"net/http"
	"strings"
)

func loadCheckerRoute(c *app.Container, router *fiber.App) *fiber.App {
	checkSvc := checkerSvc.NewCheckerService(c.Cache, c.SqlClient, c.HttpClient, c.PostgreSqlClient)
	handler := checker.NewCheckerHandler(checkSvc)

	router.Get("/health-check", handler.HealthCheck)
	return router
}
func LoadHandler(router *fiber.App, c *app.Container) *fiber.App {
	//router without jwt
	router = loadCheckerRoute(c, router)
	router = loadUserRoute(c, router)
	router = loadChatRoute(c, router)

	//using jwt
	router.Use(setBearerRule)
	router.Use(jwtFiber.New(jwtFiber.Config{
		SigningMethod: "HS256",
		SigningKey:    []byte("secret"),
	}))
	router.Use(validateJWTClient)
	return router
}

func loadUserRoute(c *app.Container, router *fiber.App) *fiber.App {
	userRepo := userRepo.NewUsersRepository(c.SupabaseClient)
	userSvc := userSvc.NewUserService(userRepo, c.HttpClient)
	userHandler := user.NewUserHandler(userSvc)

	mscGrp := router.Group("/user")
	mscGrp.Post("/create", userHandler.UserCreate)
	return router
}

func loadChatRoute(c *app.Container, router *fiber.App) *fiber.App {
	chatRepo := chatRepo.NewChatsRepository(c.SupabaseClient)
	chatSvc := chatSvc.NewChatService(chatRepo, c.HttpClient)
	chatHandler := chat.NewChatHandler(chatSvc)

	mscGrp := router.Group("/member")
	mscGrp.Post("/chat", chatHandler.UserCreate)
	return router
}

func setBearerRule(c *fiber.Ctx) error {
	tokenString := c.Get("Authorization")
	if strings.HasPrefix(tokenString, "Bearer") == false {
		c.Request().Header.Set("Authorization", "Bearer "+tokenString)
	}

	return c.Next()
}

func validateJWTClient(c *fiber.Ctx) error {
	user := c.Locals("user").(*jwt.Token)
	if claims, ok := user.Claims.(jwt.MapClaims); ok {
		c.Locals("user_data", map[string]interface{}(claims))
		return c.Next()
	}

	return c.SendStatus(http.StatusUnauthorized)
}
