FROM golang:alpine AS builder

RUN apk add --no-cache --update alpine-sdk

ENV GOPATH=/go

ENV GOOS="linux"
ENV GOARCH="amd64"
ENV GO111MODULE=on

COPY . $GOPATH/src/github.com/s4kibs4mi/movie-pie
WORKDIR $GOPATH/src/github.com/s4kibs4mi/movie-pie

RUN rm go.sum

RUN go get .
RUN rm /go/pkg/mod/github.com/coreos/etcd@v3.3.10+incompatible/client/keys.generated.go
RUN cp ./hacks/keys.generated.go /go/pkg/mod/github.com/coreos/etcd@v3.3.10+incompatible/client/

RUN go build -v -o movie-pie
RUN mv movie-pie /go/bin/movie-pie

FROM alpine

WORKDIR /root

RUN mkdir -p /root/bin

COPY --from=builder /go/bin/movie-pie /usr/local/bin/movie-pie

ENTRYPOINT ["movie-pie"]
