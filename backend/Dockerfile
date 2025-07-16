# Compilation
FROM golang:1.24.3-alpine AS builder
RUN apk add --no-cache git
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN go build ./cmd/server.go

# Serveur web
FROM alpine:latest
RUN apk add --no-cache ca-certificates
WORKDIR /app
COPY --from=builder /app/server .
COPY .env .
EXPOSE 5001
CMD ["./server"]

