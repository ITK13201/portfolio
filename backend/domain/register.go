package domain

type Register struct {
	Username string `json:"username,omitempty"`
	Password string `json:"password,omitempty"`
}
