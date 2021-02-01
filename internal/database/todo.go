package database

import (
	"github.com/google/uuid"
	"github.com/pjchender/todo-mvc-backend/internal/model"
)

func (d *GormDatabase) GetTodosByUserID(userID uuid.UUID) ([]*model.Todo, error) {
	var todos []*model.Todo

	if err := d.DB.Where("user_id = ?", userID).Find(&todos).Error; err != nil {
		return nil, err
	}

	return todos, nil
}

func (d *GormDatabase) CreateTodo(todo *model.Todo) error {
	return d.DB.Create(todo).Error
}

func (d *GormDatabase) GetTodoByID(todoID uuid.UUID) (*model.Todo, error) {
	todo := model.Todo{}
	err := d.DB.Take(&todo, todoID).Error
	if err != nil {
		return nil, err
	}

	return &todo, nil
}

func (d *GormDatabase) UpdateTodo(todo *model.Todo) error {
	// 定義能被更新的 db 欄位
	values := map[string]interface{}{
		"title":  todo.Title,
		"is_done": todo.IsDone,
	}

	return d.DB.Model(&model.Todo{ID: todo.ID}).Updates(values).Error
}

func (d *GormDatabase) DeleteTodoByID(todoID uuid.UUID) error {
	return d.DB.Delete(&model.Todo{ID: todoID}).Error
}
