# Build stage
FROM quay.io/projectquay/golang:1.23 AS builder

WORKDIR /app

# Copy dependencies
COPY go.mod go.sum ./
RUN go mod download

# Copy source code
COPY . .

# Build the Go app
RUN go build -o main ./cmd/shinko

# Final stage - Minimal runtime image from Quay
FROM quay.io/centos/centos:stream9

WORKDIR /app

# Copy only the compiled binary
COPY --from=builder /app/main /app/index.html ./
COPY --from=builder /app/metrics ./metrics/

EXPOSE 8080

CMD ["./main"]