package setup

import (
	"github.com/joho/godotenv"
	"github.com/pjchender/todo-mvc-backend/global"
	"github.com/pjchender/todo-mvc-backend/pkg/logger"
	"github.com/pjchender/todo-mvc-backend/pkg/setting"
	log "github.com/sirupsen/logrus"
	"os"
)

func Logger() error {
	log.SetFormatter(&log.TextFormatter{
		FullTimestamp: true,
	})
	logWriter := logger.NewLogWriter()
	log.SetOutput(logWriter)

	return nil
}

func Env(path ...string) error {
	m := os.Getenv("MODE")

	// 如果 mode 是 production 就不拿 local 的 .env（回直接使用 k8s 的 env）
	if m == "production" {
		return nil
	}

	// 如果 mode 不是 production 則讀取本機的 .env
	err := godotenv.Load(path...)
	if err != nil {
		log.Fatal("[runner] SetupEnv godotenv.Load() failed: ", err)
		return err
	}

	return nil
}

// Settings 會讀取 configs 與 env 中的設定，灌入到 global 的 settings 變數
func Settings() error {
	var err error
	s, err := setting.NewSetting()
	if err != nil {
		return err
	}

	global.AppSetting = s.ReadAppSetting()
	global.DatabaseSetting = s.ReadDBSetting()
	global.AuthSetting = s.ReadAuthSetting()
	global.ServerSetting = s.ReadServerSetting()
	global.GormSetting = s.ReadGormSetting()

	return nil
}
