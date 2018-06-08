FROM golang:1.10

RUN mkdir $GOPATH/src/app

WORKDIR $GOPATH/src/app

ADD . .

EXPOSE 8000

RUN curl https://glide.sh/get | sh

RUN glide install

ENV CONFIG_PATH=./config/config.json

CMD go run main.go
