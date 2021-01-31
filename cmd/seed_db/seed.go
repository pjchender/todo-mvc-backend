package main

import (
	"github.com/pjchender/todo-mvc-backend/global"
	"github.com/pjchender/todo-mvc-backend/internal/database"
	"github.com/pjchender/todo-mvc-backend/pkg/setup"
	log "github.com/sirupsen/logrus"
)

func init() {
	var err error

	// setupEnv should invoke before setupSetting()
	err = setup.Env()
	if err != nil {
		log.Fatalf("init.setupEnv failed: %v", err)
	}

	err = setup.Settings()
	if err != nil {
		log.Fatalf("init.setupSetting failed: %v", err)
	}
}

func main() {
	db, err := database.New(global.DatabaseSetting.DSN, global.GormSetting)
	if err != nil {
		panic(err)
	}

	db.AutoMigrate()
	//db.Seed(global.AuthSetting)
}
