FROM golang:1.8.0-alpine

RUN mkdir -p /usr/src/web/public

EXPOSE 8080

COPY ./ui/www /usr/src/web/public

COPY ./cmd/api/api /usr/src/app

CMD ["/usr/src/app"]