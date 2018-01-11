package main

import (
	"fmt"
	"os"

	"github.com/nlopes/slack"
)

// アーカイブされたチャンネルを除外
const excludeArchived = false

func run(api *slack.Client) {
	channels, err := api.GetChannels(excludeArchived)
	if err != nil {
		fmt.Println("GetChannels error:", err)
		return
	}

	for _, channel := range channels {

		name := channel.Name
		fmt.Println("name:", name)
	}
}

func main() {
	api := slack.New(os.Getenv("SLACKBOT"))
	run(api)
}
