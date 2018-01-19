# use golang:1.9, will call this first image 'builder'
FROM golang:1.9 AS builder
WORKDIR /go/src/hellogo
COPY . .
RUN go get github.com/tools/godep && \
        godep get && \
        go build -o hellogo .

# define new image, based in debian
FROM debian
WORKDIR /
# copy the result of the compilation from 'builder'
COPY --from=builder /go/src/hellogo/hellogo /go/bin/hellogo
USER "nobody"
CMD ["/go/bin/hellogo"]
# smaller image, we do not need a compiler in our image!
