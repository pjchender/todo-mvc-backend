package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/pjchender/todo-mvc-backend/global"
)

func ResponseHeader() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ctx.Header("Content-Type", "application/json")

		// add customize response header
		for header, value := range global.ServerSetting.ResponseHeaders {
			ctx.Header(header, value)
		}
	}
}
