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
RUN CGO_ENABLED=0 GOOS=linux go build -ldflags="-w -s" -a -installsuffix cgo -o /go/bin/character ./cmd/character
RUN CGO_ENABLED=0 GOOS=linux go build -ldflags="-w -s" -a -installsuffix cgo -o /go/bin/server ./cmd/server
RUN CGO_ENABLED=0 GOOS=linux go build -ldflags="-w -s" -a -installsuffix cgo -o /go/bin/client ./cmd/client
RUN CGO_ENABLED=0 GOOS=linux go build -ldflags="-w -s" -a -installsuffix cgo -o /go/bin/files ./cmd/files
RUN CGO_ENABLED=0 GOOS=linux go build -ldflags="-w -s" -a -installsuffix cgo -o /go/bin/namespace ./cmd/namespace

# Build a small image
FROM alpine:3.17
LABEL language="golang"
LABEL org.opencontainers.image.source = https://github.com/aestebance/erase-una-vez-2
COPY --from=builder /etc/passwd /etc/passwd
COPY --from=builder /go/bin/chracter ./usr/local/bin/character
COPY --from=builder /go/bin/server ./usr/local/bin/server
COPY --from=builder /go/bin/client ./usr/local/bin/client
COPY --from=builder /go/bin/files ./usr/local/bin/files
COPY --from=builder /go/bin/namespace ./usr/local/bin/namespace

USER aestebance

ENTRYPOINT ["server"]
