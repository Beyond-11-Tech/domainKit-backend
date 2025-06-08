FROM golang:latest as builder

# Set the working directory
WORKDIR /app

# Copy Go module files and download dependencies
COPY go.mod go.sum ./
RUN go mod download

# Copy the rest of the application code
COPY . ./

# Build the Go application
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main .

# Use a minimal base image for the final stage
FROM alpine:latest
RUN apk --no-cache add ca-certificates bind

# Set the working directory in the container
WORKDIR /root/

# Copy the built application from the builder stage
COPY --from=builder /app/main ./

# Expose the application port
EXPOSE 8080

# Command to run the application
ENTRYPOINT ["./main"]