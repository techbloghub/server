# syntax=docker/dockerfile:1

# Build the application from source
FROM golang:1.22.5 AS build-stage

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod tidy

COPY . .

RUN go build -o main ./cmd/main.go

RUN chmod +x main

EXPOSE 8080

ENTRYPOINT ["./main"]
