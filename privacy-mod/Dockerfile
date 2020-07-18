FROM golang

ADD . /go/src/github.com/privacy-mod


RUN go install github.com/privacy-mod

ENTRYPOINT /go/bin/privacy-mod

EXPOSE 5000





