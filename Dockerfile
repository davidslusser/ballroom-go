# Stage 1: build
FROM golang:1.23-alpine AS builder

WORKDIR /app

# Copy go.mod and go.sum first for caching
COPY go.mod go.sum ./

# Download dependencies
RUN go mod download

# Copy the rest of the source code
COPY . .

# Build the app
RUN go build -o ballroom-go .

# Stage 2: run
FROM alpine:latest

WORKDIR /root/

# Copy the binary from builder
COPY --from=builder /app/ballroom-go .

# Expose port
EXPOSE 8000

# Run the binary
CMD ["./ballroom-go"]
