package store

import (
	"context"
	"lectures-6/internal/models"
)

type Store interface {
	Create(ctx context.Context, laptop *models.Order) error
	All(ctx context.Context) ([]*models.Order, error)
	ByID(ctx context.Context, id int) (*models.Order, error)
	Update(ctx context.Context, laptop *models.Order) error
	Delete(ctx context.Context, id int) error
}