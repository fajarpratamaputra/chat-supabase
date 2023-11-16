package users

import (
	"chitchat/config"
	"chitchat/domain/users"
	"chitchat/domain/users/models"
	"chitchat/infrastructure/supabase"
	"context"
	"fmt"
	"github.com/afex/hystrix-go/hystrix"
	"time"
)

type repository struct {
	supabase *supabase.Client
}

func NewUsersRepository(client *supabase.Client) users.Repository {
	return &repository{
		supabase: client,
	}
}

func (r *repository) InsertUser(ctx context.Context, User models.User) error {

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
	row := models.User{
		ID:       User.ID,
		Name:     User.Name,
		Username: User.Username,
		Password: User.Password,
	}

	var results []models.User
	err = supabase.NewClient().DB.From("member").Insert(row).Execute(&results)
	if err != nil {
		panic(err)
	}

	return nil
}
