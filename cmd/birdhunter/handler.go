package main

import (
	"fmt"
	"os"
)

func handler() {
	if len(os.Args) < 2 {
		printDefaults()
	}

	switch os.Args[1] {
	case "popular":
		fmt.Println("This may take some while, popular feed...")
		popularReq()

		popularCommand.Parse(os.Args[2:])

		popular()
	case "tag":
		fmt.Println("This may take some while...")
		tagReq()

		tagCommand.Parse(os.Args[2:])

		tag()
	case "update":
		fmt.Println("Updating the configuration file...")
		updateReq()

		update()
	case "forever":
		foreverReq()

		foreverCommand.Parse(os.Args[2:])

		forever()
		start(httpAddrPtr)
	default:
		printDefaults()
	}
}

func popular() {
	NewClient(
		usernamePtr,
		passwordPtr,
		"",
	)
	NewLikes(
		likesPtr,
		minlikesPtr,
	)

	login()

	getPopularIds()
}

func tag() {
	NewClient(
		usernamePtr,
		passwordPtr,
		*listTagsPtr,
	)
	NewLikes(
		likesPtr,
		minlikesPtr,
	)

	login()

	getTagIds()
}

func update() {
	updateConfig(
		fmt.Sprintf(
			"%v/update",
			httpAddrPtr,
		),
	)
}

func forever() {
	if e := readYaml(*configFilePtr, &yamlOpt); e != nil {
		log.Error("Not be able to update the config.yanl file.")
	}

	setClient(
		yamlOpt.Client.Username,
		yamlOpt.Client.Password,
		yamlOpt.Daemon.Tags,
	)
	NewLikes(
		yamlOpt.Daemon.Likes,
		yamlOpt.Daemon.InPhoto,
	)

	login()

	schedule()
}
