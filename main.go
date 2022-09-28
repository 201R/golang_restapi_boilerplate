package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	"github.com/201R/go_api_boilerplate/database"
	"github.com/201R/go_api_boilerplate/packages/logger"
	"github.com/201R/go_api_boilerplate/packages/setting"
	"github.com/201R/go_api_boilerplate/router"
	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
	"github.com/nicksnyder/go-i18n/v2/i18n"
	"golang.org/x/text/language"
)

// @title Golang Rest API
// @version 1.0
// @description This is a real world sample golang Rest API server
// @termsOfService http://swagger.io/terms/
// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io
// @license.name MIT
// @license.url https://github.com/yhagio/github.com/yhagio/go_api_boilerplate/blob/master/LICENSE
// @host localhost:8090
// @BasePath /
// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization
func main() {
	setting.Setup()
	client := database.Connect()
	defer client.Close()

	bundle := i18n.NewBundle(language.English)
	bundle.RegisterUnmarshalFunc("json", json.Unmarshal)
	// bundle.MustLoadMessageFile("i18n/translation/en/*.en.json")
	// bundle.MustLoadMessageFile("i18n/translation/fr/*.fr.json")

	closeLogFile, err := logger.Setup()
	if err != nil {
		logger.Fatal(err)
	}
	defer closeLogFile()

	gin.SetMode(setting.ServerSetting.RunMode)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8090"
		logger.Warnf("Defaulting to port %s", port)
	}

	routersInit := router.InitRouter(*client)
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
