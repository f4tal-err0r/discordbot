package main

import (
	"fmt"

	"github.com/bwmarrin/discordgo"
)

var (
	conf	*Config
	path	string
)

func init() {
	path = ParseFlags()
	conf = NewConf(path)
}

func main() {
	
}