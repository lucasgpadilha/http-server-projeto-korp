# syntax=docker/dockerfile:1

FROM golang:1.23-alpine AS builder

WORKDIR /src

COPY go.mod ./
COPY main.go ./

RUN CGO_ENABLED=0 GOOS=linux go build \
    -trimpath \
    -ldflags="-s -w" \
    -o /out/http-server-projeto-korp .

FROM alpine:3.20

LABEL org.opencontainers.image.title="http-server-projeto-korp" \
      org.opencontainers.image.description="HTTP server em Go para o desafio técnico Projeto Korp" \
      org.opencontainers.image.source="https://github.com/lucasgpadilha/http-server-projeto-korp"

RUN addgroup -S -g 10001 app && \
    adduser -S -D -H -u 10001 -G app app

WORKDIR /app

COPY --from=builder /out/http-server-projeto-korp /app/http-server-projeto-korp

EXPOSE 8080

USER app

HEALTHCHECK --interval=30s --timeout=3s --start-period=5s --retries=3 \
  CMD wget -qO- http://127.0.0.1:8080/healthz || exit 1

ENTRYPOINT ["/app/http-server-projeto-korp"]