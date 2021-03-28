# Golang Server's Dockerfile
FROM golang:1.14.6-alpine

RUN mkdir -p /server/cmd
RUN mkdir -p /server/internal
RUN mkdir -p /server/pkg
COPY go.mod /server
COPY go.sum /server

RUN apk add --no-cache git

RUN go get github.com/lib/pq