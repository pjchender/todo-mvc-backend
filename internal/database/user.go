package database

import (
	"github.com/google/uuid"
	"github.com/pjchender/todo-mvc-backend/internal/model"
)

func (d *GormDatabase) GetUserByID(id uuid.UUID) (*model.User, error) {
	user := &model.User{}
	if err := d.DB.Preload("Todos").Take(user, id).Error; err != nil {
		return nil, err
	}

	return user, nil
}

func (d *GormDatabase) CreateUser(user *model.User) error {
	return d.DB.Create(user).Error
}

func (d *GormDatabase) FirstOrCreateUser(user *model.User) (*model.User, error) {
	err := d.DB.Where(model.User{FacebookUserID: user.FacebookUserID}).FirstOrCreate(user).Error
	if err != nil {
		return nil, err
	}

	return user, nil
}
