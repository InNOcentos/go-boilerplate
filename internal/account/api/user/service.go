package user

import (
	"github.com/InNOcentos/Doggo-track/backend/internal/account/service"
	userPb "github.com/InNOcentos/Doggo-track/backend/pkg/contracts/user"
)

type Implementation struct {
	userPb.UnimplementedUserServiceServer
	userService service.UserService
}

func NewImplementation(userService service.UserService) *Implementation {
	return &Implementation{userService: userService}
}
