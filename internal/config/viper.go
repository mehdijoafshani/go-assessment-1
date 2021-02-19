package config

import (
	"github.com/spf13/viper"
	"log"
)

type Config struct {
	Dir        string `json:"dir"`
	Concurrent bool   `json:concurrent`
}

func Setup() {
	viper.AddConfigPath(".")
	viper.SetConfigName("config")

	err := viper.ReadInConfig()
	if err != nil {
		log.Fatalf("failed to read the config file: %v", err)
	}

	conf := &Config{}
	err = viper.Unmarshal(conf)
	if err != nil {
		panic("unable to decode into config struct")
	}
}
