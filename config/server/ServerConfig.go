package server

import (
	"database/sql"
	"github.com/balqisgautama/okami-auth/config"
	"github.com/balqisgautama/okami-auth/config/database"
)

var ServerConfig serverConfig

type serverConfig struct {
	DBConnection     *sql.DB
	DBConnectionView *sql.DB
}

func SetServerConfig() {
	dbParam := config.ApplicationConfiguration.GetPostgreSQLSchema()
	dbConnection := config.ApplicationConfiguration.GetPostgreSQLAddress()
	dbMaxOpenConnection := config.ApplicationConfiguration.GetPostgreSQLMaxOpenConnection()
	dbMaxIdleConnection := config.ApplicationConfiguration.GetPostgreSQLMaxIdleConnection()
	ServerConfig.DBConnection = database.GetDbConnection(dbParam, dbConnection, dbMaxOpenConnection, dbMaxIdleConnection)

	dbParamView := config.ApplicationConfiguration.GetPostgreSQLSchemaView()
	dbConnectionView := config.ApplicationConfiguration.GetPostgreSQLAddressView()
	dbMaxOpenConnectionView := config.ApplicationConfiguration.GetPostgreSQLMaxOpenConnectionView()
	dbMaxIdleConnectionView := config.ApplicationConfiguration.GetPostgreSQLMaxIdleConnectionView()
	ServerConfig.DBConnectionView = database.GetDbConnection(dbParamView, dbConnectionView, dbMaxOpenConnectionView, dbMaxIdleConnectionView)
}
