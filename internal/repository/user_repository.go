package repository

import (
	"aurh-service/internal/db"
	"context"
)
 type UserRepository struct{}

 func (r *UserRepository) CreateUser(email string) error {
	query := "INSERT INTO users (email) VALUES ($1)"
	_,err := db.Pool.Exec(context.Background(),query,email)
	return  err
 }