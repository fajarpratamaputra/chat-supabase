package chats

import (
	"chitchat/domain/chats"
	"chitchat/domain/chats/models"
	http2 "chitchat/infrastructure/http"
	"context"
)

type service struct {
	chatRepo   chats.Repository
	httpClient *http2.Client
}

func NewChatService(repository chats.Repository, httpClient *http2.Client) chats.Service {
	return &service{
		chatRepo:   repository,
		httpClient: httpClient,
	}
}

func (s *service) CreateUser(ctx context.Context, chat models.Chat) error {

	err := s.chatRepo.InsertChat(ctx, chat)
	if err != nil {
		return err
	}

	return nil
}
