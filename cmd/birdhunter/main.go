package main

import (
	"flag"
)

func main() {
	flag.Parse()

	Init()
	client.getTagIds()
}
