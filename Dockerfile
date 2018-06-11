FROM alpine:3.7

RUN mkdir /server

WORKDIR /server

ADD build build

ADD public public

ADD config config

ADD app app

EXPOSE 8000

ENV CONFIG_PATH=./config/config.json

CMD ["build/server-linux"]
