FROM golang:alpine as build

WORKDIR /buildapp
COPY . .
RUN go build -o goapp main.go

FROM alpine:3.22
WORKDIR /app
COPY --from=build /buildapp/goapp /app/goapp
COPY .env /app/.env
COPY data.sql /app/data.sql

ENTRYPOINT [ "/app/goapp" ]