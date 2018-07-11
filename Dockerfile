FROM golang

LABEL maintainer="Narongrit Kanhanoi narongrit@3dsinteractive.com"

ENV GOPATH=/go/src

RUN apt-get update
RUN apt-get -y install imagemagick
RUN apt-get -y install libzbar-dev
RUN go get github.com/line/line-bot-sdk-go/linebot
RUN go get gopkg.in/bieber/barcode.v0

COPY src /go/src
WORKDIR /go/src

ENTRYPOINT ["/go/src/start.sh"]