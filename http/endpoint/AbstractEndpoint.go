package endpoint

import (
	"github.com/balqisgautama/okami-auth/constanta"
	"github.com/balqisgautama/okami-auth/dto/res"
	"github.com/balqisgautama/okami-auth/model"
	"github.com/balqisgautama/okami-auth/util"
	"go.uber.org/zap"
	"net/http"
	"time"
)

type AbstractEndpoint struct {
	FileName string
	FuncName string
}

func (input AbstractEndpoint) ServeEndpoint(serveFunction func(*http.Request) (res.APIResponse, map[string]string, error),
	responseWriter http.ResponseWriter, request *http.Request, sysadminAccess bool) {
	serveEndpoint(serveFunction, responseWriter, request, sysadminAccess)
}

func serveEndpoint(serveFunction func(*http.Request) (res.APIResponse, map[string]string, error),
	responseWriter http.ResponseWriter, request *http.Request, sysadminAccess bool) {
	serve(serveFunction, responseWriter, request, sysadminAccess)
}

func serve(serveFunction func(*http.Request) (res.APIResponse, map[string]string, error),
	responseWriter http.ResponseWriter, request *http.Request, sysadminAccess bool) {
	var err error
	var output res.APIResponse
	var header map[string]string

	defer func() {
		if r := recover(); r != nil {
			util.Logger.Info("recovery")
		} else {
			if err != nil {
				util.Logger.Info("server", zap.String("details", err.Error()))
			}
		}

		finish(responseWriter, err, output)
	}()

	sysadminUser := request.Header.Get(constanta.AccessHeaderNameConstanta)

	// only sysadmin can access
	if sysadminAccess && sysadminUser == "1" {
		output, header, err = serveFunction(request)
		return
	} else if sysadminAccess && sysadminUser != "1" {
		output = model.GenerateUnauthorizedError()
		return
	}

	output, header, err = serveFunction(request)

	setHeader(header, responseWriter)
}

func setHeader(header map[string]string, responseWriter http.ResponseWriter) {
	accessControlExpose := "Access-Control-Expose-Headers"
	exposeHeader := responseWriter.Header().Get(accessControlExpose)
	for key := range header {
		responseWriter.Header().Add(key, header[key])
		if exposeHeader == "" {
			exposeHeader = key
		} else {
			exposeHeader += ", " + key
		}
	}
	if exposeHeader != "" {
		responseWriter.Header().Set(accessControlExpose, exposeHeader)
	}
}

func finish(responseWriter http.ResponseWriter, err error, output res.APIResponse) {
	if err != nil {
		writeErrorResponse(responseWriter, err)
	} else {
		writeSuccessResponse(responseWriter, output)
	}
}

func writeErrorResponse(responseWriter http.ResponseWriter, err error) {
	responseWriter.WriteHeader(500)
	_, errorS := responseWriter.Write([]byte(err.Error()))
	if errorS != nil {
		util.Logger.Info("writeErrorResponse", zap.String("details", errorS.Error()))
	}
}

func writeSuccessResponse(responseWriter http.ResponseWriter, output res.APIResponse) {
	output.Timestamp = time.Now().Unix()
	bodyMessage := output.String()

	responseWriter.WriteHeader(200)
	_, errorS := responseWriter.Write([]byte(bodyMessage))
	if errorS != nil {
		util.Logger.Info("writeSuccessResponse", zap.String("details", errorS.Error()))
	}
}
