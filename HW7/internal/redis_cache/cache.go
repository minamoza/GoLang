package redis_cache

import (
	"context"
	"go/internal/models"
)

type Cache interface {
	Close() error

	Dish() DishCacheRepository

	DeleteAll(ctx context.Context) error
}

type DishCacheRepository interface {
	Set(ctx context.Context, key string, value []*models.Dish) error
	Get(ctx context.Context, key string) ([]*models.Dish, error)
}