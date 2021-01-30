package ginmode

import (
	"github.com/gin-gonic/gin"
	"github.com/pjchender/todo-mvc-backend/configs"
	log "github.com/sirupsen/logrus"
)

const (
	Dev  = "development"
	Prod = "production"
	Test = "test"
)

func IsDev(mode string) bool {
	return mode == Dev || mode == Test
}

func Set(appSetting *configs.App, mode string) {
	switch mode {
	case Prod:
		gin.SetMode(gin.ReleaseMode)
		appSetting.Mode = Prod
	case Dev:
		gin.SetMode(gin.DebugMode)
		appSetting.Mode = Dev
	case Test:
		gin.SetMode(gin.TestMode)
		appSetting.Mode = Test
	default:
		appSetting.Mode = Dev
		log.Errorf("[internal/pkg/ginmode] unknown mode '%v', reset to development", mode)
	}
}
