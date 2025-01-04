# Build Stage
FROM golang:1.23.4 AS builder
ENV CGO_ENABLED=0 \
    GOOS=linux

WORKDIR /app

# Download Go modules
COPY go.mod go.sum ./
RUN go mod download

# Copy the source code. Note the slash at the end, as explained in
# https://docs.docker.com/reference/dockerfile/#copy
COPY . ./

# Disable CGO and set target OS to Linux for a portable static binary
RUN go build -o /techbloghub-server ./cmd/main.go


# Final Stage
FROM alpine:latest AS deployment

# copy only built binaries
COPY --from=builder /techbloghub-server ./

# Optional:
# To bind to a TCP port, runtime parameters must be supplied to the docker command.
# But we can document in the Dockerfile what ports
# the application is going to listen on by default.
# https://docs.docker.com/reference/dockerfile/#expose
EXPOSE 8080

# Run
CMD ["./techbloghub-server"]
