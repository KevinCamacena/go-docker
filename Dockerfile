FROM golang:1.23 AS builder

WORKDIR /app

# Copy the Go modules manifests
COPY go.mod go.sum ./

# Download the module dependencies
RUN  go mod download 

# Copy the applicaction source code
COPY src/ .

# Build the application soruce code
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o go-api .

# Use a minimal image for the final container
FROM alpine:latest

# Set the working directory
WORKDIR /app

# Copy the compiled binary from the builder stage
COPY --from=builder /app/go-api .

EXPOSE 8080

# Command to run the application
CMD ["./go-api"]