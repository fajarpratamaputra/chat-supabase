package users

import (
	"chitchat/domain/users/models"
	"context"
)

type Service interface {
	CreateUser(ctx context.Context, user models.User) error
}

type Repository interface {
	InsertUser(ctx context.Context, User models.User) error
}
