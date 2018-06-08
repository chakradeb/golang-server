FROM golang:1.6

WORKDIR /app

ADD . app

EXPOSE 8000

RUN go get -u github.com/Masterminds/glide

RUN glide install

ENV CONFIG_PATH=./config/config.json

CMD ["go","run","main.go"]