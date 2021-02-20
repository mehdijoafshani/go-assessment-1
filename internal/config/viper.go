package config

import (
	"github.com/spf13/viper"
	"log"
	"sync"
)

type data struct {
	AccountsDir           string `json:"accountsDir"`
	LogsFile              string `json:"logsFile"`
	IsConcurrent          bool   `json:"isConcurrent"`
	IsProduction          bool   `json:"isProduction"`
	DefaultAccountNumbers int    `json:"defaultAccountNumbers"`
}

var Data *data

// to make sure viper would be setup only once
var setupOnce sync.Once

func SetupViper() {
	setupOnce.Do(func() {
		viper.AddConfigPath(".")
		viper.SetConfigName("config")

		err := viper.ReadInConfig()
		if err != nil {
			log.Fatalf("failed to read the config file: %v", err)
		}

		Data = &data{}
		err = viper.Unmarshal(Data)
		log.Println("LogsDir", Data)
		if err != nil {
			panic("unable to decode into config struct")
		}
	})
}
