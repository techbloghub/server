FROM golang:1.22.5-alpine

ENV GO111MODULE=on

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download
COPY . .

RUN go build -o main .

EXPOSE 8080


CMD ["./main"]
