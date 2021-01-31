package router

import (
	"github.com/gin-gonic/gin"
	"github.com/pjchender/todo-mvc-backend/internal/database"
	"github.com/pjchender/todo-mvc-backend/internal/middleware"
	"github.com/pjchender/todo-mvc-backend/pkg/app"
	"net/http"
)

// New 用來建立所有的 router
func New(db *database.GormDatabase) *gin.Engine {
	router := gin.New()
	router.RedirectTrailingSlash = false

	router.Use(gin.Logger(), gin.Recovery(), middleware.ErrorHandler(), middleware.ResponseHeader())

	v1 := router.Group("/api/v1")
	{
		v1.POST("/", func(ctx *gin.Context) {
			var param struct {
				Foo int    `json:"foo" binding:"required"`
				Bar string `json:"bar" binding:"required"`
			}
			err := ctx.ShouldBind(&param)
			if success := app.SuccessOrAbort(ctx, http.StatusBadRequest, err); !success {
				return
			}

			ctx.JSON(200, gin.H{
				"status": "ok",
				"foo":    param.Foo,
				"bar":    param.Bar,
			})
		})
	}

	return router
}
