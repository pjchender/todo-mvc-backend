package global

import (
	"github.com/pjchender/todo-mvc-backend/configs"
	"gorm.io/gorm"
)

var (
	AppSetting      *configs.App
	ServerSetting   *configs.Server
	DatabaseSetting *configs.Database
	AuthSetting     *configs.Auth
	GormSetting     *gorm.Config
)
