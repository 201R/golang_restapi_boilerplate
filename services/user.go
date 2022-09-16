package services

import (
	"github.com/201R/go_api_boilerplate/ent"
	"github.com/201R/go_api_boilerplate/repository"
)

type UserService interface {
	Get() ([]*ent.User, error)
}

type userService struct {
	Repo repository.UserRepository
}

func NewUserService(userRepo repository.UserRepository) UserService {
	return &userService{
		Repo: userRepo,
	}
}

// Get implements UserService
func (us *userService) Get() ([]*ent.User, error) {
	panic("unimplemented")
}
