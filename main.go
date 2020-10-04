package main

import (
	"fmt"
	"github.com/tech-showcase/covid19-service/cmd"
	"github.com/tech-showcase/covid19-service/config"
)

func init() {
	var err error
	config.Instance, err = config.Parse()
	if err != nil {
		panic(err)
	}
}

func main() {
	fmt.Println("Hi, I am Covid-19 Service!")

	err := cmd.Execute()
	if err != nil {
		panic(err)
	}
}
