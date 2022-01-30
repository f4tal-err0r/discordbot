package main

import (
	"log"

	"github.com/f4tal-err0r/discordbot/hiscore"

	"github.com/diamondburned/arikawa/v3/api"
	"github.com/diamondburned/arikawa/v3/discord"
	"github.com/diamondburned/arikawa/v3/gateway"
	"github.com/diamondburned/arikawa/v3/state"
	"github.com/diamondburned/arikawa/v3/utils/json/option"

)

var (
	commands = []discord.Command{
		{
			Name: "ping",
			Description: "Basic ping",
		},
		{
			Name: "hiscore",
			Description: "Highest shitposter over the last 7 days",
		},
	}

	commandHandlers = map[string]func(e *gateway.InteractionCreateEvent) api.InteractionResponse {
			"ping": func(e *gateway.InteractionCreateEvent) api.InteractionResponse {
				return api.InteractionResponse{
					Type: api.MessageInteractionWithSource,
					Data: &api.InteractionResponseData{
						Content: option.NewNullableString("https://c.tenor.com/5LGnXPEJU6AAAAAd/gang.gif"),
					},
				}
			},
			"hiscore": func(e *gateway.InteractionCreateEvent) api.InteractionResponse {
				hiscore.Calc(e)
				return api.InteractionResponse{
					Type: api.MessageInteractionWithSource,
					Data: &api.InteractionResponseData{
						Content: option.NewNullableString("Processing messages to logs"),
					},
				}
			},
		}
)

func initCommands(dg *state.State) {

	guildID := discord.GuildID(discord.GuildID(mustSnowflakeEnv("236649685639495680")))

	app, err := dg.CurrentApplication()
	if err != nil {
		log.Fatalln("Failed to get application ID:", err)
	}

	log.Println("Gateway connected. Getting all guild commands.")

	oldcommands, err := dg.GuildCommands(app.ID, guildID)
	if err != nil {
		log.Fatalln("failed to get guild commands:", err)
	}

	for _, command := range oldcommands {
		log.Println("Existing command", command.Name, "found.")
	}

	dg.AddHandler(func(e *gateway.InteractionCreateEvent) {
		if h, ok := commandHandlers[e.Data.(*discord.CommandInteraction).Name]; ok {
			data := h(e)
			if err := dg.RespondInteraction(e.ID, e.Token, data); err != nil {
				log.Println("failed to send interaction callback:", err)
			}
		} else {
			log.Printf("ERROR: %s does not have a command handler", e.Data.(*discord.CommandInteraction).Name )
		}
	})

		_, err = dg.BulkOverwriteGuildCommands(app.ID, guildID, commands)
		if err != nil {
			log.Fatalln("failed to create guild command:", err)
		}
}

func mustSnowflakeEnv(env string) discord.Snowflake {
	s, err := discord.ParseSnowflake(env)
	if err != nil {
		log.Fatalf("Invalid snowflake for $%s: %v", env, err)
	}
	return s
}