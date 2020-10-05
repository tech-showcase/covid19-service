package covid19

import (
	"context"
	"encoding/json"
	"github.com/go-kit/kit/endpoint"
	"github.com/tech-showcase/covid19-service/model"
	"github.com/tech-showcase/covid19-service/service"
	"net/http"
	"time"
)

type (
	GetCovid19Request struct {
		Country string    `json:"country"`
		Status  string    `json:"status"`
		From    time.Time `json:"from"`
		To      time.Time `json:"to"`
	}
	GetCovid19Response struct {
		model.Covid19Data
	}
)

func makeGetCovid19Endpoint(svc service.Covid19Service) endpoint.Endpoint {
	return func(_ context.Context, request interface{}) (interface{}, error) {
		req := request.(GetCovid19Request)

		covid19Data, err := svc.Get(req.Country, req.Status, req.From, req.To)
		if err != nil {
			return GetCovid19Response{}, err
		}
		return GetCovid19Response{Covid19Data: covid19Data}, nil
	}
}

func decodeGetCovid19Request(_ context.Context, r *http.Request) (interface{}, error) {
	fromStr := getQueryStringValue(r, "from")
	from, _ := parseDateTime(fromStr)

	toStr := getQueryStringValue(r, "to")
	to, _ := parseDateTime(toStr)

	req := GetCovid19Request{
		Country: getQueryStringValue(r, "country"),
		Status:  getQueryStringValue(r, "status"),
		From:    from,
		To:      to,
	}

	return req, nil
}

func encodeResponse(_ context.Context, w http.ResponseWriter, response interface{}) error {
	w.Header().Add("Content-Type", "application/json")
	return json.NewEncoder(w).Encode(response)
}

func getQueryStringValue(r *http.Request, key string) (value string) {
	if valueArr, ok := r.URL.Query()[key]; ok {
		value = valueArr[0]
	}

	return
}

func parseDateTime(timeStr string) (value time.Time, err error) {
	value, err = time.Parse(time.RFC3339Nano, timeStr)
	if err != nil {
		return
	}

	return
}
