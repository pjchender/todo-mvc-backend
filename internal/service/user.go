package service

import (
	"github.com/google/uuid"
	"github.com/pjchender/todo-mvc-backend/internal/model"
)

type GetUserRequest struct {
	ID uuid.UUID
}

// FirstOrCreateUserRequest 是從 JSON 取得
type FirstOrCreateUserRequest struct {
	FacebookUserID uint   `json:"facebookUserId" binding:"required"`
	Email          string `json:"email"`
}

func (svc *Service) GetUserByID(param GetUserRequest) (*model.User, error) {
	user, err := svc.db.GetUserByID(param.ID)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (svc *Service) FirstOrCreateUser(param FirstOrCreateUserRequest) (*model.User, error) {
	user := model.User{
		FacebookUserID: param.FacebookUserID,
		Email:          param.Email,
	}

	theUser, err := svc.db.FirstOrCreateUser(&user)
	if err != nil {
		return nil, err
	}

	return theUser, nil
}
