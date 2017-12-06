package main

import (
	"crypto/rand"
	"flag"
	"fmt"
	"os"

	"github.com/robfig/cron"
	"github.com/sirupsen/logrus"
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

	crono *cron.Cron

	log = logrus.New()
)

func init() {
	formatter := &logrus.TextFormatter{
		FullTimestamp: true,
	}

	log.Formatter = formatter

	filename := fmt.Sprintf("/out/birdhunter-%v.log", randToken)
	file, err := os.OpenFile(filename, os.O_CREATE|os.O_WRONLY, 0666)

	if err == nil {
		log.Out = file
	}
}

func randToken() string {
	b := make([]byte, 8)
	rand.Read(b)
	return fmt.Sprintf("%x", b)
}
