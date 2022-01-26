package main

import (
	"os"
	"log"

	"github.com/diamondburned/arikawa/v3/api"
	"github.com/diamondburned/arikawa/v3/discord"
	"github.com/diamondburned/arikawa/v3/gateway"
	"github.com/diamondburned/arikawa/v3/utils/json/option"

)

var (
	commands = []*discord.Command{
		{
			Name: "ping",
			Description: "Basic ping",
		},
	}

	commandHandlers = map[string]func(e *gateway.InteractionCreateEvent){
			"ping": func(e *gateway.InteractionCreateEvent) {
				data := api.InteractionResponse{
					Type: api.MessageInteractionWithSource,
					Data: &api.InteractionResponseData{
						Content: option.NewNullableString("https://c.tenor.com/5LGnXPEJU6AAAAAd/gang.gif"),
					},
				}
		
				if err := dg.RespondInteraction(e.ID, e.Token, data); err != nil {
					log.Println("failed to send interaction callback:", err)
				}
			},
		}
)

func init() {

	guildID := discord.GuildID(mustSnowflakeEnv("GUILD_ID"))

	token := os.Getenv("BOT_TOKEN")
	if token == "" {
		log.Fatalln("No $BOT_TOKEN given.")
	}

	app, err := dg.CurrentApplication()
	if err != nil {
		log.Fatalln("Failed to get application ID:", err)
	}

	dg.AddIntents(gateway.IntentGuilds)
	dg.AddIntents(gateway.IntentGuildMessages)

	log.Println("Gateway connected. Getting all guild commands.")

	oldcommands, err := dg.GuildCommands(app.ID, guildID)
	if err != nil {
		log.Fatalln("failed to get guild commands:", err)
	}

	for _, command := range oldcommands {
		log.Println("Existing command", command.Name, "found.")
	}

	dg.AddHandler(func(i *gateway.InteractionCreateEvent) {
		if h, ok := commandHandlers; ok {
			h(i)
		}
	})

	for _, v := range commands {
		_, err := dg.BulkOverwriteGuildCommands(app.ID, guildID, v)
		if err != nil {
			log.Panicf("Cannot create '%v' command: %v", v.Name, err)
		}
	}
}

func mustSnowflakeEnv(env string) discord.Snowflake {
	s, err := discord.ParseSnowflake(os.Getenv(env))
	if err != nil {
		log.Fatalf("Invalid snowflake for $%s: %v", env, err)
	}
	return s
}