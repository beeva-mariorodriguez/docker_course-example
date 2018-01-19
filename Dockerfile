FROM golang:1.9 AS builder
WORKDIR /go/src/hellogo
COPY . .
# alpine lacks some libraries so we compile hellogo statically
RUN go get github.com/tools/godep && \
        godep get && \
        CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o hellogo .

FROM alpine
WORKDIR /
COPY --from=builder /go/src/hellogo/hellogo /go/bin/hellogo
USER "nobody"
CMD ["/go/bin/hellogo"]

