package repository

import (
	"errors"

	"github.com/201R/go_api_boilerplate/ent"
)

type UserRepository interface {
	Get() ([]*ent.User, error)
	Create(user *ent.User) (*ent.User, error)
	Update(id string, user *ent.User) (*ent.User, error)
	Delete(id string) (*ent.User, error)
	GetByID(id string) (*ent.User, error)
}

// User repos from ent database
type userRepository struct {
	client *ent.UserClient
}

func NewUserRepository(client *ent.UserClient) UserRepository {
	return &userRepository{
		client: client,
	}
}

func (u *userRepository) Get() ([]*ent.User, error) {
	return nil, errors.New(("not implemented"))
}

func (u *userRepository) Create(user *ent.User) (*ent.User, error) {
	return nil, errors.New(("not implemented"))
}

func (u *userRepository) Update(id string, user *ent.User) (*ent.User, error) {
	return nil, errors.New(("not implemented"))
}

func (u *userRepository) Delete(id string) (*ent.User, error) {
	return nil, errors.New(("not implemented"))
}

func (u *userRepository) GetByID(id string) (*ent.User, error) {
	return nil, errors.New(("not implemented"))
}
