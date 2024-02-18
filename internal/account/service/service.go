package service

import (
	"context"

	"github.com/InNOcentos/Doggo-track/backend/internal/account/model"
)

type UserService interface {
	Get(ctx context.Context, id string) (*model.User, error)
	Create(ctx context.Context, dto *model.CreateUser) (string, error)
}
