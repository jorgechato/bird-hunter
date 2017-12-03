FROM golang:latest

RUN mkdir /bird-hunter
WORKDIR /bird-hunter
ADD . /bird-hunter
RUN go install /bird-hunter/cmd/bird-hunter

ENTRYPOINT /go/bin/bird-hunter -f -http="0.0.0.0:8080"

EXPOSE 8080
