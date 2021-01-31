package model

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
	"time"
)

type User struct {
	ID             uuid.UUID `gorm:"primaryKey;uniqueIndex;type:uuid;default:uuid_generate_v4()"`
	CreatedAt      time.Time
	UpdatedAt      time.Time
	DeletedAt      gorm.DeletedAt
	FacebookUserID uint `gorm:"uniqueIndex"`
	Email          string

	// has many todos
	Todos []*Todo
}

type UserExternal struct {
	ID             uuid.UUID       `json:"id"`
	FacebookUserID uint            `json:"facebookId"`
	Email          string          `json:"email"`
	Todos          []*TodoExternal `json:"todos"`
}

func (u *User) ToExternal() UserExternal {
	externalUser := UserExternal{
		ID:             u.ID,
		FacebookUserID: u.FacebookUserID,
		Email:          u.Email,
		Todos:          ToExternalTodos(u.Todos),
	}

	return externalUser
}
