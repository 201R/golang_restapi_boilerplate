package services

import (
	"context"

	"github.com/201R/go_api_boilerplate/dtos"
	"github.com/201R/go_api_boilerplate/ent"
	"github.com/201R/go_api_boilerplate/packages/helpers"
	"github.com/201R/go_api_boilerplate/repository"
)

type AuthService interface {
	Login(ctx context.Context, input *dtos.AuthInput) (*dtos.LoginReponse, error)
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
func (us authService) Login(ctx context.Context, input *dtos.AuthInput) (*dtos.LoginReponse, error) {
	user, err := us.Repo.GetByEmail(ctx, input.Email)
	if err != nil {
		return nil, err
	}

	if err = helpers.VerifyPassword(user.Password, input.Password); err != nil {
		return nil, err
	}

	token, err := helpers.GenerateToken(user.ID, "role")
	if err != nil {
		return nil, err
	}

	return &dtos.LoginReponse{
		Access_token: token,
		User:         *user,
		Token_type:   "Bearer",
	}, nil
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
