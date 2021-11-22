package postgres

import (
	"context"
	"github.com/jmoiron/sqlx"
	"go/internal/models"
	"go/internal/store"
)

func (db *DB) Order() store.OrderRep {
	if db.orders == nil {
		db.orders = NewOrderRepository(db.conn)
	}

	return db.orders
}

type OrderRepository struct {
	conn *sqlx.DB
}

func NewOrderRepository(conn *sqlx.DB) store.OrderRep {
	return &OrderRepository{conn: conn}
}

func (c OrderRepository) Create(ctx context.Context, Order *models.Order) error {
	_, err := c.conn.Exec("INSERT INTO order(order) VALUES ($1)", Order.Order)
	if err != nil {
		return err
	}

	return nil
}

func (c OrderRepository) All(ctx context.Context) ([]*models.Order, error) {
	Order := make([]*models.Order, 0)
	if err := c.conn.Select(&Order, "SELECT * FROM order"); err != nil {
		return nil, err
	}

	return Order, nil
}

func (c OrderRepository) ByID(ctx context.Context, id int) (*models.Order, error) {
	Order := new(models.Order)
	if err := c.conn.Get(Order, "SELECT * FROM order WHERE id=$1", id); err != nil {
		return nil, err
	}

	return Order, nil
}

func (c OrderRepository) Update(ctx context.Context, Order *models.Order) error {
	_, err := c.conn.Exec("UPDATE order SET order = $1 WHERE id = $2", Order.Name, Order.ID)
	if err != nil {
		return err
	}

	return nil
}

func (c OrderRepository) Delete(ctx context.Context, id int) error {
	_, err := c.conn.Exec("DELETE FROM order WHERE id = $1", id)
	if err != nil {
		return err
	}

	return nil
}