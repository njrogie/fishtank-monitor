# syntax=docker/dockerfile:1

FROM golang:latest

WORKDIR /app

RUN go mod download

COPY *.go ./

RUN go build -o /fishtank-container

EXPOSE 8080

CMD [ "/docker-gs-ping" ]
