package postgres

import (
	"context"
	"github.com/jmoiron/sqlx"
	"go/internal/models"
	"go/internal/store"
)

func (db *DB) Restaurant() store.RestaurantRep {
	if db.restaurants == nil {
		db.restaurants = NewRestaurantRepository(db.conn)
	}

	return db.restaurants
}

type RestaurantRepository struct {
	conn *sqlx.DB
}

func NewRestaurantRepository(conn *sqlx.DB) store.RestaurantRep {
	return &RestaurantRepository{conn: conn}
}

func (c RestaurantRepository) Create(ctx context.Context, Restaurant *models.Restaurant) error {
	_, err := c.conn.Exec("INSERT INTO restaurant(name) VALUES ($1)", Restaurant.Name)
	if err != nil {
		return err
	}

	return nil
}

func (c RestaurantRepository) All(ctx context.Context) ([]*models.Restaurant, error) {
	Restaurant := make([]*models.Restaurant, 0)
	if err := c.conn.Select(&Restaurant, "SELECT * FROM restaurant"); err != nil {
		return nil, err
	}

	return Restaurant, nil
}

func (c RestaurantRepository) ByID(ctx context.Context, id int) (*models.Restaurant, error) {
	Restaurant := new(models.Restaurant)
	if err := c.conn.Get(Restaurant, "SELECT id, name FROM restaurant WHERE id=$1", id); err != nil {
		return nil, err
	}

	return Restaurant, nil
}

func (c RestaurantRepository) Update(ctx context.Context, Restaurant *models.Restaurant) error {
	_, err := c.conn.Exec("UPDATE restaurant SET name = $1 WHERE id = $2", Restaurant.Name, Restaurant.ID)
	if err != nil {
		return err
	}

	return nil
}

func (c RestaurantRepository) Delete(ctx context.Context, id int) error {
	_, err := c.conn.Exec("DELETE FROM restaurant WHERE id = $1", id)
	if err != nil {
		return err
	}

	return nil
}