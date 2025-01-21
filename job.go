package main

import (
	"context"
	"log"
	"time"

	"gopkg.in/telebot.v4"
)

func RunJob(bot *telebot.Bot) {
	for {
		ctx := context.Background()

		location, err := time.LoadLocation("Asia/Dushanbe")
		if err != nil {
			log.Fatal(err)
		}
		now := time.Now().In(location)

		var shouldBeSent bool

		// 8:30 morning
		if now.Hour() == 8 && now.Minute() == 30 {
			shouldBeSent = true
		}

		// 19:00 evening
		if now.Hour() == 19 && now.Minute() == 0 {
			shouldBeSent = true
		}

		// Send message if it's time to send
		if shouldBeSent {
			text, err := generateText(ctx, GetRandomPrompt())
			if err != nil {
				log.Println(err)
				continue
			}
			err = SendMessage(bot, Config().Recipients, text)
			if err != nil {
				log.Println(err)
				continue
			}
		}

		time.Sleep(1 * time.Minute)
	}
}
