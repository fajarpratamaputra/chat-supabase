package users

import (
	"chitchat/domain/users"
	"chitchat/domain/users/models"
	http2 "chitchat/infrastructure/http"
	"context"
)

type service struct {
	userRepo   users.Repository
	httpClient *http2.Client
}

func NewUserService(repository users.Repository, httpClient *http2.Client) users.Service {
	return &service{
		userRepo:   repository,
		httpClient: httpClient,
	}
}

func (s *service) CreateUser(ctx context.Context, user models.User) error {

	err := s.userRepo.InsertUser(ctx, user)
	if err != nil {
		return err
	}

	return nil
}
