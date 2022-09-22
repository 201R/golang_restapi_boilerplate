package services

import (
	"github.com/201R/go_api_boilerplate/dtos"
	"github.com/201R/go_api_boilerplate/ent"
	"github.com/201R/go_api_boilerplate/repository"
)

type AuthService interface {
	Login(input *dtos.AuthInput) (*dtos.LoginReponse, error)
	Logout() (bool, error)
	refresh(token string) (*dtos.LoginReponse, error)
	me(token string) (*ent.User, error)
}

type authService struct {
	Repo repository.UserRepository
}

func NewAuthService(userRepo repository.UserRepository) AuthService {
	return authService{
		Repo: userRepo,
	}
}

// Login implements AuthService
func (us authService) Login(input *dtos.AuthInput) (*dtos.LoginReponse, error) {
	panic("unimplemented")
}

// Logout implements AuthService
func (authService) Logout() (bool, error) {
	panic("unimplemented")
}

// refresh implements AuthService
func (authService) refresh(token string) (*dtos.LoginReponse, error) {
	panic("unimplemented")
}

// me implements AuthService
func (authService) me(token string) (*ent.User, error) {
	panic("unimplemented")
}
