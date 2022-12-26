package endpoint

import (
	"github.com/balqisgautama/okami-auth/http/service"
	"net/http"
)

type healthEndpoint struct {
	AbstractEndpoint
}

var HealthEndpoint = healthEndpoint{}.New()

func (input healthEndpoint) New() (output healthEndpoint) {
	output.FileName = "HealthEndpoint.go"
	return
}

func (input healthEndpoint) CheckingHealth(responseWriter http.ResponseWriter, request *http.Request) {
	input.ServeEndpoint(service.HealthService.CheckingHealth, responseWriter, request, false)
}
