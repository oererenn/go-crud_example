# Use the official Golang Alpine image as the base image
FROM golang:alpine

# Set the working directory to /app
WORKDIR /app

# Copy the current directory contents into the container at /app
COPY . .

# Install the necessary packages for building the Go app
RUN apk add --no-cache git

# Build the Go app
RUN go build -o main ./cmd/comment-service/

# Specify the command to run when the container starts
CMD ["./main"]
