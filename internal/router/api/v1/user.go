package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/pjchender/todo-mvc-backend/internal/database"
	"github.com/pjchender/todo-mvc-backend/internal/service"
	"github.com/pjchender/todo-mvc-backend/pkg/app"
	log "github.com/sirupsen/logrus"
	"net/http"
)

type UserHandler struct {
	DB *database.GormDatabase
}

func NewUserHandler(db *database.GormDatabase) *UserHandler {
	return &UserHandler{DB: db}
}

func (d *UserHandler) Me(ctx *gin.Context) {
	var err error

	user, err := app.ParseUser(ctx)
	if success := app.SuccessOrAbort(ctx, http.StatusInternalServerError, err); !success {
		log.Println("[user] app.ParseUser failed: ", err)
		return
	}

	ctx.JSON(http.StatusOK, user.ToExternal())
}

func (d *UserHandler) Get(ctx *gin.Context) {
	var err error
	userIDStr := ctx.Param("id")
	userID, err := uuid.Parse(userIDStr)
	if success := app.SuccessOrAbort(ctx, http.StatusBadRequest, err); !success {
		log.Println("[user] uuid.Parse failed: ", err)
		return
	}

	svc := service.New(ctx, d.DB)
	user, err := svc.GetUserByID(service.GetUserRequest{ID: userID})
	if success := app.SuccessOrAbort(ctx, http.StatusInternalServerError, err); !success {
		log.Println("[user] svc.GetUserByID failed: ", err)
		return
	}

	ctx.JSON(http.StatusOK, user.ToExternal())
}

func (d *UserHandler) FirstOrCreate(ctx *gin.Context) {
	param := service.FirstOrCreateUserRequest{}

	svc := service.New(ctx, d.DB)
	user, err := svc.FirstOrCreateUser(param)
	if success := app.SuccessOrAbort(ctx, http.StatusInternalServerError, err); !success {
		log.Println("[user] svc.FirstOrCreateUser failed: ", err)
	}

	ctx.JSON(http.StatusOK, user.ToExternal())
}
