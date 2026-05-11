package repository

import (
	"auth-service/internal/db"
	"context"
)

type UserRepository struct{}

type User struct {
	ID           int
	Email        string
	PasswordHash string
}

// ---------------- CREATE USER ----------------

func (r *UserRepository) CreateUser(email string, passwordHash string) error {

	query := `
		INSERT INTO users (email, password_hash)
		VALUES ($1, $2)
	`

	_, err := db.Pool.Exec(context.Background(), query, email, passwordHash)
	return err
}

// ---------------- GET USER ----------------

func (r *UserRepository) GetUserByEmail(email string) (*User, error) {

	query := `
		SELECT id, email, password_hash
		FROM users
		WHERE email = $1
	`

	row := db.Pool.QueryRow(context.Background(), query, email)

	user := &User{}

	err := row.Scan(&user.ID, &user.Email, &user.PasswordHash)
	if err != nil {
		return nil, err
	}

	return user, nil
}