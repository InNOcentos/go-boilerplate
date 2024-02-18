package user

import (
	"net/http"
	"time"

	userPb "github.com/InNOcentos/Doggo-track/backend/pkg/contracts/user"
	response "github.com/InNOcentos/Doggo-track/backend/pkg/response"
	validation "github.com/InNOcentos/Doggo-track/backend/pkg/validation"
	"github.com/gin-gonic/gin"
)

type GetPathParams struct {
	Id string `uri:"id"`
}

type GetResponseDto struct {
	Id         string
	Name       string
	Surname    string
	CreateDate time.Time
	UpdateDate time.Time
}

// @Sumary Get
// @Tags   users
// @Description Get user by id
// @ID get-user
// @Accept json
// @Produce json
// @Param id path string true "User id"
// @Success 200 {object} GetResponseDto
// @Failure 400 {object} response.HttpErrorResponse
// @Router /api/users/{id} [get]
func (h *userHttpHandler) Get(c *gin.Context) {
	var params GetPathParams

	err := c.ShouldBindUri(&params)
	if err != nil {
		errors := validation.HttpHandleBindErrors(err)
		c.JSON(http.StatusBadRequest, response.HttpErrorResponse{Errors: errors})
		return
	}

	req := &userPb.GetRequest{
		Id: params.Id,
	}

	grpcRes, err := h.grpcClient.Get(c.Request.Context(), req)
	if err != nil {
		c.JSON(http.StatusBadRequest, response.HttpErrorResponse{Errors: []string{err.Error()}})
		return
	}

	res := GetResponseDto{
		Id:         grpcRes.Id,
		Name:       grpcRes.Name,
		Surname:    grpcRes.Name,
		CreateDate: time.Unix(grpcRes.CreatedAtDate.Seconds, int64(grpcRes.CreatedAtDate.Nanos)),
		UpdateDate: time.Unix(grpcRes.UpdatedAtDate.Seconds, int64(grpcRes.UpdatedAtDate.Nanos)),
	}

	c.JSON(http.StatusOK, response.HttpOkResponse{Data: res})
}
