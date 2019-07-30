FROM golang:1.12.6-alpine

WORKDIR /go/app

COPY ./app /go/app

RUN apk add --no-cache \
        alpine-sdk \
        git && \
    go get github.com/labstack/echo && \
    go get github.com/labstack/echo/middleware && \ 
    go get github.com/pilu/fresh

CMD [ "fresh" ]