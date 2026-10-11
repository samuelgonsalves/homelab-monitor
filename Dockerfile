# 1. Build stage using Go 1.27 on Alpine
FROM golang:1.27-alpine AS build
WORKDIR /src

# Leverage Docker cache for dependencies
COPY go.mod ./
RUN go mod download

# Copy the rest of the source code and build
COPY . .
RUN CGO_ENABLED=0 go build -o /homelab-monitor .

# 2. Final runtime stage
FROM alpine:3.20
RUN apk add --no-cache tzdata ca-certificates

# Copy the binary from the build stage
COPY --from=build /homelab-monitor /usr/local/bin/homelab-monitor

ENTRYPOINT ["homelab-monitor"]
