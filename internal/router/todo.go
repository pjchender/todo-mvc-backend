package router

import (
	"github.com/gin-gonic/gin"
	"github.com/pjchender/todo-mvc-backend/internal/database"
	v1 "github.com/pjchender/todo-mvc-backend/internal/router/api/v1"
)

func RegisterTodo(db *database.GormDatabase, routerGroup *gin.RouterGroup, middleware ...gin.HandlerFunc) {
	todoHandler := v1.NewTodoHandler(db)
	todoRouter := routerGroup.Group("/todos", middleware...)
	{
		todoRouter.GET("/", todoHandler.GetTodoByUserID)
		todoRouter.POST("/", todoHandler.CreateTodo)
		todoRouter.PATCH("/:id", todoHandler.UpdateTodo)
		todoRouter.DELETE("/:id", todoHandler.Delete)
	}
}
