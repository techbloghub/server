# Step 1: Build the Go application using an official Go image
FROM golang:1.22.5-alpine AS builder

# Set the Current Working Directory inside the container
WORKDIR /app

# Copy go.mod and go.sum files
COPY go.mod go.sum ./

# Download all dependencies. Dependencies will be cached if the go.mod and go.sum files are not changed
RUN go mod download

# Copy the source code into the container
COPY . .

# Build the Go app for Linux with static linking
RUN CGO_ENABLED=0 GOOS=linux go build -o /techbloghub-server

# Step 2: Create a minimal image to run the application
FROM alpine:latest

# Copy the binary from the builder stage
COPY --from=builder /techbloghub-server /techbloghub-server

# Expose the port the app runs on
EXPOSE 8080

# Command to run the binary
ENTRYPOINT ["/techbloghub-server"]
