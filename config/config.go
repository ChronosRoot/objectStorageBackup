package config

import (
	"fmt"

	"github.com/spf13/viper"
)

var Config *viper.Viper

func SetConfig() {
	Config = viper.New()
	Config.SetConfigFile("config/config.ini")
	Config.SetConfigType("ini")
	if err := Config.ReadInConfig(); err != nil {
		fmt.Println(err)
	}
}
