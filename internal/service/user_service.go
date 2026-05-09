package service

import (
	"auth-service/internal/repository"

	"golang.org/x/crypto/bcrypt"
)

type UserService struct {
	Repo *repository.UserRepository
}

func (s *UserService) CreateUser(email string ,password string) error{
	passwordBytes := []byte(password)
	passwordHash,err := bcrypt.GenerateFromPassword(passwordBytes,bcrypt.DefaultCost)
	if err != nil{
		return err
	}
	return  s.Repo.CreateUser(email,string(passwordHash))
}