package user

import (
	"context"
	"fmt"

	"github.com/InNOcentos/Doggo-track/backend/internal/account/converter"
	userPb "github.com/InNOcentos/Doggo-track/backend/pkg/contracts/user"
)

func (i *Implementation) Get(ctx context.Context, req *userPb.GetRequest) (*userPb.GetResponse, error) {
	user, err := i.userService.Get(ctx, req.GetId())
	if err != nil {
		return nil, err
	}

	fmt.Println(user)

	return converter.ModelToUserPbGetResponse(user), nil
}
