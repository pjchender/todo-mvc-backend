package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/pjchender/todo-mvc-backend/global"
	"github.com/pjchender/todo-mvc-backend/internal/database"
	"github.com/pjchender/todo-mvc-backend/internal/router"
	"github.com/pjchender/todo-mvc-backend/pkg/setup"
	log "github.com/sirupsen/logrus"
	"net/http"
	"os"
	"os/signal"
	"strconv"
	"syscall"
)

func init() {
	var err error
	err = setup.Logger()
	if err != nil {
		log.Fatalf("init.setupLogger failed: %v", err)
	}

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
	// init database
	db, err := database.New(global.DatabaseSetting.DSN, global.GormSetting)
	if err != nil {
		log.Fatalf("[main] database.New failed: %v", err)
	}
	db.AutoMigrate()

	// start gin server
	engine := router.New(db)
	Run(engine)

	// shut down server
	quit := make(chan os.Signal)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Println("Shutting down server...")
}

func Run(engine *gin.Engine) {
	var httpHandler http.Handler = engine

	// start tls https connection
	go func() {
		isSSLEnabled := *global.ServerSetting.SSL.Enabled
		if !isSSLEnabled {
			return
		}

		addrTLS := fmt.Sprintf("%s:%d", global.ServerSetting.SSL.ListenAddr, global.ServerSetting.SSL.Port)
		fmt.Println("Started Listening for TLS connection on " + addrTLS)
		err := http.ListenAndServeTLS(addrTLS, global.ServerSetting.SSL.CertFile, global.ServerSetting.SSL.CertKey,
			httpHandler)
		if err != nil {
			log.Fatalf("[main] run - http.ListenAndServeTLS failed: %v", err)
		}
	}()

	// start plain http connection
	go func() {
		serverPort, err := strconv.Atoi(global.ServerSetting.Port)
		if err != nil {
			log.Fatal("[main] run - strconv.Atoi failed: ", err)
		}

		addr := fmt.Sprintf("%s:%d", global.ServerSetting.ListenAddr, serverPort)
		err = http.ListenAndServe(addr, httpHandler)
		if err != nil {
			log.Fatalf("[main] run - http.ListenAndServe failed: %v", err)
		}

		fmt.Println("Started Listening for plain HTTP connection on " + addr)
	}()
}
