# syntax=docker/dockerfile:1

## Build
FROM golang:1.19.2-bullseye AS builder

WORKDIR /tempest-user-service

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY config/*.yaml ./

COPY . .
COPY *.go ./

RUN go build -o /tempest-user-service main.go

## Deploy
FROM scratch

WORKDIR /

COPY --from=builder /tempest-user-service ./

EXPOSE 8080

ENTRYPOINT ["/tempest-user-service"]