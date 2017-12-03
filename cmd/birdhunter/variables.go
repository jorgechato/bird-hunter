package main

import (
	"flag"

	"gopkg.in/ahmdrz/goinsta.v1"
)

var (
	httpAddr = flag.String(
		"http",
		"127.0.0.1:8080",
		"HTTP service address (e.g., `127.0.0.1:3999`)",
	)

	forever = flag.Bool(
		"forever",
		false,
		"Run forever the boot (DEFAULT `false`)",
	)

	tags = flag.String(
		"tags",
		"",
		"List of tags witch will be liked split by , (e.g., `tag1,tag2,tag3`)",
	)

	user = flag.String(
		"user",
		"",
		"Instagram `username`",
	)

	pass = flag.String(
		"pass",
		"",
		"Instagram `password`",
	)

	nlikes = flag.Int(
		"nlikes",
		60,
		"Maximum number of likes per hour (e.g., `60`)",
	)

	minlikes = flag.Int(
		"likesxpic",
		60,
		"Minimum number of likes a image needs to have to like it (e.g., `60`)",
	)

	insta *goinsta.Instagram

	client = Client{}
	likes  = Likes{}
)
