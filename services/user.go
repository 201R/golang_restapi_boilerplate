package services

import (
	"context"

	"github.com/201R/go_api_boilerplate/dtos"
	"github.com/201R/go_api_boilerplate/ent"
	"github.com/201R/go_api_boilerplate/repository"
)

type UserService interface {
	Get(ctx context.Context) ([]*ent.User, error)
	Create(ctx context.Context, userInput dtos.UserInput) (*ent.User, error)
	Update() (*ent.User, error)
	Delete(id int) (*ent.User, error)
	GetByID(id int) (*ent.User, error)
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
func (us *userService) Get(ctx context.Context) ([]*ent.User, error) {
	return us.Repo.Get(ctx)
}

// Create implements UserService
func (us *userService) Create(ctx context.Context, userInput dtos.UserInput) (*ent.User, error) {
	return us.Repo.Create(ctx, userInput)
}

// Delete implements UserService
func (*userService) Delete(id int) (*ent.User, error) {
	panic("unimplemented")
}

// Update implements UserService
func (*userService) Update() (*ent.User, error) {
	panic("unimplemented")
}

// GetByID implements UserService
func (*userService) GetByID(id int) (*ent.User, error) {
	panic("unimplemented")
}
