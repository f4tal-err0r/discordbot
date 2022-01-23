package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/bwmarrin/discordgo"
	"github.com/f4tal-err0r/discordbot/config"
)

func main() {
	conf := config.NewConf("./config.yaml")

	dg, err := discordgo.New("Bot " + conf.Discord.Token)

	dg.AddHandler(messageCreate)

	dg.Identify.Intents = discordgo.IntentsGuildMessages

	err = dg.Open()
	if err != nil {
		log.Fatalf("error opening connection, %s", err)
	}

	log.Println("Bot is now running.  Press CTRL-C to exit.")
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-sc

	// Cleanly close down the Discord session.
	dg.Close()
}

func messageCreate(s *discordgo.Session, m *discordgo.MessageCreate) {


	if m.Author.ID == s.State.User.ID {
		return
	}
	
	if m.Content == "ping" {
		log.Println("ping")
		s.ChannelMessageSend(m.ChannelID, "Pong!")
	}

	if m.Content == "pong" {
		log.Println("pong")
		s.ChannelMessageSend(m.ChannelID, "Ping!")
	}
}