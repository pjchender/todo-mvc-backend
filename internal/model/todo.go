package model

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
	"time"
)

type Todo struct {
	ID        uuid.UUID `gorm:"primaryKey;uniqueIndex;type:uuid;default:uuid_generate_v4()"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt
	Title     string
	IsDone    bool

	// belongs to User
	UserID uuid.UUID
}

type TodoExternal struct {
	ID     uuid.UUID `json:"id"`
	Title  string    `json:"title"`
	IsDone bool      `json:"isDone"`
}

func (t *Todo) ToExternal() TodoExternal {
	externalTodo := TodoExternal{
		ID:     t.ID,
		Title:  t.Title,
		IsDone: t.IsDone,
	}

	return externalTodo
}

func ToExternalTodos(todos []*Todo) []*TodoExternal {
	externalTodos := make([]*TodoExternal, len(todos))
	for i, todo := range todos {
		todoExternal := todo.ToExternal()
		externalTodos[i] = &todoExternal
	}

	return externalTodos
}
