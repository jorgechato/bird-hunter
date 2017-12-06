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
	Client Client `yaml:"client"`
	Daemon Daemon `yaml:"daemon"`
}

type Daemon struct {
	InPhoto int `yaml:"min-likes-x-photo"`
	Iter    Iter
}

type Iter struct {
	Tags  []string `yaml:"tags"`
	Likes int      `yaml:"likes"`
	Hour  string   `yaml:"hour"`
}

func NewClient(user string, pass string, tags string) {
	client = Client{
		Username: user,
		Password: pass,
		Tags:     strings.Split(tags, ","),
	}
}

func NewLikes(number int, min int) {
	likes = Likes{
		InPhoto: min,
		Number:  number,
	}
}
