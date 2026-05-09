package repository

import (
	"auth-service/internal/db"
	"context"
)

type UserRepository struct{}
type User struct {
	Email        string
	PasswordHash string
}

func (r *UserRepository) CreateUser(email string, passwordHash string) error {
	query := "INSERT INTO users (email,password_hash) VALUES ($1,$2)"
	_, err := db.Pool.Exec(context.Background(), query, email, passwordHash)
	return err
}

func (r *UserRepository) GetUserByEmail(email string) (*User, error) {
	query := "SELECT email,password_hash FROM users where email=$1"
	row := db.Pool.QueryRow(context.Background(), query, email)
	user := &User{}
	err := row.Scan(&user.Email, &user.PasswordHash)
	if err != nil {
		return nil,err
	}
	return user, err
}
