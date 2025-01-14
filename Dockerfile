FROM golang:alpine

# Install dependencies
RUN apk update && apk upgrade && \
  apk add --no-cache bash git

# Set working directory
WORKDIR /app

# Copy go.mod and go.sum files to the container
COPY go.mod go.sum ./

# Download dependencies
RUN go mod download

# Install air (for hot reloading)
RUN go install github.com/air-verse/air@latest

# Copy the rest of the application code
COPY . .

# Copy air.toml configuration file
COPY .air.toml .air.toml

# Build the Go application
RUN go build -o main .

# Expose the port that your app will run on
EXPOSE 8888

# Run the app with Air
CMD ["air", "-c", ".air.toml"]
