FROM golang:1.17-alpine AS Build

COPY . /code
WORKDIR /code

RUN go env -w GO111MODULE="on" \
    && go build -o easyhttpd main.go

FROM alpine:3.14

COPY --from=Build /code/easyhttpd /usr/local/bin/easyhttpd

RUN chmod +x /usr/local/bin/easyhttpd \
    && addgroup www \
    && adduser -D -H -G www http \
    && mkdir -p /srv/www \
    && chown -R http.www /srv/www

USER http

EXPOSE 8008

ENTRYPOINT ["easyhttpd"]
CMD ["-r", "/srv/www"]