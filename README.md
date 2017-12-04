# BirdHunter
[![Go Report Card](https://goreportcard.com/badge/github.com/jorgechato/birdhunter)](https://goreportcard.com/report/github.com/jorgechato/birdhunter)
[![Godoc](https://img.shields.io/badge/go-documentation-blue.svg)](https://godoc.org/github.com/jorgechato/birdhunter)

An easy tool to like pictures from Instagram without the official API

## Install & Run
#### Local
```zsh
$ go get github.com/jorgechato/birdhunter/cmd/birdhunter
$ birdhunter -user username -pass password -tags tag1,tag2,tag3 -nlikes 60 -likesxpic 20
$ birdhunter -user username -pass password -popular -nlikes 60 -likesxpic 20

```
#### Docker
```zsh
$ docker build -t birdhunter .

$ docker run -d \
-p 3999:3999 \
--name like-birdhunter\
birdhunter

# or

$ docker-compose up -d
```
## Variables
If you are running it with docker-compose, which I recommend, you can set the
variables in the .env file.
```zsh
GOPORT=3999
```
