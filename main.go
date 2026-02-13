package main

import (
	"time"

	"be-mini-project/config"
	"be-mini-project/infra/database"
	"be-mini-project/infra/logger"
	"be-mini-project/migrations"
	"be-mini-project/routers"

	_ "be-mini-project/docs"

	"github.com/spf13/viper"
)

// @title						MINI PROJECT API
// @version						1.0
// @securityDefinitions.apikey	BearerAuth
// @in							header
// @name						Authorization
func main() {
	// set timezone
	viper.SetDefault("SERVER_TIMEZONE", "Asia/Jakarta")
	loc, _ := time.LoadLocation(viper.GetString("SERVER_TIMEZONE"))
	time.Local = loc

	if err := config.SetupConfig(); err != nil {
		logger.Fatalf("config SetupConfig() error: %s", err)
	}
	psqlDBConf := config.DbConfiguration()

	if err := database.DbConnection(psqlDBConf); err != nil {
		logger.Fatalf("database DbConnection error: %s", err)
	}

	migrations.Migrate()

	router := routers.SetupRoute()
	logger.Fatalf("%v", router.Run(config.ServerConfig()))
}
