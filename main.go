package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"
	"context"

	"github.com/diamondburned/arikawa/v3/state"
	"github.com/f4tal-err0r/discordbot/config"
)

var dg *state.State

func main() {
	conf := config.NewConf("./config.yaml")

	err := state.New("Bot " + conf.Discord.Token)
	if err := dg.Open(context.Background()); err != nil {
		log.Fatalln("failed to open:", err)
	}

	log.Println("Bot is now running.  Press CTRL-C to exit.")
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-sc

	// Cleanly close down the Discord session.
	dg.Close()
}