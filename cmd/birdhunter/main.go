package main

import (
	"flag"
	"fmt"
)

func main() {
	flag.Parse()

	Init()
	client.login()

	if *popular {
		fmt.Println("This may take some while, popular feed...")
		client.getPopularIds()
	} else if *tags != "" {
		fmt.Println("This may take some while...")
		client.getTagIds()
	}

	printOut()
}
