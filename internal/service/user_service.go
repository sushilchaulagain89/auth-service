package service

import (
	"auth-service/internal/repository"
	"errors"

	"golang.org/x/crypto/bcrypt"
)

type UserService struct {
	Repo *repository.UserRepository
}

var (
	ErrEmailExists  = errors.New("email already exists")
	ErrInvalidLogin = errors.New("invalid credentials")
	ErrUserNotFound = errors.New("user not found")
)

func (s *UserService) CreateUser(email string, password string) error {
	passwordBytes := []byte(password)
	passwordHash, err := bcrypt.GenerateFromPassword(passwordBytes, bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	return s.Repo.CreateUser(email, string(passwordHash))
}

func (s *UserService) Login(email, password string) error {
	user, err := s.Repo.GetUserByEmail(email)
	if err != nil {
		return ErrInvalidLogin
	}

	err = bcrypt.CompareHashAndPassword(
		[]byte(user.PasswordHash),
		[]byte(password),
	)
	if err != nil {
		return ErrInvalidLogin
	}

	return nil
}
