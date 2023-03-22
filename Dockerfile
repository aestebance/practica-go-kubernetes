# Using golang 1.20.2 with alpine 3.17
FROM golang:1.20.2-alpine AS builder
# Create user
RUN adduser -D -g '' aestebance
# Create workspace
WORKDIR /opt/app/
COPY go.mod go.sum ./
# Fetch dependencies
RUN go mod download
RUN go mod verify
# Copy the source code
COPY . .
# Build the binary
RUN CGO_ENABLED=0 GOOS=linux go build -ldflags="-w -s" -a -installsuffix cgo -o /go/bin/server ./cmd/server
RUN CGO_ENABLED=0 GOOS=linux go build -ldflags="-w -s" -a -installsuffix cgo -o /go/bin/client ./cmd/client

# Build a small image
FROM alpine:3.17
