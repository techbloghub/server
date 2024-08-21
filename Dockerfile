FROM golang:1.20-alpine

RUN apk add --no-cache gcc musl-dev

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN CGO_ENABLED=1 GOOS=linux go build -o /techbloghub-server

FROM alpine:latest
COPY --from=0 /techbloghub-server /techbloghub-server
EXPOSE 8080
ENTRYPOINT ["/techbloghub-server"]
