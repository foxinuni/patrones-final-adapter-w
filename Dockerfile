# Use the official Golang image from the Docker Hub
FROM golang:latest

# Set the Current Working Directory inside the container
WORKDIR /app

# Copy the source code into the container
COPY . .

# Build the Go app
RUN go build -o main .

# Expose port 8080
EXPOSE 8080

# Command to run the executable
CMD ["./main"]