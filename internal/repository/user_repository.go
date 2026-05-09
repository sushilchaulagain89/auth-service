package repository

import (
	"auth-service/internal/db"
	"context"
)
 type UserRepository struct{}

 func (r *UserRepository) CreateUser(email string,passwordHash string) error {
	query := "INSERT INTO users (email,password_hash) VALUES ($1,$2)"
	_,err := db.Pool.Exec(context.Background(),query,email,passwordHash)
	return  err
 }