FROM alpine:3.7

WORKDIR /go

ADD . /go

EXPOSE 8000

ENV CONFIG_PATH=./config/config.json

ENTRYPOINT ["./app"]
