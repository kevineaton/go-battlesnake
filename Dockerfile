FROM golang:1.16 as base

ADD ./ /go/src/github.com/kevineaton/go-battlesnake
WORKDIR /go/src/github.com/kevineaton/go-battlesnake

RUN go build .

FROM busybox:glibc
WORKDIR /go/src/github.com/kevineaton/go-battlesnake
COPY --from=base /go/src/github.com/kevineaton/go-battlesnake/go-battlesnake .
COPY --from=base /etc/ssl/certs /etc/ssl/certs
CMD ["./go-battlesnake"]