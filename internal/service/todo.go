package service

import (
	"github.com/google/uuid"
	"github.com/pjchender/todo-mvc-backend/internal/model"
)

type GetTodosByUserIDRequest struct {
	UserID uuid.UUID
}

type CreateTodoRequest struct {
	Title  string    `json:"title"`
	IsDone bool      `json:"isDone"`
	UserID uuid.UUID `json:"userId"`
}

type UpdateTodoRequest struct {
	TodoID uuid.UUID `json:"todoId"`
	Title  string    `json:"title"`
	IsDone bool      `json:"isDone"`
}

type DeleteTodoRequest struct {
	TodoID uuid.UUID
}

func (svc *Service) GetTodosByUserID(param GetTodosByUserIDRequest) ([]*model.Todo, error) {
	todos, err := svc.db.GetTodosByUserID(param.UserID)

	if err != nil {
		return nil, err
	}

	return todos, nil
}

func (svc *Service) CreateTodo(param CreateTodoRequest) (*model.Todo, error) {
	todo := model.Todo{
		Title:  param.Title,
		IsDone: param.IsDone,
		UserID: param.UserID,
	}

	err := svc.db.CreateTodo(&todo)

	if err != nil {
		return nil, err
	}

	return &todo, nil
}

func (svc *Service) UpdateTodo(param UpdateTodoRequest) (*model.Todo, error) {
	var err error

	// 先確認該筆 record 存在，找不到該 record 則回傳 404
	_, err = svc.db.GetTodoByID(param.TodoID)
	if err != nil {
		return nil, err
	}

	todo := model.Todo{
		ID:     param.TodoID,
		Title:  param.Title,
		IsDone: param.IsDone,
	}

	// 更新 record
	err = svc.db.UpdateTodo(&todo)
	if err != nil {
		return nil, err
	}

	// 取得更新後的 record
	updatedTodo, err := svc.db.GetTodoByID(param.TodoID)
	if err != nil {
		return nil, err
	}

	return updatedTodo, nil
}

func (svc *Service) DeleteTodo(param DeleteTodoRequest) error {
	var err error

	// 先確認該筆 record 存在，找不到該 record 則回傳 404
	_, err = svc.db.GetTodoByID(param.TodoID)
	if err != nil {
		return err
	}

	err = svc.db.DeleteTodoByID(param.TodoID)
	if err != nil {
		return err
	}

	return nil
}
