package config

import (
	"fmt"

	"github.com/spf13/viper"
)

type DatabaseConfiguration struct {
	Driver   string
	Dbname   string
	Username string
	Password string
	Host     string
	Port     string
	LogMode  bool
}

func DbConfiguration() string {
	psqlDBName := viper.GetString("PSQL_DB_NAME")
	psqlDBUser := viper.GetString("PSQL_DB_USER")
	psqlDBPassword := viper.GetString("PSQL_DB_PASSWORD")
	psqlDBHost := viper.GetString("PSQL_DB_HOST")
	psqlDBPort := viper.GetString("PSQL_DB_PORT")
	psqlDBSslMode := viper.GetString("PSQL_SSL_MODE")

	psqlDBConn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s",
		psqlDBHost, psqlDBUser, psqlDBPassword, psqlDBName,
	)

	if psqlDBPort != "" {
		psqlDBConn = fmt.Sprintf(
			"host=%s user=%s password=%s dbname=%s port=%s sslmode=%s",
			psqlDBHost, psqlDBUser, psqlDBPassword, psqlDBName, psqlDBPort, psqlDBSslMode,
		)
	}

	return psqlDBConn
}
