package postgres

import (
	_ "github.com/jackc/pgx/stdlib"
  "github.com/jmoiron/sqlx"
	"go/internal/store"
)

type DB struct {
	conn *sqlx.DB

	users 		store.UserRep
	orders      store.OrderRep
	restaurants store.RestaurantRep
	categories  store.CategoriesRep
}

func NewDB() store.Store {
	return &DB{}
}

func (db *DB) Connect(url string) error {
	conn, err := sqlx.Connect("pgx", url)
	if err != nil {
		return err
	}

	if err := conn.Ping(); err != nil {
		return err
	}

	db.conn = conn
	return nil
}

func (db *DB) Close() error {
	return db.conn.Close()
}

