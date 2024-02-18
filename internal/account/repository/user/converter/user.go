package converter

import (
	"github.com/InNOcentos/Doggo-track/backend/internal/account/model"
	"github.com/InNOcentos/Doggo-track/backend/internal/account/repository/user/entity"
)

func EntityToModel(u *entity.User) *model.User {
	return &model.User{
		Id:        u.HexId(),
		Name:      u.Name,
		Surname:   u.Surname,
		CreatedAt: u.CreatedAt,
		UpdatedAt: u.UpdatedAt,
	}
}
