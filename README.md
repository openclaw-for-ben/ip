# IP

Simple HTTP service that returns the client's IP address.

[![Test](https://github.com/bborbe/ip/actions/workflows/test.yml/badge.svg)](https://github.com/bborbe/ip/actions/workflows/test.yml)
[![Go Reference](https://pkg.go.dev/badge/github.com/bborbe/ip.svg)](https://pkg.go.dev/github.com/bborbe/ip)

## Features

- Returns client IP from `X-Forwarded-For`, `X-Real-IP`, or `RemoteAddr`
- Prometheus metrics at `/metrics`
- Health check at `/healthz`
- Readiness check at `/readiness`
- Graceful shutdown

## Installation

```bash
go install github.com/bborbe/ip@latest
```

## Usage

```bash
ip -listen=:8080
```

### Environment Variables

| Variable | Default | Description |
|----------|---------|-------------|
| `LISTEN` | `:8080` | Address to listen on |
| `SENTRY_DSN` | | Optional Sentry DSN for error tracking |

### Docker

```bash
docker run -p 8080:8080 bborbe/ip:latest
```

## API

### GET /

Returns the client's IP address as plain text.

**Headers checked (in order):**
1. `X-Forwarded-For` (first IP)
2. `X-Real-IP`
3. `X-Remote-Addr`
4. `RemoteAddr`

**Example:**
```bash
curl http://localhost:8080/
# 192.168.1.1
```

### GET /healthz

Health check endpoint. Returns `OK`.

### GET /readiness

Readiness check endpoint. Returns `OK`.

### GET /metrics

Prometheus metrics endpoint.

## Requirements

- Go 1.26+

## License

BSD-style license. See [LICENSE](LICENSE) for details.
