FROM golang:latest as builder
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . ./
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main .

FROM alpine:latest
RUN apk --no-cache add ca-certificates
WORKDIR /root/
COPY --from=builder /app/main ./
EXPOSE 8080
ENTRYPOINT [ "./main", "-apiKey", "test", "-appKey=test" ]



# FROM golang:1.24.2 AS build

# WORKDIR /app

# # Copy the Go module files
# COPY go.mod .
# COPY go.sum .

# # Download the Go module dependencies
# RUN go mod download

# COPY . .

# RUN go build -o apikit .

# FROM alpine:latest AS run
# WORKDIR /app

# # Copy the application executable from the build image
# COPY --from=build /app/apikit apikit

# ENV webKey=
# ENV appKey=

# EXPOSE 8080
# ENTRYPOINT [ "./main", "-apiKey=test", "-appKey=test" ]
