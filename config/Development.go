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

func convertStringParamToInt(value string) int {
	intPort, _ := strconv.Atoi(value)
	return intPort
}
