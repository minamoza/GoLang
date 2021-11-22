package redis_cache

import (
	"context"
	"encoding/json"
	"github.com/go-redis/redis/v8"
	"go/internal/models"
	"time"
)

func (rc *RedisCache) Dish() DishCacheRepository {
	if rc.dishes == nil {
		rc.dishes = newDishRepo(rc.client, rc.expires)
	}

	return rc.dishes
}

type DishRepo struct {
	client  *redis.Client
	expires time.Duration
}

func newDishRepo(client *redis.Client, exp time.Duration) DishCacheRepository {
	return &DishRepo{client: client, expires: exp}
}

func (c *DishRepo) Set(ctx context.Context, key string, value []*models.Dish) error {
	dishBytes, err := json.Marshal(value)
	if err != nil {
		return err
	}

	if err = c.client.Set(ctx, key, dishBytes, c.expires*time.Second).Err(); err != nil {
		return err
	}

	return nil
}

func (c *DishRepo) Get(ctx context.Context, key string) ([]*models.Dish, error) {
	result, err := c.client.Get(ctx, key).Result()
	switch err {
	case nil:
		break
	case redis.Nil:
		return nil, nil
	default:
		return nil, err
	}

	Dish := make([]*models.Dish, 0)
	if err = json.Unmarshal([]byte(result), &Dish); err != nil {
		return nil, err
	}

	return Dish, nil
}