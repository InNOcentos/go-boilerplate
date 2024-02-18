package user

import (
	"context"

	"github.com/InNOcentos/Doggo-track/backend/internal/account/model"
)

func (s *serv) Get(ctx context.Context, id string) (*model.User, error) {
	return s.repo.Get(ctx, id)
}
