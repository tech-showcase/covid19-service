package covid19

import (
	stdopentracing "github.com/opentracing/opentracing-go"
	generalEndpoint "github.com/tech-showcase/covid19-service/endpoint"
	"github.com/tech-showcase/covid19-service/middleware"
	"github.com/tech-showcase/covid19-service/service"
)

type (
	Endpoint struct {
		Get generalEndpoint.HTTPEndpoint
	}
)

func NewCovid19Endpoint(svc service.Covid19Service, tracer stdopentracing.Tracer) Endpoint {
	covid19Endpoint := Endpoint{}

	getCovid19Endpoint := makeGetCovid19Endpoint(svc)
	getCovid19Endpoint = middleware.ApplyTracerClient("getCovid19-endpoint", getCovid19Endpoint, tracer)
	covid19Endpoint.Get = generalEndpoint.HTTPEndpoint{
		Endpoint: getCovid19Endpoint,
		Decoder:  decodeGetCovid19Request,
		Encoder:  encodeResponse,
	}

	return covid19Endpoint
}
