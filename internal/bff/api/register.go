package api

import (
	user "github.com/InNOcentos/Doggo-track/backend/internal/bff/api/user"
	userPb "github.com/InNOcentos/Doggo-track/backend/pkg/contracts/user"
	"github.com/gin-gonic/gin"
)

func RegisterUserHttpHandlers(r *gin.RouterGroup, c userPb.UserServiceClient) {
	h := user.NewHttpHandler(c)

	g := r.Group("/users")

	g.GET("/:id", h.Get)
	g.POST("/", h.Create)
}
