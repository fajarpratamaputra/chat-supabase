package checker

import "context"

type Service interface {
	HealthCheck(ctx context.Context) (map[string]interface{}, error)
}

type Repository interface {
}
