package main

import (
	log "github.com/sirupsen/logrus"
	"github.com/softarch-project/menu-api/config"
	"github.com/softarch-project/menu-api/httpserver"
	"github.com/softarch-project/menu-api/pkg/database"
	"github.com/softarch-project/menu-api/pkg/logger"
)

var serverConfig *config.Config

func init() {
	serverConfig = config.LoadConfig()
	logger.InitLogger(serverConfig.App)
}

func main() {
	db, err := database.NewMySQLDatabaseConnection(serverConfig)
	if err != nil {
		log.Fatalf("error, create mysql database connection, %s", err.Error())
	}

	server := httpserver.NewHTTPServer(serverConfig, db)

	server.Start()
}
