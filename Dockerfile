# syntax=docker/dockerfile:1

## Build
FROM golang:1.19.2-bullseye AS builder

WORKDIR /tempest-user-service

ENV GO111MODULE=on

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY config/*.yaml ./

COPY . .
COPY *.go ./

RUN CGO_ENABLED=0 go build -o /tempest-user-service

## Deploy
FROM scratch

WORKDIR /

COPY --from=builder /tempest-user-service ./

EXPOSE 8080

USER nonroot:nonroot


ENTRYPOINT ["/tempest-user-service"]