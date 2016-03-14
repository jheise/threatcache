FROM golang

RUN mkdir /go/src/threatcache

ADD threatcache.go /go/src/threatcache/threatcache.go
RUN go get threatcache
RUN go install threatcache

ENTRYPOINT /go/bin/threatcache

EXPOSE 8888
