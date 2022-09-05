package repository

import "github.com/201R/go_api_boilerplate/ent"

// User repos from ent database
type UserRepository struct {
	client *ent.UserClient
}

func NewUserRepository(client *ent.UserClient) UserRepository {
	return UserRepository{
		client: client,
	}
}
