package config

import (
	"fmt"
	"os"

	"github.com/spf13/viper"
)

type Discord struct {
	Token			string
}
type Youtube struct {
	Token 		string
}
type Spotify struct {
	Token	string
}

type Config struct {
	Discord		*Discord
	Youtube		*Youtube
	Spotify		*Spotify
}

func NewConf(path string) *Config {
	viper.SetConfigName("discordbot-config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("..")
	viper.AddConfigPath(path)
	viper.AutomaticEnv()

	err := viper.ReadInConfig()
	if err != nil {
			fmt.Println("fatal error config file: default \n", err)
			os.Exit(1)
	}
	

	conf := &Config{
		Discord: &Discord{
			Token: viper.GetString("app.discord.token"),
		},
		Youtube: &Youtube{
			Token: viper.GetString("app.youtube.token"),
		},
		Spotify: &Spotify{
			Token: viper.GetString("app.spotify.token"),
		},
	}

	return conf
}