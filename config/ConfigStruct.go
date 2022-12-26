package config

import (
	"fmt"
	"github.com/kelseyhightower/envconfig"
	"github.com/tkanos/gonfig"
	"os"
)

var ApplicationConfiguration Configuration

type Configuration interface {
	GetServerHost() string
	GetServerPort() int
	GetServerVersion() string
	GetServerResourceID() string
	GetServerPrefixPath() string
	GetPostgreSQLAddress() string
	GetPostgreSQLSchema() string
	GetPostgreSQLMaxOpenConnection() int
	GetPostgreSQLMaxIdleConnection() int
	GetPostgreSQLAddressView() string
	GetPostgreSQLSchemaView() string
	GetPostgreSQLMaxOpenConnectionView() int
	GetPostgreSQLMaxIdleConnectionView() int
	GetJWTKey() string
}

func GenerateConfiguration(arg string) {
	enviName := "GeneralConfiguration"
	var err error

	temp := DevelopmentConfig{}
	err = gonfig.GetConf(os.Getenv(enviName)+"config_development.json", &temp)
	if err != nil {
		fmt.Print(err)
		os.Exit(2) // exit 2
	}
	err = envconfig.Process(os.Getenv(enviName)+"config_development.json", &temp)
	if err != nil {
		fmt.Print(err)
		os.Exit(2)
	}
	ApplicationConfiguration = &temp
}
