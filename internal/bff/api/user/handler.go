package user

import (
	userPb "github.com/InNOcentos/Doggo-track/backend/pkg/contracts/user"
)

type userHttpHandler struct {
	grpcClient userPb.UserServiceClient
}

func NewHttpHandler(c userPb.UserServiceClient) *userHttpHandler {
	return &userHttpHandler{c}
}
