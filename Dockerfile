FROM golang:1.14-alpine3.11 AS builder

RUN apk update && apk add --no-cache git

WORKDIR $GOPATH/src/app/
COPY . .

RUN go mod tidy
RUN CGO_ENABLED=0 go build -o /go/bin/app

ENTRYPOINT ["/go/bin/app"]
