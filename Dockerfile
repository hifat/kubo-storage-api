FROM golang:1.25-alpine AS builder

WORKDIR /app

RUN apk add --no-cache git make

# Copy go mod files
COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o /app/bin/storage-api ./cmd/r2/

# Runtime stage
FROM alpine:3.18

WORKDIR /app

# Install ca-certificates for HTTPS/TLS
RUN apk add --no-cache ca-certificates

# Create non-root user
RUN addgroup -g 1000 appuser && \
    adduser -D -u 1000 -G appuser appuser

# Copy binary from builder
COPY --from=builder /app/bin/storage-api .

# Create config directory
RUN mkdir -p /app/env && chown -R appuser:appuser /app

# Switch to non-root user
USER appuser

# Expose ports
EXPOSE 4000 9000

# Health check
HEALTHCHECK --interval=10s --timeout=5s --start-period=5s --retries=3 \
    CMD grpcurl -plaintext localhost:9000 grpc.health.v1.Health/Check || exit 1

# Run the application
ENTRYPOINT ["./storage-api"]