package transport

import (
	"github.com/go-kit/kit/log"
	httptransport "github.com/go-kit/kit/transport/http"
	stdopentracing "github.com/opentracing/opentracing-go"
	"github.com/tech-showcase/covid19-service/endpoint"
	"github.com/tech-showcase/covid19-service/helper"
)

func MakeGetCovid19HTTPHandler(getCovid19Endpoint endpoint.HTTPEndpoint, tracer stdopentracing.Tracer, logger log.Logger) (handler *httptransport.Server) {
	var options []httptransport.ServerOption
	options = append(options, httptransport.ServerBefore(helper.HTTPToContext(tracer, "getCovid19-transport", logger)))

	handler = httptransport.NewServer(
		getCovid19Endpoint.Endpoint,
		getCovid19Endpoint.Decoder,
		getCovid19Endpoint.Encoder,
		options...,
	)

	return
}
