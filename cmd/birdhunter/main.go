package main

import (
	"flag"
	"fmt"
	//"github.com/janeczku/go-spinner"
)

func main() {
	flag.Parse()

	Init()
	client.login()

	if *popular {
		//s := spinner.StartNew("This may take some while, popular feed...")
		fmt.Println("This may take some while, popular feed...")
		client.getPopularIds()
		//s.Stop()
	} else {
		//s := spinner.StartNew("This may take some while...")
		fmt.Println("This may take some while...")
		client.getTagIds()
		//s.Stop()
	}

	printOut()
}
