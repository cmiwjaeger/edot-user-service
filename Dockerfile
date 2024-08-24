# Use the official Golang image from the Docker Hub
FROM golang:1.20 AS builder

# Set the Current Working Directory inside the container
WORKDIR /app

# Copy go.mod and go.sum files
COPY Makefile ./
COPY go.mod go.sum ./


# Download all dependencies. Dependencies will be cached if the go.mod and go.sum files are not changed
RUN go mod download

# Copy the source code into the container
COPY . .

# Build the Go app
RUN go build -o main .

# Start from a smaller image to keep the final image size small
FROM alpine:latest

# Copy the Pre-built binary file from the previous stage
COPY --from=builder /app/main .

# Expose port 3100 to the outside world
EXPOSE 3100

# Command to run the executable
CMD ["make build", "./main"]