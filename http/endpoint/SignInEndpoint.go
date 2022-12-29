package endpoint

import (
	"github.com/balqisgautama/okami-auth/http/service/sign"
	"net/http"
)

type signInEndpoint struct {
	AbstractEndpoint
}

var SignInEndpoint = signInEndpoint{}.New()

func (input signInEndpoint) New() (output signInEndpoint) {
	output.FileName = "SignInEndpoint.go"
	return
}

func (input signInEndpoint) Step1(responseWriter http.ResponseWriter, request *http.Request) {
	input.ServeEndpoint(sign.SignInService.Step1, responseWriter, request, false)
}

func (input signInEndpoint) Step2(responseWriter http.ResponseWriter, request *http.Request) {
	input.ServeEndpoint(sign.SignInService.Step2, responseWriter, request, false)
}

func (input signInEndpoint) Step3(responseWriter http.ResponseWriter, request *http.Request) {
	input.ServeEndpoint(sign.SignInService.Step3, responseWriter, request, false)
}
