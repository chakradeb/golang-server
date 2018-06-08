FROM golang:1.10

RUN mkdir /app

WORKDIR /app

ADD build build

ADD public public

ADD config config

EXPOSE 8000

RUN curl https://glide.sh/get | sh

CMD ["./build/server-linux"]
