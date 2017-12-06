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

	updateCommand = flag.NewFlagSet(
		"update",
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
	notifyPtr = foreverCommand.Bool(
		"notify",
		false,
		"Will notify to your telegram chat the status of the bot",
	)

	configFilePtr = foreverCommand.String(
		"configFile",
		"/config.yaml",
		"Config file `location`",
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

	//http common
	httpAddrPtr string

	insta *goinsta.Instagram

	client  = Client{}
	likes   = Likes{}
	yamlOpt = YamlOpt{}
)
