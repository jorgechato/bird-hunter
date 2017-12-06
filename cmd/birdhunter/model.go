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

type YamlObj struct {
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
