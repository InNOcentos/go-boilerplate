package user

import (
	"context"

	"github.com/InNOcentos/Doggo-track/backend/internal/account/model"
)

func (s *serv) Create(ctx context.Context, dto *model.CreateUser) (string, error) {
	return s.repo.Create(ctx, dto)
}
