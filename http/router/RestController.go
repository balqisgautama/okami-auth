package router

import (
	"github.com/balqisgautama/okami-auth/config"
	"github.com/balqisgautama/okami-auth/constanta"
	"github.com/balqisgautama/okami-auth/http/endpoint"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

func ApiController(port int) {
	handler := mux.NewRouter()

	handler.HandleFunc(setPath("/health"), endpoint.HealthEndpoint.CheckingHealth).Methods("GET", "OPTIONS")
	handler.HandleFunc(setPath("/sign"), endpoint.SignEndpoint.UserSign)
	handler.HandleFunc(setPath("/resource"), endpoint.ResourceEndpoint.Resource)
	handler.HandleFunc(setPath("/registration"), endpoint.RegistrationEndpoint.Registration).Methods("POST", "OPTIONS")
	handler.HandleFunc(setPath("/active/{"+constanta.VarToken+"}/{"+constanta.VarCode+"}"),
		endpoint.RegistrationEndpoint.ActivationUser).Methods("GET", "OPTIONS")
	handler.HandleFunc(setPath("/active/resend/{"+constanta.VarToken+"}/{"+constanta.VarCode+"}"),
		endpoint.RegistrationEndpoint.ActivationUserResend).Methods("GET", "OPTIONS")

	//util.Logger.Info("Hello World")
	//util.Logger.Error("Not able to reach blog.", zap.String("url", "localhost"))
	//util.Logger.Debug("logger debug", zap.String("debug", "try"))

	handler.Use(MiddlewareCORSOrigin)
	handler.Use(MiddlewareCheckUserToken)
	http.ListenAndServe(":"+strconv.Itoa(port), handler)
}

func setPath(path string) string {
	prefixPath := config.ApplicationConfiguration.GetServerPrefixPath()
	return "/" + prefixPath + path
}
