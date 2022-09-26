package repository

import (
	"context"

	"github.com/201R/go_api_boilerplate/dtos"
	"github.com/201R/go_api_boilerplate/ent"
	"github.com/201R/go_api_boilerplate/ent/predicate"
	"github.com/201R/go_api_boilerplate/ent/user"
)

type UserRepository interface {
	Get(ctx context.Context) ([]*ent.User, error)
	Create(ctx context.Context, userInput dtos.UserInput) (*ent.User, error)
	Update(id string, user *ent.User) (*ent.User, error)
	Delete(id string) (*ent.User, error)
	GetByID(id string) (*ent.User, error)
	GetByEmail(ctx context.Context, email string) (*ent.User, error)
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

// Create implements UserRepository
func (u *userRepository) Create(ctx context.Context, userInput dtos.UserInput) (*ent.User, error) {
	return u.client.Create().
		SetFirstName(*userInput.FirstName).
		SetLastName(*userInput.LastName).
		SetEmail(userInput.Email).
		SetStatus(userInput.Status).
		SetPassword(userInput.Password).
		Save(ctx)
}

// Delete implements UserRepository
func (u *userRepository) Delete(id string) (*ent.User, error) {
	panic("unimplemented")
}

// Get implements UserRepository
func (u *userRepository) Get(ctx context.Context) ([]*ent.User, error) {
	return u.client.Query().All(ctx)
}

// GetByID implements UserRepository
func (u *userRepository) GetByID(id string) (*ent.User, error) {
	panic("unimplemented")
}

// Update implements UserRepository
func (u *userRepository) Update(id string, user *ent.User) (*ent.User, error) {
	panic("unimplemented")
}

// GetByEmail implements UserRepository
func (u *userRepository) GetByEmail(ctx context.Context, email string) (*ent.User, error) {
	return u.client.Query().Where(predicate.User(user.EmailEQ(email))).First(ctx)
}
