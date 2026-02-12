# syntax=docker/dockerfile:1

# Build stage
FROM golang:1.26.0-alpine AS builder

WORKDIR /app

# Download dependencies (with cache mount)
COPY go.mod go.sum ./
RUN --mount=type=cache,target=/go/pkg/mod \
    go mod download

# Build (with cache mount for build cache)
COPY . .
RUN --mount=type=cache,target=/go/pkg/mod \
    --mount=type=cache,target=/root/.cache/go-build \
    CGO_ENABLED=0 GOOS=linux go build -ldflags "-s -w" -o ip .

# Final stage
FROM alpine:3.21

RUN apk --no-cache add ca-certificates tzdata

COPY --from=builder /app/ip /ip

EXPOSE 8080

ENTRYPOINT ["/ip"]
CMD ["-listen=:8080", "-logtostderr", "-v=2"]
