FROM golang:latest

WORKDIR /go/src/github.com/jorgechato/birdhunter
ADD . .

RUN go get -v ./...
RUN go install ./cmd/birdhunter

ENTRYPOINT /go/bin/birdhunter forever

EXPOSE 8080
