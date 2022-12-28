package endpoint

import (
	"github.com/balqisgautama/okami-auth/http/service"
	"net/http"
)

type registrationEndpoint struct {
	AbstractEndpoint
}

var RegistrationEndpoint = registrationEndpoint{}.New()

func (input registrationEndpoint) New() (output registrationEndpoint) {
	output.FileName = "RegistrationEndpoint.go"
	return
}

func (input registrationEndpoint) Registration(responseWriter http.ResponseWriter, request *http.Request) {
	input.ServeEndpoint(service.RegistrationService.UserRegistration, responseWriter, request, false)
}

func (input registrationEndpoint) ActivationUser(responseWriter http.ResponseWriter, request *http.Request) {
	input.ServeEndpoint(service.RegistrationService.ActivationAccount, responseWriter, request, false)
}

func (input registrationEndpoint) ActivationUserResend(responseWriter http.ResponseWriter, request *http.Request) {
	input.ServeEndpoint(service.RegistrationService.ActivationAccountResend, responseWriter, request, false)
}
