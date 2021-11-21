package store

import (
	"context"
	"go/internal/models"
)

type Store interface{
	Connect(url string) error
	Close() error

	Order() OrderRep
	Categories() CategoriesRep
	Restaurant() RestaurantRep
	User() UserRep
}

type OrderRep interface {
	Create(ctx context.Context, order *models.Order) error
	All(ctx context.Context) ([]*models.Order, error)
	ByID(ctx context.Context, id int) (*models.Order, error)
	Update(ctx context.Context, order *models.Order) error
	Delete(ctx context.Context, id int) error
}

type RestaurantRep interface{
	Create(ctx context.Context, res *models.Restaurant) error
	All(ctx context.Context) ([]*models.Restaurant, error)
	ByID(ctx context.Context, id int) (*models.Restaurant, error)
	Update(ctx context.Context, res *models.Restaurant) error
	Delete(ctx context.Context, id int) error
}

type UserRep interface{
	Create(ctx context.Context, res *models.User) error
	All(ctx context.Context) ([]*models.User, error)
	ByID(ctx context.Context, id int) (*models.User, error)
	Update(ctx context.Context, user *models.User) error
	Delete(ctx context.Context, id int) error
}

type CategoriesRep interface {
	Create(ctx context.Context, category *models.Category) error
	All(ctx context.Context) ([]*models.Category, error)
	ByID(ctx context.Context, id int) (*models.Category, error)
	Update(ctx context.Context, category *models.Category) error
	Delete(ctx context.Context, id int) error
}