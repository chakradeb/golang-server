FROM alpine:3.7

RUN mkdir /app

WORKDIR /app

ADD build build

ADD public public

ADD config config

EXPOSE 8000

ENV CONFIG_PATH=./config/config.json

CMD ["build/server-linux"]
