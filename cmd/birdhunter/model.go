package main

import "strings"

type Client struct {
	User     string
	Password string
	Tags     []string
}

type Likes struct {
	Minimum int
	NLikes  int
}

func model(user string, pass string, tags string) {
	client = Client{
		User:     user,
		Password: pass,
		Tags:     strings.Split(tags, ","),
	}
}

func Init() {
	model(*user, *pass, *tags)

	likes = Likes{
		Minimum: *minlikes,
		NLikes:  *nlikes,
	}
}
