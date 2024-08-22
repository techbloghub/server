FROM golang:alpine AS builder

ENV GO111MODULE=on \
    CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64

WORKDIR /build

COPY go.mod go.sum ./cmd/main.go ./

RUN go mod download
RUN go build -o techbloghub-server .

WORKDIR /dist

RUN cp /build/techbloghub-server .

FROM scratch

COPY --from=builder /dist/techbloghub-server /techbloghub-server

EXPOSE 8080
ENTRYPOINT ["/techbloghub-server"]
