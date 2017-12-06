package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {
	handler()
}

func required(flagCommand *flag.FlagSet, fl *string) {
	if *fl == "" {
		flagCommand.PrintDefaults()
		os.Exit(0)
	}
}

func printDefaults() {
	fmt.Printf("\npopular\n")
	popularReq()
	popularCommand.PrintDefaults()
	fmt.Printf("\ntag\n")
	tagReq()
	tagCommand.PrintDefaults()
	fmt.Printf("\nforever\n")
	foreverReq()
	foreverCommand.PrintDefaults()
	os.Exit(0)
}

func popularReq() {
	credentialCommon(popularCommand)
	likeCommon(popularCommand)
}

func tagReq() {
	credentialCommon(tagCommand)
	likeCommon(tagCommand)
}

func foreverReq() {
}

func credentialCommon(command *flag.FlagSet) {
	command.StringVar(
		&usernamePtr,
		"u",
		"",
		"Instagram `username`",
	)

	command.StringVar(
		&passwordPtr,
		"p",
		"",
		"Instagram `password`",
	)
}

func likeCommon(command *flag.FlagSet) {
	command.IntVar(
		&likesPtr,
		"likes",
		60,
		"Maximum number of likes per hour (e.g., `60`)",
	)

	command.IntVar(
		&minlikesPtr,
		"min-likes-x-photo",
		60,
		"Minimum number of likes a image needs to have to like it (e.g., `60`)",
	)
}
