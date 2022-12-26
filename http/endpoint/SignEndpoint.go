package endpoint

import (
	"github.com/balqisgautama/okami-auth/constanta"
	"github.com/balqisgautama/okami-auth/http/service"
	"net/http"
)

type signEndpoint struct {
	AbstractEndpoint
}

var SignEndpoint = signEndpoint{}.New()

func (input signEndpoint) New() (output signEndpoint) {
	output.FileName = "SignEndpoint.go"
	return
}

func (input signEndpoint) UserSign(responseWriter http.ResponseWriter, request *http.Request) {
	input.FuncName = "UserSign"

	switch request.Method {
	case constanta.RequestPOST:
		input.ServeEndpoint(service.SignService.SignIn, responseWriter, request, false)
		break
	}
}
