FROM golang:alpine AS build

WORKDIR /Build

COPY . .

RUN go build -o goapp main.go

FROM alpine:3.22

WORKDIR /app

COPY --from=build /Build/goapp /app/goapp

ENTRYPOINT ["/app/goapp"]