package config

import (
	"flag"
	"fmt"
	"os"

	"github.com/spf13/viper"
)

type Config struct {
	type discord struct {
		token			string
	}
	type youtube struct {
		token 		string
	}
	type spotify struct {
		token	string
	}
}

func NewConf(path string) *Config {
	viper.SetConfigName("default")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")
	viper.AddConfigPath(path)
	viper.AutomaticEnv()

	err := viper.ReadInConfig()
	if err != nil {
			fmt.Println("fatal error config file: default \n", err)
			os.Exit(1)
	}
	

	conf := &Config{
		discord{
			token: viper.GetString("app.discord.token")
		}
		youtube{
			token: viper.GetString("app.youtube.token")
		}
		spotify{
			token: viper.GetString("app.spotify.token")
		}
	}

	return conf
}