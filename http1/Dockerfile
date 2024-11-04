# Start from the official Golang image
FROM golang:1.20 as builder

# Set the Current Working Directory inside the container
WORKDIR /app

# Copy go.mod and go.sum files
COPY go.mod go.sum ./

# Download all dependencies
RUN go mod download

# Copy the source from the current directory to the Working Directory inside the container
COPY . .

# Build the Go app
RUN go build -o my-knative-function

# Start a new stage from scratch
FROM gcr.io/distroless/base-debian10

# Copy the Pre-built binary file from the previous stage
COPY --from=builder /app/my-knative-function /my-knative-function

# Command to run the executable
ENTRYPOINT ["/my-knative-function"]
