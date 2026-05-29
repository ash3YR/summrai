package repository

import (
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
)

type UserRepository struct {
	DB *pgxpool.Pool
}

func NewUserRepository(db *pgxpool.Pool) *UserRepository {
	return &UserRepository{
		DB: db,
	}
}

func (r *UserRepository) CreateUser(
	name string,
	email string,
	passwordHash string,
) error {

	query := `
		INSERT INTO users
		(name, email, password_hash)
		VALUES ($1, $2, $3)
	`

	_, err := r.DB.Exec(
		context.Background(),
		query,
		name,
		email,
		passwordHash,
	)

	return err
}