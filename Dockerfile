FROM golang:1.24.2 AS build

WORKDIR /app

# Copy the Go module files
COPY go.mod .
COPY go.sum .

# Download the Go module dependencies
RUN go mod download

COPY . .

RUN go build -o apikit .

FROM alpine:latest AS run
WORKDIR /app

# Copy the application executable from the build image
COPY --from=build /app/apikit apikit

ENV webKey=
ENV appKey=

EXPOSE 8080
ENTRYPOINT [ "apikit", "-apiKey=test", "-appKey=test" ]
