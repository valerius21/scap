# Use the official Go image as the base image
FROM golang:1.20-alpine

# Set the working directory inside the container
WORKDIR /app

# Copy the Go module files
COPY go.mod go.sum ./

# Download the Go module dependencies
RUN go mod download

# Copy the entire project directory into the container
COPY . .

# Build the Go application
RUN go build -o app ./cmd/main.go

# Set the command to run when the container starts
CMD ["./app", "--nsqHost=0.0.0.0"]
