# syntax=docker/dockerfile:1

FROM golang:latest

WORKDIR /app

RUN go mod download

COPY source/*.go ./
COPY source/go.mod ./
COPY source/go.sum ./
COPY source/*.go ./

RUN go build -o fishtank-container

EXPOSE 8080

CMD [ "/docker-gs-ping" ]
