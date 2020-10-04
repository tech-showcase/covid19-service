package service

import (
	"github.com/tech-showcase/covid19-service/model"
	"time"
)

type (
	covid19Service struct {
		covid19Repo model.Covid19Repo
	}
	Covid19Service interface {
		Get(string, string, time.Time, time.Time) (model.Covid19Data, error)
	}
)

func NewCovid19Service(covid19Repo model.Covid19Repo) Covid19Service {
	instance := covid19Service{}
	instance.covid19Repo = covid19Repo

	return &instance
}

func (instance *covid19Service) Get(country, status string, from, to time.Time) (covid19Data model.Covid19Data, err error) {
	covid19Data, err = instance.covid19Repo.Get(country, status, from, to)
	if err != nil {
		return
	}

	return
}
