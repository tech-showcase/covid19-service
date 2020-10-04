## COVID-19 SERVICE

### Description
This repo contains project that act as a **service**.
This service is part of a big system. 
The whole system will be used to present **micro-services without an orchestrator**.

### Features
- Serve Covid-19 data through HTTP.

### How to run
#### Docker
- Install docker
- Create config file `.env` under root dir which contains following content
```
ADDRESS=https://api.covid19api.com
```
- Fill env var `CONFIG_FILEPATH` with directory path where config file is contained
- Fill env var `CONFIG_FILENAME` with the name of config file (e.g `.env`)
- Build and run docker image as below
```shell script
$ docker build -t covid19-service .
$ docker run -p 8083:8080 covid19-service
```

### Tech / Dependency
- [Go kit - service](https://github.com/go-kit/kit)
- [Cobra - cli app](https://github.com/spf13/cobra)
- [Viper - config](https://github.com/spf13/viper)
- [Gorilla http mux - api](https://github.com/gorilla/mux)
