package domain

import "github.com/ITK13201/portfolio/backend/ent"

type Login struct {
	Username string `json:"username,omitempty"`
	Password string `json:"password,omitempty"`
}

type LoginResponse struct {
	Token      string   `json:"token,omitempty"`
	Expiration string   `json:"expiration,omitempty"`
	User       ent.User `json:"user"`
}

type RefreshResponse struct {
	Token      string `json:"token,omitempty"`
	Expiration string `json:"expiration,omitempty"`
}
