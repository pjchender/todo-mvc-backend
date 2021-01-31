package app

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/pjchender/todo-mvc-backend/internal/model"
)

func ParseUser(ctx *gin.Context) (*model.User, error) {
	user, exist := ctx.Get("user")
	if !exist {
		return nil, errors.New("parseUser failed: user is not exist")
	}
	u, ok := user.(*model.User)
	if !ok {
		return nil, errors.New("parseUser failed: can not do type assertion")
	}
	return u, nil
}
