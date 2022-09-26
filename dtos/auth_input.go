package dtos

import "github.com/201R/go_api_boilerplate/ent"

type AuthInput struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type LoginReponse struct {
	Access_token string   `json:"access_token"`
	Token_type   string   `json:"token_type"`
	User         ent.User `json:"user"`
}
