FROM golang:1.17 as build

ADD . /app/
WORKDIR /app

RUN CGO_ENABLED=0 go build

FROM alpine:3.14

COPY --from=build /app/dummy /app/
WORKDIR /app

CMD ["/app/dummy"]
