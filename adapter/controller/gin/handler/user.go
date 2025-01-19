package handler

import (
	"go-api-arch-clean/adapter/controller/gin/presenter"
	"go-api-arch-clean/api"
	"net/http"

	"github.com/gin-gonic/gin"
)

type UserHandler struct {
}

func NewUserHandler() *UserHandler {
	return &UserHandler{}
}

func (a *UserHandler) GetUserById(c *gin.Context, ID int) {
	c.JSON(http.StatusOK, &presenter.UserResponse{
		ApiVersion: api.Version,
		Data: presenter.User{
			Kind: "user",
			Id:   1,
			Name: "jun",
		},
	})
}
