FROM golang

ADD . /go/src/observer-mod

RUN go install observer-mod

ENTRYPOINT /go/bin/observer-mod

VOLUME [ "/home/debian/" ]

CMD [ "go run", "main.go" ]

EXPOSE 6000