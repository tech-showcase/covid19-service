package main

import "github.com/tech-showcase/covid19-service/config"

func init() {
	var err error
	config.Instance, err = config.Parse()
	if err != nil {
		panic(err)
	}
}

func main() {

}
