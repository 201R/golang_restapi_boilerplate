package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/201R/go_api_boilerplate/packages/logger"
	"github.com/201R/go_api_boilerplate/packages/setting"
	"github.com/201R/go_api_boilerplate/router"
	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
)

func main() {
	setting.Setup()
	closeLogFile, err := logger.Setup()
	if err != nil {
		logger.Fatal(err)
	}
	defer closeLogFile()

	gin.SetMode(setting.ServerSetting.RunMode)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
		logger.Warnf("Defaulting to port %s", port)
	}

	routersInit := router.InitRouter()
	readTimeout := setting.ServerSetting.ReadTimeout
	writeTimeout := setting.ServerSetting.WriteTimeout
	endPoint := fmt.Sprintf(":%s", port)

	server := &http.Server{
		Addr:         endPoint,
		Handler:      routersInit,
		ReadTimeout:  readTimeout,
		WriteTimeout: writeTimeout,
	}

	logger.Infof("[info] start http server listening %s", endPoint)

	server.ListenAndServe()
}
