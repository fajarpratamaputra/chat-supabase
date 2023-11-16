package checker

import (
	"chitchat/domain/checker"
	"chitchat/infrastructure/http"
	"chitchat/infrastructure/postgresql"
	"chitchat/infrastructure/redis"
	"chitchat/infrastructure/sql"
	"context"
)

type service struct {
	redis      *redis.Client
	sql        *sql.Client
	http       *http.Client
	postgresql *postgresql.Client
}

func NewCheckerService(redisC *redis.Client, sqlC *sql.Client, httpC *http.Client, postgresqlC *postgresql.Client) checker.Service {
	return &service{
		redis:      redisC,
		sql:        sqlC,
		http:       httpC,
		postgresql: postgresqlC,
	}
}

func (s *service) HealthCheck(ctx context.Context) (map[string]interface{}, error) {

	res := map[string]interface{}{
		"redis_status":      s.checkRedis(ctx, s.redis),
		"mysql_status":      s.checkMysql(ctx),
		"postgresql_status": s.checkPostgresql(ctx),
	}

	return res, nil
}

func (s *service) checkRedis(ctx context.Context, client *redis.Client) bool {

	if _, err := client.Set(ctx, "health_check", "OK", 3600).Result(); err != nil {
		return false
	}

	if _, err := client.Get(ctx, "health_check").Result(); err != nil {
		return false
	}

	return true
}

func (s *service) checkMysql(ctx context.Context) bool {
	if s.sql == nil {
		return false
	}

	db, err := s.sql.WithContext(ctx).DB()
	if err = db.Ping(); err != nil {
		return false
	}

	return true
}

func (s *service) checkPostgresql(ctx context.Context) bool {
	if s.sql == nil {
		return false
	}

	db, err := s.postgresql.WithContext(ctx).DB()
	if err = db.Ping(); err != nil {
		return false
	}

	return true
}
