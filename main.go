package main

import (
	"fmt"
	"github.com/balqisgautama/okami-auth/config"
	"github.com/balqisgautama/okami-auth/seeder"
	"github.com/balqisgautama/okami-auth/server"
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

	fmt.Println(config.ApplicationConfiguration.GetServerHost())
	fmt.Println(config.ApplicationConfiguration.GetServerPort())
	fmt.Println(config.ApplicationConfiguration.GetServerVersion())
	fmt.Println(config.ApplicationConfiguration.GetServerResourceID())
	fmt.Println(config.ApplicationConfiguration.GetServerPrefixPath())
}
