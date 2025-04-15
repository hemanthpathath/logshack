# Stage 1: Build the Go application
FROM golang:1.22.1-alpine as builder

# Install build dependencies
RUN apk add --no-cache build-base git

# Set working directory
WORKDIR /app

# Copy and download deps
COPY go.mod ./
RUN go mod download

# Copy the source code
COPY . .

# Build the Go binary for Linux (amd64 to ensure cross-compat)
RUN GOOS=linux GOARCH=amd64 go build -o main ./cmd/server/main.go

# Stage 2: Minimal runtime container
FROM alpine:latest

RUN apk add --no-cache tzdata

WORKDIR /app

# Copy the compiled Go binary
COPY --from=builder /app/main .

# Expose port
EXPOSE 8080

# Start the server
ENTRYPOINT ["/app/main"]
