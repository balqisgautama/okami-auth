package endpoint

import (
	"github.com/balqisgautama/okami-auth/constanta"
	"github.com/balqisgautama/okami-auth/http/service"
	"net/http"
)

type resourceEndpoint struct {
	AbstractEndpoint
}

var ResourceEndpoint = resourceEndpoint{}.New()

func (input resourceEndpoint) New() (output resourceEndpoint) {
	output.FileName = "ResourceEndpoint.go"
	return
}

func (input resourceEndpoint) Resource(responseWriter http.ResponseWriter, request *http.Request) {
	input.FuncName = "Resource"

	switch request.Method {
	case constanta.RequestPOST:
		input.ServeEndpoint(service.ResourceService.CreateResource, responseWriter, request, true)
		break
	case constanta.RequestPUT:
		input.ServeEndpoint(service.ResourceService.EditResource, responseWriter, request, true)
		break
	case constanta.RequestDELETE:
		input.ServeEndpoint(service.ResourceService.DeleteResource, responseWriter, request, true)
		break
	case constanta.RequestGET:
		input.ServeEndpoint(service.ResourceService.GetResource, responseWriter, request, true)
		break
	}
}
