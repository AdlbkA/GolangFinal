# Stage 1 - Build the Go binary
FROM golang:1.23.2 AS builder

WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o AuthService ./cmd/app/

# Stage 2 - Run the binary
FROM alpine:latest

WORKDIR /app
COPY --from=builder /app/AuthService .
COPY internal/app/config/.env internal/app/config/.env



EXPOSE 8080

CMD ["./AuthService"]