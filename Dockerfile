FROM golang:1.9-alpine
WORKDIR /go/src/hellogo
COPY . .
# golang:1.9-alpine has no git, so godep wont work
# we need to install git using apk
RUN \
    apk add --no-cache git && \
    go get github.com/tools/godep && \
    godep get && \
    go build
USER "nobody"
CMD ["/go/bin/hellogo"]

