package user

import (
	"context"

	"github.com/InNOcentos/Doggo-track/backend/internal/account/converter"
	userPb "github.com/InNOcentos/Doggo-track/backend/pkg/contracts/user"
)

func (i *Implementation) Create(ctx context.Context, req *userPb.CreateRequest) (*userPb.CreateResponse, error) {
	id, err := i.userService.Create(ctx, converter.UserPbCreateRequestToModel(req))
	if err != nil {
		return nil, err
	}

	return &userPb.CreateResponse{Id: id}, nil
}
