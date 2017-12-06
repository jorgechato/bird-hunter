package main

import "strings"

type Client struct {
	Username string
	Password string
	Tags     []string
}

type Likes struct {
	InPhoto int
	Number  int
}

type YamlOpt struct {
	Client Client `yaml:"credentials"`
	Daemon Daemon `yaml:"daemon"`
}

type Daemon struct {
	InPhoto  int      `yaml:"min-likes-x-photo"`
	Interval string   `yaml:"interval"`
	Likes    int      `yaml:"likes"`
	Tags     []string `yaml:"tags"`
}

func NewClient(user string, pass string, tags string) {
	client = Client{
		Username: user,
		Password: pass,
		Tags:     strings.Split(tags, ","),
	}
}

func setClient(user string, pass string, tags []string) {
	client = Client{
		Username: user,
		Password: pass,
		Tags:     tags,
	}
}

func NewLikes(number int, min int) {
	likes = Likes{
		InPhoto: min,
		Number:  number,
	}
}
