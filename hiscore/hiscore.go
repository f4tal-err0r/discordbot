package hiscore

import (
	"log"

	"github.com/diamondburned/arikawa/v3/discord"
	"github.com/diamondburned/arikawa/v3/gateway"
	"github.com/diamondburned/arikawa/v3/api"
	"github.com/f4tal-err0r/discordbot/config"

)

func fetchMessages(apiClient *api.Client, channel discord.ChannelID) ([]discord.Message, error) {
	// var MessagesUntil []discord.Message

	// timestampUntil := time.Now().AddDate(0, 0, -7)

	return apiClient.Messages(channel, 0)
}

func Calc(e *gateway.InteractionCreateEvent) {
	conf := config.NewConf("./config.yaml")

	apiClient := api.NewClient(conf.Discord.Token)

	messages, err := fetchMessages(apiClient, e.ChannelID)
	if err != nil {
		log.Println("Unable to fetch messages")
	}

	for _, message := range messages {
		log.Printf("%s, %s: %s\n", message.Timestamp, message.Author, message.Content)
	}

}