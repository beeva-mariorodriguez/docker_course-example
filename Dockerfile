FROM golang:1.9
WORKDIR /go/src/hellogo
COPY . .
RUN go get github.com/tools/godep && \
        godep get && \
        go build
USER "nobody"
CMD ["/go/bin/hellogo"]

