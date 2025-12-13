# Use the official Golang image as the build environment
FROM golang:1.20-alpine AS builder

WORKDIR /app

# Copy go mod and sum files
COPY go.mod .
COPY go.sum .

# Download dependencies
RUN go mod download

# Copy the source code
COPY . .

# Build the application
RUN go build -o myapp ./cmd/main.go

# Use a minimal image for the final container
FROM alpine:latest
WORKDIR /root/

# Copy the binary from the builder
COPY --from=builder /app/myapp .

# Command to run
CMD ["./myapp"]
