# BirdHunter
[![Go Report Card](https://goreportcard.com/badge/github.com/jorgechato/birdhunter)](https://goreportcard.com/report/github.com/jorgechato/birdhunter)
[![Godoc](https://img.shields.io/badge/go-documentation-blue.svg)](https://godoc.org/github.com/jorgechato/birdhunter)

An easy tool to like pictures from Instagram without the official API

## Install & Run
#### Local
```zsh
$ go get github.com/jorgechato/birdhunter/cmd/birdhunter

$ birdhunter tag -u username -p password -list tag1,tag2,tag3 -likes 60 -min-likes-x-photo 20
$ birdhunter popular -u username -p password -likes 60 -min-likes-x-photo 20
```
#### Docker
```zsh
$ docker build -t birdhunter .

$ docker run -d \
-p 3888:8080 \
-v {PWD}/config.yaml:/config.yaml \
-v {PWD}/out:/out \
--name like-birdhunter \
birdhunter

# or

$ docker-compose up -d
```

## Use
```zsh
$ birdhunter -h

# output

popular
  -likes 60
        Maximum number of likes per hour (e.g., 60) (default 60)
  -min-likes-x-photo 60
        Minimum number of likes a image needs to have to like it (e.g., 60) (default 60)
  -p password
        Instagram password
  -u username
        Instagram username

tag
  -likes 60
        Maximum number of likes per hour (e.g., 60) (default 60)
  -list tag1,tag2,tag3
        List of tags witch will be liked split by , (e.g., tag1,tag2,tag3)
  -min-likes-x-photo 60
        Minimum number of likes a image needs to have to like it (e.g., 60) (default 60)
  -p password
        Instagram password
  -u username
        Instagram username

forever
  -configFile location
        Config file location (default "/config.yaml")
  -http localhost:8080
        HTTP service address (e.g., localhost:8080) (default "0.0.0.0:8080")
  -notify
        Will notify to your telegram chat the status of the bot

update
  -http localhost:8080
        HTTP service address (e.g., localhost:8080) (default "http://localhost:3888")
```

## Variables
Create a file ./config/config.yaml.
Variables in config.yaml file. In order to request a valid TOKEN create a bot and you will get the ID if you visit https://api.telegram.org/bot<TOKEN>/getUpdates.
```zsh
credentials:
  username: ""
  password: ""

notify:
  token: ""
  id: ""

daemon:
  min-likes-x-photo: 20
  interval: "@every 3h30m"
  likes: 60
  tags:
      - ""
      - ""
```
