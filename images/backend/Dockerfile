FROM golang:latest

RUN apt-get update && apt-get install docker.io -y
RUN go get github.com/docker/docker/client
RUN go get github.com/ant0ine/go-json-rest/rest