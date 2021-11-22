package postgres

import (
	"context"
	"fmt"
	"github.com/jmoiron/sqlx"
	"github.com/lib/pq"
	"go/internal/models"
	"go/internal/store"
)

func (db *DB) Dish() store.DishRep {
	if db.dishes == nil {
		db.dishes = NewDishRepository(db.conn)
	}

	return db.dishes
}

type DishRepository struct {
	conn *sqlx.DB
}

func NewDishRepository(conn *sqlx.DB) store.DishRep {
	return &DishRepository{conn: conn}
}

func (c DishRepository) Create(ctx context.Context, Dish *models.Dish) error {
	_, err := c.conn.Exec("INSERT INTO Dish(dish, cost, ingredients) VALUES ($1, $2, $3)", Dish.Dish, Dish.Cost, pq.Array(Dish.Ingredients))
	if err != nil {
		return err
	}

	return nil
}

func (c DishRepository) All(ctx context.Context, filter *models.DishFilter) ([]*models.Dish, error) {
	dishes := make([]*models.Dish, 0)
	basicQuery := "SELECT id, dish, cost FROM dish"

	if filter.Query != nil {
		basicQuery = fmt.Sprintf("%s WHERE name ILIKE $1", basicQuery)

		if err := c.conn.Select(&dishes, basicQuery, "%"+*filter.Query+"%"); err != nil {
			return nil, err
		}

		return dishes, nil
	}

	if err := c.conn.Select(&dishes, basicQuery); err != nil {
		return nil, err
	}

	return dishes, nil
}

func (c DishRepository) ByID(ctx context.Context, id int) (*models.Dish, error) {
	Dish := new(models.Dish)
	if err := c.conn.Get(Dish, "SELECT id, dish, cost FROM Dish WHERE id=$1", id); err != nil {
		return nil, err
	}

	return Dish, nil
}

func (c DishRepository) Update(ctx context.Context, Dish *models.Dish) error {
	_, err := c.conn.Exec("UPDATE dish SET dish=$1, cost=$2 WHERE id=$3", Dish.Dish, Dish.Cost, Dish.ID)
	if err != nil {
		return err
	}

	return nil
}

func (c DishRepository) Delete(ctx context.Context, id int) error {
	_, err := c.conn.Exec("DELETE FROM Dish WHERE id=$1", id)
	if err != nil {
		return err
	}

	return nil
}