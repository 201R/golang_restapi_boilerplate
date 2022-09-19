package services

import (
	"github.com/201R/go_api_boilerplate/dtos"
	"github.com/201R/go_api_boilerplate/repository"
)

type AuthService interface {
	Login(input *dtos.AuthInput) (*dtos.LoginReponse, error)
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
func (authService) Login(input *dtos.AuthInput) (*dtos.LoginReponse, error) {
	panic("unimplemented")
}
