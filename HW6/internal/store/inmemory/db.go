package inmemory

import (
	"context"
	"fmt"
	"lectures-6/internal/models"
	"lectures-6/internal/store"
	"sync"
)

type DB struct {
	data map[int]*models.Order

	mu *sync.RWMutex
}

func NewDB() store.Store {
	return &DB{
		data: make(map[int]*models.Order),
		mu:   new(sync.RWMutex),
	}
}

func (db *DB) Create(ctx context.Context, order *models.Order) error {
	db.mu.Lock()
	defer db.mu.Unlock()

	db.data[order.ID] = order
	return nil
}

func (db *DB) All(ctx context.Context) ([]*models.Order, error) {
	db.mu.RLock()
	defer db.mu.RUnlock()

	orders := make([]*models.Order, 0, len(db.data))
	for _, order := range db.data {
		orders = append(orders, order)
	}

	return orders, nil
}

func (db *DB) ByID(ctx context.Context, id int) (*models.Order, error) {
	db.mu.RLock()
	defer db.mu.RUnlock()

	order, ok := db.data[id]
	if !ok {
		return nil, fmt.Errorf("No order with id %d", id)
	}

	return order, nil
}

func (db *DB) Update(ctx context.Context, order *models.Order) error {
	db.mu.Lock()
	defer db.mu.Unlock()

	db.data[order.ID] = order
	return nil
}

func (db *DB) Delete(ctx context.Context, id int) error {
	db.mu.Lock()
	defer db.mu.Unlock()

	delete(db.data, id)
	return nil
}
