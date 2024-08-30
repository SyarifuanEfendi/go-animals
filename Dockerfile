# Stage 1: Build the Go binary
FROM golang:1.22.1-alpine as builder

# Set the working directory
WORKDIR /app

# Copy go.mod and go.sum and download dependencies
COPY go.mod go.sum ./
RUN go mod download

# Copy the source code
COPY . .

# Build the Go binary
RUN CGO_ENABLED=0 GOOS=linux go build -o main .

# Stage 2: Create a minimal runtime image
FROM alpine:latest

# Set the working directory
WORKDIR /app

# Copy the built binary and environment file from the builder stage
COPY --from=builder /app/main .
COPY .env.example .env
COPY .env ./app/main


# Expose the application port
EXPOSE 8080

# Run the binary
CMD ["./main"]
