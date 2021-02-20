package config

import (
	"github.com/spf13/viper"
	"log"
)

type data struct {
	Dir        string `json:"dir"`
	LogsDir    string `json:"logs_dir"`
	Concurrent bool   `json:concurrent`
	Production bool   `json:"production"`
}

var Data *data

func Setup() {
	viper.AddConfigPath(".")
	viper.SetConfigName("config")

	err := viper.ReadInConfig()
	if err != nil {
		log.Fatalf("failed to read the config file: %v", err)
	}

	Data = &data{}
	err = viper.Unmarshal(Data)
	if err != nil {
		panic("unable to decode into config struct")
	}
}
