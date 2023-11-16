package chats

import (
	"chitchat/domain/chats/models"
	"context"
)

type Service interface {
	CreateUser(ctx context.Context, chat models.Chat) error
}

type Repository interface {
	InsertChat(ctx context.Context, Chat models.Chat) error
}
