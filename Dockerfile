FROM alpine

ADD build build

ADD public public

ENV PORT=8000

EXPOSE 8000

CMD ["build/server-linux"]
