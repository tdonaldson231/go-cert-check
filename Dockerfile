# Stage 1: Build the Go binary
FROM golang:1.20 AS builder

WORKDIR /app
COPY go.mod ./
RUN go mod tidy
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -o ssl-check ./cmd/ssl-check

# Stage 2: Use Alpine and include CA certificates
FROM alpine:latest
# Install certificates to trust SSL connections
RUN apk --no-cache add ca-certificates
WORKDIR /root/
COPY --from=builder /app/ssl-check .

# Command to run the binary
ENTRYPOINT ["./ssl-check"]
