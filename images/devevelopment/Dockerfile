FROM golang:1.14 AS build-stage

ADD . $GOPATH/src/github.com/naoto67/clean-architecture
WORKDIR $GOPATH/src/github.com/naoto67/clean-architecture

RUN GOOS=linux CGO_ENABLED=0 go build --tags local cmd/server/main.go
RUN mv main /usr/local/bin

FROM alpine:3.10.3
COPY --from=build-stage /usr/local/bin/main /app/
WORKDIR /app
ADD ./env /app/env

RUN chmod +x /app/main
