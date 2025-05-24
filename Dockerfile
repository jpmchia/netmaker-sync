FROM golang:1.23-alpine AS builder

WORKDIR /app

# Copy go mod and sum files
COPY go.mod go.sum ./

# Download dependencies
RUN go mod download

# Copy source code
COPY . .

# Build the application
RUN CGO_ENABLED=0 GOOS=linux go build -o netmaker-sync .

# Use a smaller image for the final container
FROM alpine:latest

WORKDIR /app

# Install CA certificates for HTTPS requests
RUN apk --no-cache add ca-certificates

# Copy the binary from the builder stage
COPY --from=builder /app/netmaker-sync /app/netmaker-sync

# Copy the .env.example file (will be used if no .env is provided)
COPY .env.example /app/.env.example

# Create a non-root user and switch to it
RUN addgroup -S appgroup && adduser -S appuser -G appgroup
USER appuser

# Expose the API port
EXPOSE 8080

# Run the application
ENTRYPOINT ["/app/netmaker-sync"]
CMD ["serve"]
