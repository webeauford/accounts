package accounts

import "time"

type User struct {
	ID        string    `json:"id"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	Disabled  bool      `json:"disabled"`
	Roles     []Role    `json:"roles"`
	CreatedAT time.Time `json:"created_at"`
}

type Role struct {
	ID          string      `json:"id"`
	Name        string      `json:"name"`
	Permissions []Permisson `json:"permissions"`
}

type Permisson struct {
	Code string `json:"code"`
}
