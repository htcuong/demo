FROM golang:1.19-alpine AS base

ENV GOPRIVATE=github.com/htcuong

# Install git, tooling
RUN apk update && apk add --no-cache git make bash

# Set the current working directory inside the container
WORKDIR /root


# Copy go mod and sum files
COPY go.mod go.sum ./

# Download all dependencies. Dependencies will be cached if the go.mod and the go.sum files are not changed
RUN go mod download

# Copy the source from the current directory to the working Directory inside the container
COPY . .
RUN ls -alh

FROM base AS tester
RUN CGO_ENABLED=0 go test --tags=unit ./...

FROM base AS builder
# Build the Go app
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o server cmd/main.go || exit 1
RUN ls -alh

# Start a new stage from scratch
FROM alpine:latest

ENV USER=demo UID=1001 GID=1001 \
    WD=/app

WORKDIR ${WD}

# Copy the Pre-built binary file from the previous stage. Observe we also copied the .env file
COPY --from=builder /root/server .
COPY --from=builder /root/config.yml .
RUN ls -alh
# Expose port 8080 to the outside world
EXPOSE 8080

# Add a health check
# HEALTHCHECK --interval=30s --timeout=10s --start-period=5s --retries=3 CMD curl --fail http://localhost:8080/health || exit 1

# Command to run the executable
ENTRYPOINT [ "./server" ]
