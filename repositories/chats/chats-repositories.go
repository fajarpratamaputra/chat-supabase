package chats

import (
	"chitchat/config"
	"chitchat/domain/chats"
	"chitchat/domain/chats/models"
	"chitchat/infrastructure/supabase"
	"context"
	"fmt"
	"github.com/afex/hystrix-go/hystrix"
	"time"
)

type repository struct {
	supabase *supabase.Client
}

func NewChatsRepository(client *supabase.Client) chats.Repository {
	return &repository{
		supabase: client,
	}
}

func (r *repository) InsertChat(ctx context.Context, Chat models.Chat) error {

	hystrix.ConfigureCommand("InsertUser", hystrix.CommandConfig{
		Timeout:               config.Cfg.GetInt("CB_TIMEOUT"),
		MaxConcurrentRequests: config.Cfg.GetInt("CB_MAX_CONCURRENT"),
		ErrorPercentThreshold: config.Cfg.GetInt("CB_ERROR_PERCENT_THRESHOLD"),
	})
	_, err := time.LoadLocation("Asia/Jakarta")
	if err != nil {
		fmt.Println("Error loading Jakarta time zone:", err)
		return err
	}

	// Get the current time in Jakarta
	row := models.Chat{
		ID:           Chat.ID,
		UserID:       Chat.UserID,
		UserIdTarget: Chat.UserIdTarget,
		Message:      Chat.Message,
		CreatedAt:    Chat.CreatedAt,
	}

	var results []models.Chat
	err = supabase.NewClient().DB.From("chat").Insert(row).Execute(&results)
	if err != nil {
		panic(err)
	}

	return nil
}
