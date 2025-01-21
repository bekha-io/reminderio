package main

import (
	"context"
	"log"
	"os"
	"slices"
	"time"

	_ "time/tzdata"

	"github.com/joho/godotenv"
	tele "gopkg.in/telebot.v4"
)

func main() {
	godotenv.Load()
	InitClient()

	pref := tele.Settings{
		Token:  os.Getenv("BOT_TOKEN"),
		Poller: &tele.LongPoller{Timeout: 10 * time.Second},
	}

	b, err := tele.NewBot(pref)
	if err != nil {
		log.Fatal(err)
		return
	}

	b.Handle("/start", func(ctx tele.Context) error {
		// If user is not in whitelist, then skip
		if !slices.Contains(Config().Recipients, int(ctx.Message().Sender.ID)) {
			log.Printf("user %v is not in whitelist %v. skipping...", int(ctx.Message().Sender.ID), Config().Recipients)
			return nil
		}
		text, err := generateText(context.Background(), GetRandomPrompt())
		if err != nil {
			log.Println(err)
			ctx.Reply(text)
			return err
		}
		ctx.Reply(text)
		return nil
	})

	go RunJob(b)
	b.Start()
}
