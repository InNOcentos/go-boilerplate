package user

import (
	"net/http"

	userPb "github.com/InNOcentos/Doggo-track/backend/pkg/contracts/user"
	response "github.com/InNOcentos/Doggo-track/backend/pkg/response"
	validation "github.com/InNOcentos/Doggo-track/backend/pkg/validation"
	"github.com/gin-gonic/gin"
)

type CreateDto struct {
	Name    string `json:"name" binding:"required,min=1,max=30"`
	Surname string `json:"surname" binding:"required,min=1,max=30"`
}

type CreateResponseDto struct {
	Id string `json:"id"`
}

// @Sumary Create
// @Tags   users
// @Description Create new user
// @ID create-user
// @Accept json
// @Produce json
// @Param input body CreateDto true "user info"
// @Success 200 {object} CreateResponseDto
// @Failure 400 {object} response.HttpErrorResponse
// @Router /api/users [post]
func (h *userHttpHandler) Create(c *gin.Context) {
	var dto CreateDto

	err := c.ShouldBindJSON(&dto)
	if err != nil {
		errors := validation.HttpHandleBindErrors(err)
		c.JSON(http.StatusBadRequest, response.HttpErrorResponse{Errors: errors})
		return
	}

	req := &userPb.CreateRequest{
		Name:    dto.Name,
		Surname: dto.Surname,
	}

	grpcRes, err := h.grpcClient.Create(c.Request.Context(), req)
	if err != nil {
		c.JSON(http.StatusBadRequest, response.HttpErrorResponse{Errors: []string{err.Error()}})
		return
	}

	res := CreateResponseDto{
		Id: grpcRes.Id,
	}

	c.JSON(http.StatusOK, response.HttpOkResponse{Data: res})
}
