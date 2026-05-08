package service
import (
	"auth-service/internal/repository"
)

type UserService struct {
	Repo *repository.UserRepository
}

func (s *UserService) CreateUser(email string) error{
	return  s.Repo.CreateUser(email)
} 