# Use the specified Go image
FROM golang:1.22

# Set the working directory inside the container
WORKDIR /usr/src/app

# Copy go.mod and go.sum for dependency management (if applicable)
COPY go.mod go.sum ./

# Download dependencies using Go modules
RUN go mod download

# Copy the Go application source code into the container
COPY ../. .

# Build the Go binary (replace "main.go" with your actual entrypoint)
RUN CGO_ENABLED=0 go build -o main ./main.go

# Command to run the executable
CMD ["./main"]