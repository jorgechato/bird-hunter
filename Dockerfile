FROM golang:latest

RUN mkdir /birdhunter
WORKDIR /birdhunter
ADD . /birdhunter
RUN go install /birdhunter/cmd/birdhunter

ENTRYPOINT /go/bin/birdhunter forever -notify

EXPOSE 8080
