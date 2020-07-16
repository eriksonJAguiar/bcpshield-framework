FROM golang

ADD . /go/src/privacy-mod

RUN go install privacy-mod

ENTRYPOINT /go/bin/privacy-mod

VOLUME [ "/home/debian/" ]

CMD [ "go run", "main.go" ]

EXPOSE 5000



