package router

import (
	"github.com/naoto67/clean-architecture/src/interface/api/handler"
	"github.com/naoto67/clean-architecture/src/registry"

	"github.com/gin-gonic/gin"
)

func New() *gin.Engine {
	router := gin.Default()
	reg := registry.New()

	userHandler := handler.NewUserHandler(reg)
	userGroup := router.Group("/users")
	{
		userGroup.GET("/:id", userHandler.Show)
	}

	return router
}
