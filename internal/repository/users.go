package repository

import (
	"context"
	"database/sql"
)

type UserRepository struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) *UserRepository {
	return &UserRepository{db: db}
}

func (r *UserRepository) CreateUser(ctx context.Context, email, passwordHash string) error {
	_, err := r.db.ExecContext(ctx,
		"INSERT INTO users (email, password_hash) VALUES ($1, $2)",
		email, passwordHash,
	)
	return err
}
