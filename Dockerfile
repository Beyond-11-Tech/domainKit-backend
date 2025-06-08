FROM golang:1.24.2 as build

WORKDIR /app

# Copy the Go module files
COPY go.mod .
COPY go.sum .

# Download the Go module dependencies
RUN go mod download

COPY . .

RUN go build -o /app .
 
FROM alpine:latest as run

# Copy the application executable from the build image
COPY --from=build /app /app

ENV webKey=
ENV appKey=

WORKDIR /app
EXPOSE 8080
CMD app -apiKey=${webKey} appKey=${appKey}
