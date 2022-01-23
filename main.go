package main

import (
	"fmt"
	"log"
	"flags"

	"github.com/bwmarrin/discordgo"
	"github.com/f4tal-err0r/discordbot/config"
)

var (
	DSession	*discordgo.session
)

func main() {
	conf := config.NewConf("./config")

	discord, err := discordgo.New("Bot " + conf.token)
}