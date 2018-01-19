FROM golang:1.9 AS builder
WORKDIR /go/src/hellogo
COPY . .
RUN go get github.com/tools/godep && \
        godep get && \
        CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o hellogo .

FROM scratch
WORKDIR /
COPY --from=builder /go/src/hellogo/hellogo /go/bin/hellogo
COPY passwd.minimal /etc/passwd
USER "nobody"
ENV REDISHOST="localhost" \
    REDISDB="0" \
    REDISPORT="6379"
CMD ["/go/bin/hellogo"]

