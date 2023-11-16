package models

import "time"

type Chat struct {
	ID           int       `json:"id"`
	UserID       int       `json:"user_id"`
	UserIdTarget int       `json:"user_id_target"`
	Message      string    `json:"message"`
	CreatedAt    time.Time `json:"created_at"`
}
