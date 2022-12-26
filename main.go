package main

import (
	"fmt"
	"github.com/balqisgautama/okami-auth/config"
	"github.com/balqisgautama/okami-auth/config/server"
	"github.com/balqisgautama/okami-auth/http/router"
	"github.com/balqisgautama/okami-auth/seeder"
	"github.com/balqisgautama/okami-auth/util"
	"os"
)

func main() {
	var arguments = "development"
	args := os.Args
	if len(args) > 1 {
		arguments = args[1]
	}

	config.GenerateConfiguration(arguments)
	server.SetServerConfig()
	seeder.DBMigrate()
	util.InitializeLogger()

	err := server.ServerConfig.DBConnection.Ping()
	if err != nil {
		fmt.Println("Connecting failed (PostgreSQL)", err)
	}

	defer func() {
		err := server.ServerConfig.DBConnection.Close()
		if err != nil {
			fmt.Println("Connecting failed (PostgreSQL)", err)
		}
	}()

	router.ApiController(config.ApplicationConfiguration.GetServerPort())
}
