package main

import (
	"flag"

	"gopkg.in/ahmdrz/goinsta.v1"
)

var (
	//Flag sets
	foreverCommand = flag.NewFlagSet(
		"forever",
		flag.ExitOnError,
	)

	popularCommand = flag.NewFlagSet(
		"popular",
		flag.ExitOnError,
	)

	tagCommand = flag.NewFlagSet(
		"tag",
		flag.ExitOnError,
	)

	//Custom Ptr
	httpAddrPtr = foreverCommand.String(
		"http",
		"127.0.0.1:3999",
		"HTTP service address (e.g., `127.0.0.1:3999`)",
	)

	listTagsPtr = tagCommand.String(
		"list",
		"",
		"List of tags witch will be liked split by , (e.g., `tag1,tag2,tag3`)",
	)

	//Common
	usernamePtr string
	passwordPtr string

	//Popular & Tag
	likesPtr    int
	minlikesPtr int

	insta *goinsta.Instagram

	client = Client{}
	likes  = Likes{}
)
