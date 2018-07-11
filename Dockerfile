FROM golang

LABEL maintainer="Narongrit Kanhanoi narongrit@3dsinteractive.com"

ENV GOPATH=/go/src

RUN go get github.com/line/line-bot-sdk-go/linebot
RUN apt-get update
RUN apt-get -y install imagemagick

COPY src /go/src
WORKDIR /go/src

ENTRYPOINT ["/go/src/start.sh"]