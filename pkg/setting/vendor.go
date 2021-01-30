package setting

import (
	"github.com/pjchender/todo-mvc-backend/configs"
	"gorm.io/gorm"
)

func (s *Setting) ReadGormSetting() *gorm.Config {
	return configs.Gorm
}
