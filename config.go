package main

import (
	"log"
	"math/rand"
	"sync"

	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

var (
	once sync.Once
	cfg  *config
)

type config struct {
	Recipients []int
	Prompts    []string
}

func Config() *config {
	once.Do(func() {
		viper.SetConfigName("config")
		viper.SetConfigType("yml")
		viper.AddConfigPath(".")

		if err := viper.ReadInConfig(); err != nil {
			log.Fatalln(err)
		}

		err := viper.Unmarshal(&cfg)
		if err != nil {
			log.Fatal(err)
		}
		viper.WatchConfig()
		viper.OnConfigChange(func(in fsnotify.Event) {
			if err := viper.ReadInConfig(); err != nil {
				log.Fatalln(err)
			}
			err := viper.Unmarshal(&cfg)
			if err != nil {
				log.Fatal(err)
			}
		})

		if len(cfg.Recipients) == 0 {
			log.Fatalf("no recipients configured")
		}
	})
	return cfg
}

func GetRandomPrompt() string {
	prompts := Config().Prompts
	log.Printf("len(prompts) = %v", len(prompts))
	var index int
	if len(prompts) == 1 {
		index = 0
	} else {
		index = rand.Intn(len(prompts))
	}
	return prompts[index]
}
