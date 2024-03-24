# Stage 1: Build the Go application
FROM golang:1.22 AS builder

# Set the working directory
WORKDIR /app

# Copy the Go module files
COPY go.mod go.sum ./

# Download the Go module dependencies
RUN go mod download

# Copy the application source code
COPY . .

# Build the Go application
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o sh_account .

# Stage 2: Create the final runtime image
FROM alpine:latest

# Set the working directory
WORKDIR /app

# Copy the compiled binary from the builder stage
COPY --from=builder /app/sh_account .

# Expose the port on which the application will run (if needed)
EXPOSE 8080

# Run the application
CMD ["./sh_account", "serve"]