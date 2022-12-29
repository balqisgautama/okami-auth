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

	// registration
	handler.HandleFunc(setPath("/registration"),
		endpoint.RegistrationEndpoint.Registration).Methods("POST", "OPTIONS")
	handler.HandleFunc(setPath("/active/{"+constanta.VarToken+"}/{"+constanta.VarCode+"}"),
		endpoint.RegistrationEndpoint.ActivationUser).Methods("GET", "OPTIONS")
	handler.HandleFunc(setPath("/active/resend/{"+constanta.VarToken+"}/{"+constanta.VarCode+"}"),
		endpoint.RegistrationEndpoint.ActivationUserResend).Methods("GET", "OPTIONS")

	// login
	handler.HandleFunc(setPath("/sign/in/step1"),
		endpoint.SignInEndpoint.Step1).Methods("POST", "OPTIONS")
	handler.HandleFunc(setPath("/sign/in/step2"),
		endpoint.SignInEndpoint.Step2).Methods("POST", "OPTIONS")
	handler.HandleFunc(setPath("/sign/in/step3/{"+constanta.VarToken+"}"),
		endpoint.SignInEndpoint.Step3).Methods("GET", "OPTIONS")

	// crud resource
	handler.HandleFunc(setPath("/resource"), endpoint.ResourceEndpoint.Resource)

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
