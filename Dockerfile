FROM scratch

ADD build/server-linux build/server-linux
ADD public public

ENV PORT=8000

EXPOSE 8000

CMD ["build/server-linux"]
