FROM golang:1.24.1 AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go build -o main ./cmd/v1

FROM debian:bookworm-slim

WORKDIR /app

COPY --from=builder /app/main .

EXPOSE $SERVER_PORT
CMD ["./main"]


