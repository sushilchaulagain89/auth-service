package service

import (
	"auth-service/internal/auth"
	"auth-service/internal/repository"
	"errors"

	"golang.org/x/crypto/bcrypt"
)

type UserService struct {
	Repo *repository.UserRepository
}

// ---------------- ERRORS ----------------

var (
	ErrInvalidLogin = errors.New("invalid credentials")
	ErrUserExists   = errors.New("user already exists")
)

// ---------------- REGISTER ----------------

func (s *UserService) CreateUser(email string, password string) error {

	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	return s.Repo.CreateUser(email, string(hash))
}

// ---------------- LOGIN ----------------

func (s *UserService) Login(email, password string) (string, string, error) {

	user, err := s.Repo.GetUserByEmail(email)
	if err != nil {
		return "", "", ErrInvalidLogin
	}

	err = bcrypt.CompareHashAndPassword(
		[]byte(user.PasswordHash),
		[]byte(password),
	)

	if err != nil {
		return "", "", ErrInvalidLogin
	}

	// 🔐 JWT GENERATION
	access, refresh, err := auth.GenerateAllTokens(user.ID, user.Email)
	if err != nil {
		return "", "", err
	}

	return access, refresh, nil
}