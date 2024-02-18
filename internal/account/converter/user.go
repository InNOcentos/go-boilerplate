package converter

import (
	"github.com/InNOcentos/Doggo-track/backend/internal/account/model"
	userPb "github.com/InNOcentos/Doggo-track/backend/pkg/contracts/user"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func UserPbCreateRequestToModel(req *userPb.CreateRequest) *model.CreateUser {
	return &model.CreateUser{
		Name:    req.Name,
		Surname: req.Surname,
	}
}

func ModelToUserPbGetResponse(u *model.User) *userPb.GetResponse {
	return &userPb.GetResponse{
		Id:            u.Id,
		Name:          u.Name,
		Surname:       u.Surname,
		CreatedAtDate: timestamppb.New(u.CreatedAt),
		UpdatedAtDate: timestamppb.New(u.UpdatedAt),
	}
}
