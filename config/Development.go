package config

import (
	"strconv"
)

type DevelopmentConfig struct {
	Configuration
	Server struct {
		Host       string `envconfig:"HOST"`
		Port       string `envconfig:"PORT"`
		Version    string `envconfig:"VERSION"`
		ResourceID string `envconfig:"RESOURCE_ID"`
		PrefixPath string `envconfig:"PREFIX_PATH"`
	} `json:"server"`
	Postgresql struct {
		Address           string `envconfig:"DB_CONNECTION"`
		Schema            string `envconfig:"DB_SCHEMA"`
		MaxOpenConnection int    `json:"max_open_connection"`
		MaxIdleConnection int    `json:"max_idle_connection"`
	} `json:"postgresql"`
	PostgresqlView struct {
		Address           string `envconfig:"DB_VIEW_CONNECTION"`
		Schema            string `envconfig:"DB_VIEW_SCHEMA"`
		MaxOpenConnection int    `json:"max_open_connection"`
		MaxIdleConnection int    `json:"max_idle_connection"`
	} `json:"postgresql_view"`
	JWT struct {
		Key string `envconfig:"JWT_KEY"`
	}
}

func (input DevelopmentConfig) GetServerHost() string {
	return input.Server.Host
}
func (input DevelopmentConfig) GetServerPort() int {
	return convertStringParamToInt(input.Server.Port)
}
func (input DevelopmentConfig) GetServerVersion() string {
	return input.Server.Version
}
func (input DevelopmentConfig) GetServerResourceID() string {
	return input.Server.ResourceID
}
func (input DevelopmentConfig) GetServerPrefixPath() string {
	return input.Server.PrefixPath
}
func (input DevelopmentConfig) GetPostgreSQLAddress() string {
	return input.Postgresql.Address
}
func (input DevelopmentConfig) GetPostgreSQLSchema() string {
	return input.Postgresql.Schema
}
func (input DevelopmentConfig) GetPostgreSQLMaxOpenConnection() int {
	return input.Postgresql.MaxOpenConnection
}
func (input DevelopmentConfig) GetPostgreSQLMaxIdleConnection() int {
	return input.Postgresql.MaxIdleConnection
}
func (input DevelopmentConfig) GetPostgreSQLAddressView() string {
	return input.PostgresqlView.Address
}
func (input DevelopmentConfig) GetPostgreSQLSchemaView() string {
	return input.PostgresqlView.Schema
}
func (input DevelopmentConfig) GetPostgreSQLMaxOpenConnectionView() int {
	return input.PostgresqlView.MaxOpenConnection
}
func (input DevelopmentConfig) GetPostgreSQLMaxIdleConnectionView() int {
	return input.PostgresqlView.MaxIdleConnection
}

func (input DevelopmentConfig) GetJWTKey() string {
	return input.JWT.Key
}

func convertStringParamToInt(value string) int {
	intPort, _ := strconv.Atoi(value)
	return intPort
}
