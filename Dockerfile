FROM golang:1.24 AS builder

WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN go build -a -o to-do-app ./cmd/api/main.go

FROM debian:stable-slim
WORKDIR /app
COPY --from=builder /app/to-do-app .
EXPOSE 8083
CMD ["./to-do-app"]