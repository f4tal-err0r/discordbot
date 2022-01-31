package hiscore

import (
	"log"
	"time"

	"github.com/diamondburned/arikawa/v3/discord"
	"github.com/diamondburned/arikawa/v3/gateway"
	"github.com/diamondburned/arikawa/v3/api"
	"github.com/f4tal-err0r/discordbot/config"

)

func fetchMessages(apiClient *api.Client, channel discord.ChannelID) ([]discord.Message, error) {
	timestampUntil := time.Now().AddDate(0, 0, -7)

	s:= discord.NewSnowflake(timestampUntil)

	return apiClient.MessagesAfter(channel, discord.MessageID(s), 0)
}
//TODO: This takes too long, need a loading feature to get around the timeout
func Calc(e *gateway.InteractionCreateEvent) {
	conf := config.NewConf("./config.yaml")
	m := make(map[string]int)

	apiClient := api.NewClient("Bot " + conf.Discord.Token)

	messages, err := fetchMessages(apiClient, e.ChannelID)
	if err != nil {
		log.Printf("Unable to fetch messages: %s", err)
	}

	for _, message := range messages {
		log.Printf("%s, %s: %s\n", message.Timestamp.Format("Mon 01/02/06 03:04"), message.Author, message.Content)
		for _, reactions := range message.Reactions {
			m[message.Author.Username] += reactions.Count
		}
	}

	log.Println(m)

}