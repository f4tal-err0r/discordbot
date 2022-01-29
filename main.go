package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"
	"context"

	"github.com/diamondburned/arikawa/v3/state"
	"github.com/diamondburned/arikawa/v3/gateway"
	"github.com/f4tal-err0r/discordbot/config"
)


func main() {
	conf := config.NewConf("./config.yaml")

	dg := state.New("Bot " + conf.Discord.Token)

	dg.AddIntents(gateway.IntentGuilds)
	dg.AddIntents(gateway.IntentGuildMessages)
	if err := dg.Open(context.Background()); err != nil {
		log.Fatalln("failed to open:", err)
	}

	initConfig(dg)

	log.Println("Bot is now running.  Press CTRL-C to exit.")
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-sc

	// Cleanly close down the Discord session.
	dg.Close()
}