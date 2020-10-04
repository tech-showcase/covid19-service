package model

import (
	"encoding/json"
	"github.com/tech-showcase/covid19-service/helper"
	"io/ioutil"
	"net/http"
	"time"
)

type (
	Covid19Data []struct {
		Country     string    `json:"Country"`
		CountryCode string    `json:"CountryCode"`
		Province    string    `json:"Province"`
		City        string    `json:"City"`
		CityCode    string    `json:"CityCode"`
		Lat         string    `json:"Lat"`
		Lon         string    `json:"Lon"`
		Cases       int       `json:"Cases"`
		Status      string    `json:"Status"`
		Date        time.Time `json:"Date"`
	}

	covid19Repo struct {
		address string
	}
	Covid19Repo interface {
		Get(country, status string, from, to time.Time) (covid19Data Covid19Data, err error)
	}
)

func NewCovid19Repo(address string) Covid19Repo {
	instance := covid19Repo{}
	instance.address = address

	return &instance
}

func (instance *covid19Repo) Get(country, status string, from, to time.Time) (covid19Data Covid19Data, err error) {
	path := "/country/" + country + "/status/" + status
	endpoint, _ := helper.JoinURL(instance.address, path)

	req, err := http.NewRequest("GET", endpoint, nil)
	if err != nil {
		return
	}

	q := req.URL.Query()
	q.Add("from", from.Format(time.RFC3339Nano))
	q.Add("to", to.Format(time.RFC3339Nano))
	req.URL.RawQuery = q.Encode()

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return
	}

	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return
	}
	defer resp.Body.Close()

	err = json.Unmarshal(respBody, &covid19Data)
	if err != nil {
		return
	}

	return
}
