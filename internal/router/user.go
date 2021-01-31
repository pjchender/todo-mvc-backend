package router

import (
	"github.com/gin-gonic/gin"
	"github.com/pjchender/todo-mvc-backend/internal/database"
	v1 "github.com/pjchender/todo-mvc-backend/internal/router/api/v1"
)

func RegisterUser(db *database.GormDatabase, routerGroup *gin.RouterGroup, middleware ...gin.HandlerFunc) {
	userHandler := v1.NewUserHandler(db)
	userRouter := routerGroup.Group("/users", middleware...)
	{
		userRouter.GET("/me", userHandler.Me)
		//userRouter.GET("/:id", userHandler.Get)
		//userRouter.POST("/", userHandler.FirstOrCreate)
	}
}
