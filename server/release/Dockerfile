FROM alpine:3.11.5

RUN apk update
RUN apk add curl
RUN apk add logrotate
RUN apk add --no-cache tzdata && ln -snf /usr/share/zoneinfo/$TZ /etc/localtime

WORKDIR /go-app

ENV CONFIG_FILE_PATH /go-app/config
COPY inventory.o .
COPY config/. /go-app/config

ENTRYPOINT [ "/go-app/inventory.o" ]