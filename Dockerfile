# syntax=docker/dockerfile:1

## Build
FROM golang:1.16-buster AS build

WORKDIR /app

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY *.go ./
COPY . $SRC_DIR

RUN go build -o /tempest-user-service

## Deploy
FROM gcr.io/distroless/base-debian10

WORKDIR /

COPY --from=build /tempest-user-service /tempest-user-service

EXPOSE 8080

USER nonroot:nonroot

ENTRYPOINT ["/tempest-user-service"]