package transport

import (
	httptransport "github.com/go-kit/kit/transport/http"
	"github.com/tech-showcase/covid19-service/endpoint"
)

func MakeGetCovid19HTTPHandler(getCovid19Endpoint endpoint.HTTPEndpoint) (handler *httptransport.Server) {
	handler = httptransport.NewServer(
		getCovid19Endpoint.Endpoint,
		getCovid19Endpoint.Decoder,
		getCovid19Endpoint.Encoder,
		[]httptransport.ServerOption{}...,
	)

	return
}
