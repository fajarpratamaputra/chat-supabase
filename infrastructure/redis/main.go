package redis

import (
	"chitchat/config"
	"github.com/redis/go-redis/v9"
)

type Client struct {
	*redis.Client
}

func NewClient() *Client {
	cl := redis.NewClient(&redis.Options{
		Addr: config.Cfg.GetString("REDIS_HOST"),
		DB:   config.Cfg.GetInt("REDIS_DB"),
	})

	return &Client{cl}
}
