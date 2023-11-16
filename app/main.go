package app

import (
	"chitchat/infrastructure/http"
	"chitchat/infrastructure/postgresql"
	"chitchat/infrastructure/redis"
	"chitchat/infrastructure/sql"
	"chitchat/infrastructure/supabase"
	"context"
)

type Container struct {
	Cache            *redis.Client
	HttpClient       *http.Client
	SqlClient        *sql.Client
	PostgreSqlClient *postgresql.Client
	SupabaseClient   *supabase.Client
}

func NewContainer(ctx context.Context) *Container {
	c := redis.NewClient()
	hc := http.NewClient()
	sc := sql.NewClient()
	psc := postgresql.NewClient()

	return &Container{
		Cache:            c,
		HttpClient:       hc,
		SqlClient:        sc,
		PostgreSqlClient: psc,
	}
}
