# Build stage
FROM golang:1.25-alpine AS builder

WORKDIR /app

# Copy go.mod first for better layer caching
COPY go.mod ./

# Download dependencies (none currently, but good practice)
RUN go mod download

# Copy source code
COPY *.go ./

# Build the binary
RUN CGO_ENABLED=0 GOOS=linux go build -o httpecho

# Runtime stage
FROM alpine:latest

RUN apk --no-cache add ca-certificates

WORKDIR /root/

# Copy the binary from builder
COPY --from=builder /app/httpecho .

EXPOSE 8080

CMD ["./httpecho"]
