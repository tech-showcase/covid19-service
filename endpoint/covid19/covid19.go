package covid19

import (
	generalEndpoint "github.com/tech-showcase/covid19-service/endpoint"
	"github.com/tech-showcase/covid19-service/service"
)

type (
	Endpoint struct {
		Get generalEndpoint.HTTPEndpoint
	}
)

func NewCovid19Endpoint(svc service.Covid19Service) Endpoint {
	instance := Endpoint{}
	instance.Get = generalEndpoint.HTTPEndpoint{
		Endpoint: makeGetCovid19Endpoint(svc),
		Decoder:  decodeGetCovid19Request,
		Encoder:  encodeResponse,
	}

	return instance
}
