FROM golang:1.27.1-alpine AS builder

WORKDIR /app

COPY backend/go.mod backend/go.sum ./
RUN go mod download

COPY backend/ .

RUN CGO_ENABLED=0 go build -ldflags="-s -w" -o stockflow-api ./cmd/server

FROM alpine:3.20
WORKDIR /app
COPY --from=builder /app/stockflow-api .
RUN adduser -D -H -u 10001 appuser
USER appuser
CMD ["./stockflow-api"]