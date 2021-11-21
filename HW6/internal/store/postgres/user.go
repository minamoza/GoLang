package postgres

import (
	"context"
	"github.com/jmoiron/sqlx"
	"go/internal/models"
	"go/internal/store"
)

func (db *DB) User() store.UserRep {
	if db.users == nil {
		db.users = NewUserRepository(db.conn)
	}

	return db.users
}

type UserRepository struct {
	conn *sqlx.DB
}

func NewUserRepository(conn *sqlx.DB) store.UserRep {
	return &UserRepository{conn: conn}
}

func (c UserRepository) Create(ctx context.Context, User *models.User) error {
	_, err := c.conn.Exec("INSERT INTO User(name) VALUES ($1)", User.Name)
	if err != nil {
		return err
	}

	return nil
}

func (c UserRepository) All(ctx context.Context) ([]*models.User, error) {
	User := make([]*models.User, 0)
	if err := c.conn.Select(&User, "SELECT * FROM User"); err != nil {
		return nil, err
	}

	return User, nil
}

func (c UserRepository) ByID(ctx context.Context, id int) (*models.User, error) {
	User := new(models.User)
	if err := c.conn.Get(User, "SELECT id, name FROM User WHERE id=$1", id); err != nil {
		return nil, err
	}

	return User, nil
}

func (c UserRepository) Update(ctx context.Context, User *models.User) error {
	_, err := c.conn.Exec("UPDATE User SET name = $1 WHERE id = $2", User.Name, User.ID)
	if err != nil {
		return err
	}

	return nil
}

func (c UserRepository) Delete(ctx context.Context, id int) error {
	_, err := c.conn.Exec("DELETE FROM User WHERE id = $1", id)
	if err != nil {
		return err
	}

	return nil
}
