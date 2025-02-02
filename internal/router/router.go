package router

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/pjchender/todo-mvc-backend/internal/database"
	"github.com/pjchender/todo-mvc-backend/internal/middleware"
)

// New 用來建立所有的 router
func New(db *database.GormDatabase) *gin.Engine {
	router := gin.New()
	router.RedirectTrailingSlash = false

	router.Use(gin.Logger(), gin.Recovery(), middleware.ErrorHandler(), middleware.ResponseHeader())

	config := cors.DefaultConfig()
	config.AllowAllOrigins = true
	config.AllowHeaders = []string{"Origin", "Content-Length", "Content-Type", "Facebook-Client-Token"}
	router.Use(cors.New(config))

	facebookMiddleware := middleware.AuthWithFacebook{DB: db}
	router.Use(facebookMiddleware.CheckFacebookLogin())

	v1 := router.Group("/api/v1")
	{
		RegisterUser(db, v1)
		RegisterTodo(db, v1)
	}

	return router
}
