package v1

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/pjchender/todo-mvc-backend/internal/database"
	"github.com/pjchender/todo-mvc-backend/internal/model"
	"github.com/pjchender/todo-mvc-backend/internal/service"
	"github.com/pjchender/todo-mvc-backend/pkg/app"
	log "github.com/sirupsen/logrus"
	"net/http"
)

type TodoHandler struct {
	DB *database.GormDatabase
}

func NewTodoHandler(db *database.GormDatabase) *TodoHandler {
	return &TodoHandler{DB: db}
}

func (t *TodoHandler) GetTodoByUserID(ctx *gin.Context) {
	var err error
	user, err := app.ParseUser(ctx)
	if success := app.SuccessOrAbort(ctx, http.StatusInternalServerError, err); !success {
		log.Println("[todo] app.ParseUser failed: ", err)
		return
	}

	svc := service.New(ctx, t.DB)
	todos, err := svc.GetTodosByUserID(service.GetTodosByUserIDRequest{UserID: user.ID})
	if success := app.SuccessOrAbort(ctx, http.StatusInternalServerError, err); !success {
		log.Println("[todo] svc.GetTodosByUserID failed: ", err)
		return
	}

	ctx.JSON(http.StatusOK, model.ToExternalTodos(todos))
}

func (t *TodoHandler) CreateTodo(ctx *gin.Context) {
	var err error
	user, err := app.ParseUser(ctx)
	if success := app.SuccessOrAbort(ctx, http.StatusInternalServerError, err); !success {
		log.Println("[todo] app.ParseUser failed: ", err)
		return
	}

	param := service.CreateTodoRequest{
		UserID: user.ID,
	}
	err = ctx.ShouldBind(&param)
	if success := app.SuccessOrAbort(ctx, http.StatusBadRequest, err); !success {
		log.Println("[todo] ctx.ShouldBind failed: ", err)
		return
	}

	svc := service.New(ctx, t.DB)
	todo, err := svc.CreateTodo(param)
	if success := app.SuccessOrAbort(ctx, http.StatusInternalServerError, err); !success {
		log.Println("[todo] svc.CreateTodo failed: ", err)
		return
	}

	ctx.JSON(http.StatusOK, todo.ToExternal())
}

func (t *TodoHandler) UpdateTodo(ctx *gin.Context) {
	var err error
	todoIDStr := ctx.Param("id")
	todoID, err := uuid.Parse(todoIDStr)
	if success := app.SuccessOrAbort(ctx, http.StatusBadRequest, err); !success {
		log.Println("[todo] uuid.Parse failed: ", err)
		return
	}

	param := service.UpdateTodoRequest{TodoID: todoID}
	err = ctx.ShouldBind(&param)
	if success := app.SuccessOrAbort(ctx, http.StatusBadRequest, err); !success {
		log.Println("[todo] ctx.ShouldBind failed: ", err)
		return
	}

	fmt.Println("param", param)

	svc := service.New(ctx, t.DB)
	updatedCategory, err := svc.UpdateTodo(param)
	if success := app.SuccessOrAbort(ctx, http.StatusInternalServerError, err); !success {
		log.Println("[todo] svc.UpdateTodo failed: ", err)
		return
	}

	ctx.JSON(http.StatusOK, updatedCategory.ToExternal())
}

func (t *TodoHandler) Delete(ctx *gin.Context) {
	var err error
	todoIDStr := ctx.Param("id")
	todoID, err := uuid.Parse(todoIDStr)
	if success := app.SuccessOrAbort(ctx, http.StatusBadRequest, err); !success {
		log.Println("[todo] uuid.Parse failed: ", err)
		return
	}

	svc := service.New(ctx, t.DB)
	err = svc.DeleteTodo(service.DeleteTodoRequest{TodoID: todoID})
	if success := app.SuccessOrAbort(ctx, http.StatusInternalServerError, err); !success {
		log.Println("[todo] svc.DeleteTodo failed: ", err)
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"status": "OK",
	})
}
