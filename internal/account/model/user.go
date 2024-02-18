package model

import (
	"time"
)

type CreateUser struct {
	Name    string
	Surname string
}

type User struct {
	Id        string
	Name      string
	Surname   string
	CreatedAt time.Time
	UpdatedAt time.Time
}
