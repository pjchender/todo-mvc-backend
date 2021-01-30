package setting

import (
	"github.com/pjchender/todo-mvc-backend/configs"
	"github.com/pjchender/todo-mvc-backend/pkg/ginmode"
	"os"
)

// ReadAppSetting 會設定 APP
func (s *Setting) ReadAppSetting() *configs.App {
	modeENV := os.Getenv("MODE")
	if modeENV != "" {
		ginmode.Set(&s.defaultConfig.App, modeENV)
	}

	return &s.defaultConfig.App
}

func (s *Setting) ReadDBSetting() *configs.Database {
	DSN := os.Getenv("DSN")
	if DSN != "" {
		s.defaultConfig.Database.DSN = DSN
	}

	return &s.defaultConfig.Database
}

func (s *Setting) ReadAuthSetting() *configs.Auth {
	passwordSalt := os.Getenv("PASSWORD_SALT")
	if passwordSalt != "" {
		s.defaultConfig.Auth.Password.Salt = passwordSalt
	}

	JWTSecret := os.Getenv("JWT_SECRET")
	if JWTSecret != "" {
		s.defaultConfig.Auth.JWT.Secret = JWTSecret
	}

	return &s.defaultConfig.Auth
}

func (s *Setting) ReadServerSetting() *configs.Server {
	serverPort := os.Getenv("SERVER_PORT")
	if serverPort != "" {
		s.defaultConfig.Server.Port = serverPort
	}

	return &s.defaultConfig.Server
}
