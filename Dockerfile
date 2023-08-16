FROM alpine

WORKDIR app

COPY duolingo .

RUN chmod +x duolingo

CMD ["/app/duolingo"]
