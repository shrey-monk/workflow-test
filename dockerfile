# Build Stage
FROM golang:1.22 as builder

# Set the working directory
WORKDIR /app

# Copy go.mod and go.sum
COPY go.mod ./

# Download dependencies
RUN go mod download

# Copy the rest of the application code
COPY . .

# Build the application for amd64 architecture
RUN GOOS=linux GOARCH=amd64 go build -o myservice

# Final Stage (small base image)
FROM alpine:latest

# Set the working directory
WORKDIR /app

# Copy the binary from the builder stage
COPY --from=builder /app/myservice .

# Ensure the binary is executable
RUN chmod +x myservice

# Command to run the service
CMD ["./myservice"]
