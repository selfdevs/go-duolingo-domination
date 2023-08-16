FROM alpine

WORKDIR app

COPY executable .

COPY assets ./assets

RUN chmod +x executable

CMD ["/app/executable"]
