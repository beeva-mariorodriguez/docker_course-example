FROM golang:1.9 AS builder
WORKDIR /go/src/hellogo
COPY main.go .
RUN go get github.com/tools/godep && \
        godep get && \
        CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o hellogo .

FROM scratch
WORKDIR /
COPY --from=builder /go/src/hellogo/hellogo /go/bin/hellogo
# scratch is _reallly_ empty, so we need a minimal passwd file to avoid running as root
COPY passwd.minimal /etc/passwd
USER "nobody"
CMD ["/go/bin/hellogo"]

