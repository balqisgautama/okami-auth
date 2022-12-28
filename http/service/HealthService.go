package service

import (
	"github.com/balqisgautama/okami-auth/config/server"
	"github.com/balqisgautama/okami-auth/dto/res"
	"net/http"
)

type healthService struct {
	AbstractService
}

var HealthService = healthService{}.New()

func (input healthService) New() (output healthService) {
	output.FileName = "HealthService.go"
	return
}

func (input healthService) CheckingHealth(request *http.Request) (output res.APIResponse, header map[string]string, err error) {
	input.FuncName = "CheckingHealth"

	output.Status.Success = true
	output.Status.Message = "UP"
	if server.ServerConfig.DBConnection.Ping() != nil {
		output.Status.Message = "DOWN"
		output.Status.Detail = []string{"PostgreSQL"}
	}
	return
}
