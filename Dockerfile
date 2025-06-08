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

ARG APP_KEY
ARG WEB_KEY

ENV appKey ${APP_KEY}
ENV webKey ${WEB_KEY}

EXPOSE 8080
ENTRYPOINT ./main -webKey=$webKey -appKey=$appKey