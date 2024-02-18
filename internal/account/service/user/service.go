package user

import "github.com/InNOcentos/Doggo-track/backend/internal/account/repository"

type serv struct {
	repo repository.UserRepository
}

func NewService(repo repository.UserRepository) *serv {
	return &serv{repo: repo}
}
