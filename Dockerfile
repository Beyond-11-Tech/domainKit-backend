FROM golang:latest as builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . ./

RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main .

FROM alpine:latest
RUN apk --no-cache add ca-certificates curl bind

WORKDIR /root/

COPY --from=builder /app/main ./

# Expose the application port
EXPOSE 8080

# Add a health check to call the /v1/health endpoint
HEALTHCHECK --interval=30s --timeout=5s --start-period=5s --retries=3 \
  CMD curl --fail http://localhost:8080/v1/health || exit 1

ENTRYPOINT ["./main"]