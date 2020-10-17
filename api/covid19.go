package api

import (
	"github.com/gorilla/mux"
	"github.com/tech-showcase/covid19-service/config"
	"github.com/tech-showcase/covid19-service/endpoint/covid19"
	"github.com/tech-showcase/covid19-service/helper"
	"github.com/tech-showcase/covid19-service/model"
	"github.com/tech-showcase/covid19-service/service"
	"github.com/tech-showcase/covid19-service/transport"
	"net/http"
)

func RegisterCovid19HTTPAPI(r *mux.Router) {
	configInstance := config.Instance
	tracerInstance := helper.TracerInstance
	loggerInstance := helper.LoggerInstance

	covid19Repo := model.NewCovid19Repo(configInstance.Address)
	covid19Service := service.NewCovid19Service(covid19Repo)
	covid19Endpoint := covid19.NewCovid19Endpoint(covid19Service, tracerInstance, loggerInstance)
	r.Handle("/covid19", transport.MakeGetCovid19HTTPHandler(covid19Endpoint.Get, tracerInstance, loggerInstance)).Methods(http.MethodGet)
}
