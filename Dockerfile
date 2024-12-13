# Use the official Golang image from the Docker Hub
FROM golang:1.23-alpine

# Set the working directory to /app
WORKDIR /app

# Copy the current directory contents into the container at /app
COPY . .

# Build the main executable
RUN go build -o bin/main cmd/main.go

# Expose port 8080
EXPOSE 8080

# Run the main executable
ENTRYPOINT [ "./bin/main" ]