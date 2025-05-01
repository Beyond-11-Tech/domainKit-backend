FROM golang:1.24.2 as builder

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go build .

FROM golang:1.24.2-bookworm as app

COPY --from=builder domainKit /usr/local/bin/app/

ENV webKey=
ENV appKey=
CMD domainKit -apiKey=${webKey} appKey=${appKey}