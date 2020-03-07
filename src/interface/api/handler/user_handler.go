package handler

import (
	"net/http"
	"strconv"

	"github.com/naoto67/clean-architecture/src/registry"
	"github.com/naoto67/clean-architecture/src/usecase"

	"github.com/gin-gonic/gin"
)

type UserHandler interface {
	Show(c *gin.Context)
}

type userHandler struct {
	userInteractor usecase.UserInteractor
}

func NewUserHandler(reg registry.Registry) UserHandler {
	return &userHandler{
		userInteractor: usecase.NewUserInteractor(reg.NewUserRepository()),
	}
}

func (handler *userHandler) Show(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{})
		return
	}
	user, err := handler.userInteractor.FindById(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{})
		return
	}
	c.JSON(http.StatusOK, user)
}
