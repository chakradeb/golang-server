FROM alpine:3.7

RUN mkdir /server

WORKDIR /server

ADD build build
ADD public public
ADD config config

ENV PORT=8000
EXPOSE 8000

CMD ["build/server-linux"]
