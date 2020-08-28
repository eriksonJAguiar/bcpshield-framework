FROM golang

ADD . /go/src/github.com/observer-mod


RUN go install github.com/observer-mod

ENTRYPOINT /go/bin/observer-mod

EXPOSE 6000